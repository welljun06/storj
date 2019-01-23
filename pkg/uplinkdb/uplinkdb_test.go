// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package uplinkdb_test

import (
	"context"
	"testing"

	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/stretchr/testify/assert"

	"storj.io/storj/internal/testcontext"
	"storj.io/storj/pkg/uplinkdb"
	"storj.io/storj/satellite"
	"storj.io/storj/satellite/satellitedb/satellitedbtest"
)

func TestUplinkdb(t *testing.T) {
	satellitedbtest.Run(t, func(t *testing.T, db satellite.DB) {
		ctx := testcontext.New(t)
		defer ctx.Cleanup()

		testDatabase(ctx, t, db.UplinkDB())
	})
}

func testDatabase(ctx context.Context, t *testing.T, upldb uplinkdb.DB) {
	//testing variables
	uplinkInfo := &uplinkdb.Agreement{
		Agreement: []byte("IamUplinkID"),
		Signature: []byte("IamUplinkPublicKey"),
	}

	serialNum, err := uuid.New()
	assert.NoError(t, err)

	{ // New entry
		err := upldb.CreateAgreement(ctx, serialNum.String(), *uplinkInfo)
		assert.NoError(t, err)
	}

	{ // Get the corresponding Public key for the serialnum
		agreement, err := upldb.GetSignature(ctx, serialNum.String())
		assert.NoError(t, err)
		assert.EqualValues(t, uplinkInfo.Agreement, agreement.Agreement)
		assert.EqualValues(t, uplinkInfo.Signature, agreement.Signature)
	}
}