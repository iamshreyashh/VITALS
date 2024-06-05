package keeper

import (
	"context"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"insurance/x/insurance/types"
)

func (k msgServer) Liquidateinsurance(goCtx context.Context, msg *types.MsgLiquidateinsurance) (*types.MsgLiquidateinsuranceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	insurance, found := k.Getinsurance(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	if insurance.Lender != msg.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot liquidate: not the lender")
	}
	if insurance.State != "approved" {
		return nil, errorsmod.Wrapf(types.ErrWronginsuranceState, "%v", insurance.State)
	}
	lender, _ := sdk.AccAddressFromBech32(insurance.Lender)
	collateral, _ := sdk.ParseCoinsNormalized(insurance.Collateral)
	deadline, err := strconv.ParseInt(insurance.Deadline, 10, 64)
	if err != nil {
		panic(err)
	}
	if ctx.BlockHeight() < deadline {
		return nil, errorsmod.Wrap(types.ErrDeadline, "Cannot liquidate before deadline")
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, lender, collateral)
	if err != nil {
		return nil, err
	}
	insurance.State = "liquidated"
	k.Setinsurance(ctx, insurance)
	return &types.MsgLiquidateinsuranceResponse{}, nil
}
