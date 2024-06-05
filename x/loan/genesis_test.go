package insurance_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "insurance/testutil/keeper"
	"insurance/testutil/nullify"
	"insurance/x/insurance"
	"insurance/x/insurance/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		insuranceList: []types.insurance{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		insuranceCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.insuranceKeeper(t)
	insurance.InitGenesis(ctx, *k, genesisState)
	got := insurance.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.insuranceList, got.insuranceList)
	require.Equal(t, genesisState.insuranceCount, got.insuranceCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
