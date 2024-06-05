package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"insurance/x/insurance/types"
)

func (k msgServer) Repayinsurance(goCtx context.Context, msg *types.MsgRepayinsurance) (*types.MsgRepayinsuranceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	insurance, found := k.Getinsurance(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrKeyNotFound, "key %d doesn't exist", msg.Id)
	}
	lender, _ := sdk.AccAddressFromBech32(insurance.Lender)
	borrower, _ := sdk.AccAddressFromBech32(insurance.Borrower)
	if msg.Creator != insurance.Borrower {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "Cannot repay: not the borrower")
	}
	amount, _ := sdk.ParseCoinsNormalized(insurance.Amount)
	fee, _ := sdk.ParseCoinsNormalized(insurance.Fee)
	collateral, _ := sdk.ParseCoinsNormalized(insurance.Collateral)
	err := k.bankKeeper.SendCoins(ctx, borrower, lender, amount)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoins(ctx, borrower, lender, fee)
	if err != nil {
		return nil, err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	insurance.State = "repayed"
	k.Setinsurance(ctx, insurance)
	return &types.MsgRepayinsuranceResponse{}, nil
}
