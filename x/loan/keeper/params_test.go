package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "insurance/testutil/keeper"
	"insurance/x/insurance/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.insuranceKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
