package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"insurance/x/insurance/types"
)

func (k msgServer) Approveinsurance(goCtx context.Context, msg *types.MsgApproveinsurance) (*types.MsgApproveinsuranceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	insurance, found := k.Getinsurance(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(insurance.Borrower)
	amount, err := sdk.ParseCoinsNormalized(insurance.Amount)
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}
	insurance.Lender = msg.Creator
	insurance.State = "approved"
	k.Setinsurance(ctx, insurance)
	return &types.MsgApproveinsuranceResponse{}, nil
}
