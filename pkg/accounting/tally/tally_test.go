// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package tally_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/internal/testidentity"
	"storj.io/storj/internal/teststorj"
	"storj.io/storj/pkg/accounting/tally"
	"storj.io/storj/pkg/bwagreement"
	"storj.io/storj/pkg/bwagreement/testbwagreement"
	"storj.io/storj/pkg/identity"
	"storj.io/storj/pkg/overlay/mocks"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/pointerdb"
	"storj.io/storj/satellite/satellitedb"
	"storj.io/storj/storage/teststore"
)

func TestQueryNoAgreements(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	// TODO: use testplanet

	service := pointerdb.NewService(zap.NewNop(), teststore.New())
	overlayServer := mocks.NewOverlay([]*pb.Node{})
	db, err := satellitedb.NewInMemory()
	assert.NoError(t, err)
	defer ctx.Check(db.Close)
	assert.NoError(t, db.CreateTables())

	tally := tally.New(zap.NewNop(), db.Accounting(), db.BandwidthAgreement(), service, overlayServer, 0, time.Second)

	err = tally.QueryBW(ctx)
	assert.NoError(t, err)
}

func TestQueryWithBw(t *testing.T) {
	ctx := testcontext.New(t)
	defer ctx.Cleanup()

	// TODO: use testplanet

	service := pointerdb.NewService(zap.NewNop(), teststore.New())
	overlayServer := mocks.NewOverlay([]*pb.Node{})

	db, err := satellitedb.NewInMemory()
	assert.NoError(t, err)
	defer ctx.Check(db.Close)

	assert.NoError(t, db.CreateTables())

	bwDb := db.BandwidthAgreement()
	tally := tally.New(zap.NewNop(), db.Accounting(), bwDb, service, overlayServer, 0, time.Second)

	//get a private key
	fiC, err := testidentity.NewTestIdentity(ctx)
	assert.NoError(t, err)

	makeBWA(ctx, t, bwDb, fiC, pb.BandwidthAction_PUT)
	makeBWA(ctx, t, bwDb, fiC, pb.BandwidthAction_GET)
	makeBWA(ctx, t, bwDb, fiC, pb.BandwidthAction_GET_AUDIT)
	makeBWA(ctx, t, bwDb, fiC, pb.BandwidthAction_GET_REPAIR)
	makeBWA(ctx, t, bwDb, fiC, pb.BandwidthAction_PUT_REPAIR)

	//check the db
	err = tally.QueryBW(ctx)
	assert.NoError(t, err)
}

func makeBWA(ctx context.Context, t *testing.T, bwDb bwagreement.DB, fiC *identity.FullIdentity, action pb.BandwidthAction) {
	//generate an agreement with the key
	pba, err := testbwagreement.GeneratePayerBandwidthAllocation(action, fiC, fiC, time.Hour)
	assert.NoError(t, err)
	rba, err := testbwagreement.GenerateRenterBandwidthAllocation(pba, teststorj.NodeIDFromString("StorageNodeID"), fiC, 666)
	assert.NoError(t, err)
	//save to db
	err = bwDb.CreateAgreement(ctx, rba)
	assert.NoError(t, err)
}
