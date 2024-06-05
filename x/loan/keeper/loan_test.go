package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "insurance/testutil/keeper"
	"insurance/testutil/nullify"
	"insurance/x/insurance/keeper"
	"insurance/x/insurance/types"
)

func createNinsurance(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.insurance {
	items := make([]types.insurance, n)
	for i := range items {
		items[i].Id = keeper.Appendinsurance(ctx, items[i])
	}
	return items
}

func TestinsuranceGet(t *testing.T) {
	keeper, ctx := keepertest.insuranceKeeper(t)
	items := createNinsurance(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.Getinsurance(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestinsuranceRemove(t *testing.T) {
	keeper, ctx := keepertest.insuranceKeeper(t)
	items := createNinsurance(keeper, ctx, 10)
	for _, item := range items {
		keeper.Removeinsurance(ctx, item.Id)
		_, found := keeper.Getinsurance(ctx, item.Id)
		require.False(t, found)
	}
}

func TestinsuranceGetAll(t *testing.T) {
	keeper, ctx := keepertest.insuranceKeeper(t)
	items := createNinsurance(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllinsurance(ctx)),
	)
}

func TestinsuranceCount(t *testing.T) {
	keeper, ctx := keepertest.insuranceKeeper(t)
	items := createNinsurance(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetinsuranceCount(ctx))
}
