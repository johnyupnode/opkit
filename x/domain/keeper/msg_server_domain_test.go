package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"opkit/x/domain/types"
)

func TestDomainMsgServerCreate(t *testing.T) {
	_, srv, ctx := setupMsgServer(t)
	wctx := sdk.UnwrapSDKContext(ctx)

	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateDomain(wctx, &types.MsgCreateDomain{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestDomainMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateDomain
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateDomain{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateDomain{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateDomain{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateDomain(wctx, &types.MsgCreateDomain{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateDomain(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestDomainMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteDomain
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteDomain{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteDomain{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteDomain{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			_, srv, ctx := setupMsgServer(t)
			wctx := sdk.UnwrapSDKContext(ctx)

			_, err := srv.CreateDomain(wctx, &types.MsgCreateDomain{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteDomain(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
