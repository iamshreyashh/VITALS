package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"insurance/x/insurance/types"
)

func (k msgServer) Requestinsurance(goCtx context.Context, msg *types.MsgRequestinsurance) (*types.MsgRequestinsuranceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var insurance = types.insurance{
		Amount:     msg.Amount,
		Fee:        msg.Fee,
		Collateral: msg.Collateral,
		Deadline:   msg.Deadline,
		State:      "requested",
		Borrower:   msg.Creator,
	}
	borrower, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	collateral, err := sdk.ParseCoinsNormalized(insurance.Collateral)
	if err != nil {
		panic(err)
	}
	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, borrower, types.ModuleName, collateral)
	if sdkError != nil {
		return nil, sdkError
	}
	k.Appendinsurance(ctx, insurance)
	return &types.MsgRequestinsuranceResponse{}, nil
}
