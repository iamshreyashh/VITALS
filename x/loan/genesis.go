package insurance

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"insurance/x/insurance/keeper"
	"insurance/x/insurance/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the insurance
	for _, elem := range genState.insuranceList {
		k.Setinsurance(ctx, elem)
	}

	// Set insurance count
	k.SetinsuranceCount(ctx, genState.insuranceCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.insuranceList = k.GetAllinsurance(ctx)
	genesis.insuranceCount = k.GetinsuranceCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
