package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"insurance/x/insurance/types"
)

func (k msgServer) Cancelinsurance(goCtx context.Context, msg *types.MsgCancelinsurance) (*types.MsgCancelinsuranceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	insurance, found := k.Getinsurance(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if insurance.Borrower != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot cancel: not the borrower")
	}

	borrower, _ := sdk.AccAddressFromBech32(insurance.Borrower)
	collateral, _ := sdk.ParseCoinsNormalized(insurance.Collateral)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	insurance.State = "cancelled"
	k.Setinsurance(ctx, insurance)
	return &types.MsgCancelinsuranceResponse{}, nil
}
