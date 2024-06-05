package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"insurance/x/insurance/types"
)

func (k Keeper) insuranceAll(goCtx context.Context, req *types.QueryAllinsuranceRequest) (*types.QueryAllinsuranceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var insurances []types.insurance
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	insuranceStore := prefix.NewStore(store, types.KeyPrefix(types.insuranceKey))

	pageRes, err := query.Paginate(insuranceStore, req.Pagination, func(key []byte, value []byte) error {
		var insurance types.insurance
		if err := k.cdc.Unmarshal(value, &insurance); err != nil {
			return err
		}

		insurances = append(insurances, insurance)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllinsuranceResponse{insurance: insurances, Pagination: pageRes}, nil
}

func (k Keeper) insurance(goCtx context.Context, req *types.QueryGetinsuranceRequest) (*types.QueryGetinsuranceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	insurance, found := k.Getinsurance(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetinsuranceResponse{insurance: insurance}, nil
}
