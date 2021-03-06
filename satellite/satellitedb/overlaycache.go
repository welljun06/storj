// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package satellitedb

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeebo/errs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/storj"
	"storj.io/storj/pkg/utils"
	dbx "storj.io/storj/satellite/satellitedb/dbx"
	"storj.io/storj/storage"
)

var _ overlay.DB = (*overlaycache)(nil)

type overlaycache struct {
	db *dbx.DB
}

type getNodesRequest struct {
	minReputation         *pb.NodeStats
	freeBandwidth         int64
	freeDisk              int64
	excluded              []pb.NodeID
	reputableNodeAmount   int64
	newNodeAmount         int64
	newNodeAuditThreshold int64
}

// FilterNodes looks up nodes based on reputation requirements
func (cache *overlaycache) FilterNodes(ctx context.Context, req *overlay.FilterNodesRequest) ([]*pb.Node, error) {
	reputableNodeAmount := req.MinNodes
	if reputableNodeAmount <= 0 {
		reputableNodeAmount = req.Opts.GetAmount()
	}

	getReputableReq := &getNodesRequest{
		minReputation:         req.MinReputation,
		freeBandwidth:         req.Opts.GetRestrictions().FreeBandwidth,
		freeDisk:              req.Opts.GetRestrictions().FreeDisk,
		excluded:              req.Opts.ExcludedNodes,
		reputableNodeAmount:   reputableNodeAmount,
		newNodeAuditThreshold: req.NewNodeAuditThreshold,
	}

	reputableNodes, err := cache.getReputableNodes(ctx, getReputableReq)
	if err != nil {
		return nil, err
	}

	newNodeAmount := int64(float64(reputableNodeAmount) * req.NewNodePercentage)

	getNewReq := &getNodesRequest{
		freeBandwidth:         req.Opts.GetRestrictions().FreeBandwidth,
		freeDisk:              req.Opts.GetRestrictions().FreeDisk,
		excluded:              req.Opts.ExcludedNodes,
		newNodeAmount:         newNodeAmount,
		newNodeAuditThreshold: req.NewNodeAuditThreshold,
	}

	newNodes, err := cache.getNewNodes(ctx, getNewReq)
	if err != nil {
		return nil, err
	}

	var allNodes []*pb.Node
	allNodes = append(allNodes, reputableNodes...)
	allNodes = append(allNodes, newNodes...)

	if int64(len(reputableNodes)) < reputableNodeAmount {
		err := status.Errorf(codes.ResourceExhausted, fmt.Sprintf("requested %d reputable nodes, only %d reputable nodes matched the criteria requested",
			reputableNodeAmount, len(reputableNodes)))
		return allNodes, err
	}

	return allNodes, nil
}

func (cache *overlaycache) getReputableNodes(ctx context.Context, req *getNodesRequest) ([]*pb.Node, error) {
	rows, err := cache.findReputableNodesQuery(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = utils.CombineErrors(err, rows.Close())
	}()

	reputableNodes, err := sqlRowsToNodes(rows)
	if err != nil {
		return nil, err
	}

	return reputableNodes, nil
}

func (cache *overlaycache) getNewNodes(ctx context.Context, req *getNodesRequest) ([]*pb.Node, error) {
	rows, err := cache.findNewNodesQuery(ctx, req)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = utils.CombineErrors(err, rows.Close())
	}()

	newNodes, err := sqlRowsToNodes(rows)
	if err != nil {
		return nil, err
	}

	return newNodes, nil
}

