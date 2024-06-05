package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"insurance/x/insurance/keeper"
	"insurance/x/insurance/types"
)

func SimulateMsgLiquidateinsurance(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgLiquidateinsurance{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the Liquidateinsurance simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "Liquidateinsurance simulation not implemented"), nil, nil
	}
}
