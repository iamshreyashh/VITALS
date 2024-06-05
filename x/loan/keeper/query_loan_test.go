package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "insurance/testutil/keeper"
	"insurance/testutil/nullify"
	"insurance/x/insurance/types"
)

func TestinsuranceQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.insuranceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNinsurance(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetinsuranceRequest
		response *types.QueryGetinsuranceResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetinsuranceRequest{Id: msgs[0].Id},
			response: &types.QueryGetinsuranceResponse{insurance: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetinsuranceRequest{Id: msgs[1].Id},
			response: &types.QueryGetinsuranceResponse{insurance: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetinsuranceRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.insurance(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestinsuranceQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.insuranceKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNinsurance(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllinsuranceRequest {
		return &types.QueryAllinsuranceRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.insuranceAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.insurance), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.insurance),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.insuranceAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.insurance), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.insurance),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.insuranceAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.insurance),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.insuranceAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