func sqlRowsToNodes(rows *sql.Rows) (nodes []*pb.Node, err error) {
	for rows.Next() {
		overlayNode := &dbx.OverlayCacheNode{}
		err = rows.Scan(&overlayNode.NodeId, &overlayNode.NodeType,
			&overlayNode.Address, &overlayNode.FreeBandwidth, &overlayNode.FreeDisk,
			&overlayNode.AuditSuccessRatio, &overlayNode.AuditUptimeRatio,
			&overlayNode.AuditCount, &overlayNode.AuditSuccessCount,
			&overlayNode.UptimeCount, &overlayNode.UptimeSuccessCount)
		if err != nil {
			return nil, err
		}

		node, err := convertOverlayNode(overlayNode)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

func (cache *overlaycache) findReputableNodesQuery(ctx context.Context, req *getNodesRequest) (*sql.Rows, error) {
	auditCount := req.minReputation.AuditCount
	if req.newNodeAuditThreshold > auditCount {
		auditCount = req.newNodeAuditThreshold
	}
	auditSuccessRatio := req.minReputation.AuditSuccessRatio
	uptimeCount := req.minReputation.UptimeCount
	uptimeRatio := req.minReputation.UptimeRatio
	nodeAmt := req.reputableNodeAmount

	var rows *sql.Rows
	var err error
	var nodeTypeStorage int32 = 2

	args := make([]interface{}, len(req.excluded))
	for i, id := range req.excluded {
		args[i] = id.Bytes()
	}

	args = append(args, auditCount, auditSuccessRatio, uptimeCount, uptimeRatio,
		req.freeBandwidth, req.freeDisk, nodeTypeStorage, nodeAmt)

	// This queries for nodes whose audit counts are greater than or equal to
	// the new node audit threshold and the minimum reputation audit count.
	rows, err = cache.db.Query(`SELECT node_id,
	node_type, address, free_bandwidth, free_disk, audit_success_ratio,
	audit_uptime_ratio, audit_count, audit_success_count, uptime_count,
	uptime_success_count
	FROM overlay_cache_nodes
	WHERE node_id NOT IN (`+strings.Join(sliceOfCopies("?", len(req.excluded)), ", ")+`)
	AND audit_count >= ?
	AND audit_success_ratio >= ?
	AND uptime_count >= ?
	AND audit_uptime_ratio >= ?
	AND free_bandwidth >= ?
	AND free_disk >= ?
	AND node_type == ?
	LIMIT ?
	`,
		args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func sliceOfCopies(val string, count int) []string {
	slice := make([]string, count)
	for i := range slice {
		slice[i] = val
	}
	return slice
}

func (cache *overlaycache) findNewNodesQuery(ctx context.Context, req *getNodesRequest) (*sql.Rows, error) {
	var rows *sql.Rows
	var err error

	var nodeTypeStorage int32 = 2

	args := make([]interface{}, len(req.excluded))
	for i, id := range req.excluded {
		args[i] = id.Bytes()
	}

	args = append(args, req.newNodeAuditThreshold, req.freeBandwidth,
		req.freeDisk, nodeTypeStorage, req.newNodeAmount)

	rows, err = cache.db.Query(cache.db.Rebind(`SELECT node_id,
		node_type, address, free_bandwidth, free_disk, audit_success_ratio,
		audit_uptime_ratio, audit_count, audit_success_count, uptime_count,
		uptime_success_count
		FROM overlay_cache_nodes
		WHERE node_id NOT IN (`+strings.Join(sliceOfCopies("?", len(req.excluded)), ", ")+`)	
		AND audit_count < ?
		AND free_bandwidth >= ?
		AND free_disk >= ?
		AND node_type == ?
		LIMIT ?
	`),
		args...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

// Get looks up the node by nodeID
func (cache *overlaycache) Get(ctx context.Context, id storj.NodeID) (*pb.Node, error) {
	if id.IsZero() {
		return nil, overlay.ErrEmptyNode
	}

	node, err := cache.db.Get_OverlayCacheNode_By_NodeId(ctx,
		dbx.OverlayCacheNode_NodeId(id.Bytes()),
	)
	if err == sql.ErrNoRows {
		return nil, overlay.ErrNodeNotFound
	}
	if err != nil {
		return nil, err
	}

	return convertOverlayNode(node)
}

// GetAll looks up nodes based on the ids from the overlay cache
func (cache *overlaycache) GetAll(ctx context.Context, ids storj.NodeIDList) ([]*pb.Node, error) {
	infos := make([]*pb.Node, len(ids))
	for i, id := range ids {
		// TODO: abort on canceled context
		info, err := cache.Get(ctx, id)
		if err != nil {
			continue
		}
		infos[i] = info
	}
	return infos, nil
}

// List lists nodes starting from cursor
func (cache *overlaycache) List(ctx context.Context, cursor storj.NodeID, limit int) ([]*pb.Node, error) {
	// TODO: handle this nicer
	if limit <= 0 || limit > storage.LookupLimit {
		limit = storage.LookupLimit
	}

	dbxInfos, err := cache.db.Limited_OverlayCacheNode_By_NodeId_GreaterOrEqual(ctx,
		dbx.OverlayCacheNode_NodeId(cursor.Bytes()),
		limit, 0,
	)
	if err != nil {
		return nil, err
	}

	infos := make([]*pb.Node, len(dbxInfos))
	for i, dbxInfo := range dbxInfos {
		infos[i], err = convertOverlayNode(dbxInfo)
		if err != nil {
			return nil, err
		}
	}
	return infos, nil
}

// Paginate will run through
func (cache *overlaycache) Paginate(ctx context.Context, offset int64, limit int) ([]*pb.Node, bool, error) {
	cursor := storj.NodeID{}

	// more represents end of table. If there are more rows in the database, more will be true.
	more := true

	if limit <= 0 || limit > storage.LookupLimit {
		limit = storage.LookupLimit
	}

	dbxInfos, err := cache.db.Limited_OverlayCacheNode_By_NodeId_GreaterOrEqual(ctx,
		dbx.OverlayCacheNode_NodeId(cursor.Bytes()),
		limit, offset,
	)

	if err != nil {
		return nil, false, err
	}

	if len(dbxInfos) < limit {
		more = false
	}

	infos := make([]*pb.Node, len(dbxInfos))
	for i, dbxInfo := range dbxInfos {
		infos[i], err = convertOverlayNode(dbxInfo)
		if err != nil {
			return nil, false, err
		}
	}
	return infos, more, nil
}

// Update updates node information
func (cache *overlaycache) Update(ctx context.Context, info *pb.Node) (err error) {
	if info == nil || info.Id.IsZero() {
		return overlay.ErrEmptyNode
	}

	tx, err := cache.db.Open(ctx)
	if err != nil {
		return Error.Wrap(err)
	}

	// TODO: use upsert
	_, err = tx.Get_OverlayCacheNode_By_NodeId(ctx,
		dbx.OverlayCacheNode_NodeId(info.Id.Bytes()),
	)

	address := info.Address
	if address == nil {
		address = &pb.NodeAddress{}
	}

	metadata := info.Metadata
	if metadata == nil {
		metadata = &pb.NodeMetadata{}
	}

	restrictions := info.Restrictions
	if restrictions == nil {
		restrictions = &pb.NodeRestrictions{
			FreeBandwidth: -1,
			FreeDisk:      -1,
		}
	}

	reputation := info.Reputation
	if reputation == nil {
		reputation = &pb.NodeStats{}
	}

	if err != nil {
		_, err = tx.Create_OverlayCacheNode(
			ctx,
			dbx.OverlayCacheNode_NodeId(info.Id.Bytes()),

			dbx.OverlayCacheNode_NodeType(int(info.Type)),
			dbx.OverlayCacheNode_Address(address.Address),
			dbx.OverlayCacheNode_Protocol(int(address.Transport)),

			dbx.OverlayCacheNode_OperatorEmail(metadata.Email),
			dbx.OverlayCacheNode_OperatorWallet(metadata.Wallet),

			dbx.OverlayCacheNode_FreeBandwidth(restrictions.FreeBandwidth),
			dbx.OverlayCacheNode_FreeDisk(restrictions.FreeDisk),

			dbx.OverlayCacheNode_Latency90(reputation.Latency_90),
			dbx.OverlayCacheNode_AuditSuccessRatio(reputation.AuditSuccessRatio),
			dbx.OverlayCacheNode_AuditUptimeRatio(reputation.UptimeRatio),
			dbx.OverlayCacheNode_AuditCount(reputation.AuditCount),
			dbx.OverlayCacheNode_AuditSuccessCount(reputation.AuditSuccessCount),

			dbx.OverlayCacheNode_UptimeCount(reputation.UptimeCount),
			dbx.OverlayCacheNode_UptimeSuccessCount(reputation.UptimeSuccessCount),
		)
		if err != nil {
			return Error.Wrap(errs.Combine(err, tx.Rollback()))
		}
	} else {
		update := dbx.OverlayCacheNode_Update_Fields{
			// TODO: should we be able to update node type?
			Address:  dbx.OverlayCacheNode_Address(address.Address),
			Protocol: dbx.OverlayCacheNode_Protocol(int(address.Transport)),

			Latency90:          dbx.OverlayCacheNode_Latency90(info.Reputation.Latency_90),
			AuditSuccessRatio:  dbx.OverlayCacheNode_AuditSuccessRatio(info.Reputation.AuditSuccessRatio),
			AuditUptimeRatio:   dbx.OverlayCacheNode_AuditUptimeRatio(info.Reputation.UptimeRatio),
			AuditCount:         dbx.OverlayCacheNode_AuditCount(info.Reputation.AuditCount),
			AuditSuccessCount:  dbx.OverlayCacheNode_AuditSuccessCount(info.Reputation.AuditSuccessCount),
			UptimeCount:        dbx.OverlayCacheNode_UptimeCount(info.Reputation.UptimeCount),
			UptimeSuccessCount: dbx.OverlayCacheNode_UptimeSuccessCount(info.Reputation.UptimeSuccessCount),
		}

		if info.Metadata != nil {
			update.OperatorEmail = dbx.OverlayCacheNode_OperatorEmail(info.Metadata.Email)
			update.OperatorWallet = dbx.OverlayCacheNode_OperatorWallet(info.Metadata.Wallet)
		}

		if info.Restrictions != nil {
			update.FreeBandwidth = dbx.OverlayCacheNode_FreeBandwidth(restrictions.FreeBandwidth)
			update.FreeDisk = dbx.OverlayCacheNode_FreeDisk(restrictions.FreeDisk)
		}

		_, err := tx.Update_OverlayCacheNode_By_NodeId(ctx,
			dbx.OverlayCacheNode_NodeId(info.Id.Bytes()),
			update,
		)
		if err != nil {
			return Error.Wrap(errs.Combine(err, tx.Rollback()))
		}
	}

	return Error.Wrap(tx.Commit())
}

// Delete deletes node based on id
func (cache *overlaycache) Delete(ctx context.Context, id storj.NodeID) error {
	_, err := cache.db.Delete_OverlayCacheNode_By_NodeId(ctx,
		dbx.OverlayCacheNode_NodeId(id.Bytes()),
	)
	return err
}

func convertOverlayNode(info *dbx.OverlayCacheNode) (*pb.Node, error) {
	if info == nil {
		return nil, Error.New("missing info")
	}

	id, err := storj.NodeIDFromBytes(info.NodeId)
	if err != nil {
		return nil, err
	}

	node := &pb.Node{
		Id:   id,
		Type: pb.NodeType(info.NodeType),
		Address: &pb.NodeAddress{
			Address:   info.Address,
			Transport: pb.NodeTransport(info.Protocol),
		},
		Metadata: &pb.NodeMetadata{
			Email:  info.OperatorEmail,
			Wallet: info.OperatorWallet,
		},
		Restrictions: &pb.NodeRestrictions{
			FreeBandwidth: info.FreeBandwidth,
			FreeDisk:      info.FreeDisk,
		},
		Reputation: &pb.NodeStats{
			NodeId:             id,
			Latency_90:         info.Latency90,
			AuditSuccessRatio:  info.AuditSuccessRatio,
			UptimeRatio:        info.AuditUptimeRatio,
			AuditCount:         info.AuditCount,
			AuditSuccessCount:  info.AuditSuccessCount,
			UptimeCount:        info.UptimeCount,
			UptimeSuccessCount: info.UptimeSuccessCount,
		},
	}

	if node.Address.Address == "" {
		node.Address = nil
	}
	if node.Metadata.Email == "" && node.Metadata.Wallet == "" {
		node.Metadata = nil
	}
	if node.Restrictions.FreeBandwidth < 0 && node.Restrictions.FreeDisk < 0 {
		node.Restrictions = nil
	}
	if node.Reputation.Latency_90 < 0 {
		node.Reputation = nil
	}

	return node, nil
}

//GetWalletAddress gets the node's wallet address
func (cache *overlaycache) GetWalletAddress(ctx context.Context, id storj.NodeID) (string, error) {
	w, err := cache.db.Get_OverlayCacheNode_OperatorWallet_By_NodeId(ctx, dbx.OverlayCacheNode_NodeId(id.Bytes()))
	if err != nil {
		return "", err
	}
	return w.OperatorWallet, nil
}
