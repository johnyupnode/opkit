package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "opkit/testutil/keeper"
	"opkit/testutil/nullify"
	"opkit/x/domain/types"
)

func TestDomainQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	msgs := createNDomain(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetDomainRequest
		response *types.QueryGetDomainResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetDomainRequest{Id: msgs[0].Id},
			response: &types.QueryGetDomainResponse{Domain: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetDomainRequest{Id: msgs[1].Id},
			response: &types.QueryGetDomainResponse{Domain: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetDomainRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Domain(ctx, tc.request)
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

func TestDomainQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.DomainKeeper(t)
	msgs := createNDomain(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllDomainRequest {
		return &types.QueryAllDomainRequest{
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
			resp, err := keeper.DomainAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Domain), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Domain),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.DomainAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Domain), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Domain),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.DomainAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Domain),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.DomainAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
