package keeper

import (
	"context"

	"opkit/x/domain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) SetPrimaryDomain(goCtx context.Context, msg *types.MsgSetPrimaryDomain) (*types.MsgSetPrimaryDomainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.logger.Error("msgServer.SetPrimaryDomain", "msg", msg.String())

	err := k.Keeper.SetPrimaryDomain(ctx, msg.Domain, msg.Creator)
	if err != nil {
		return nil, err
	}

	return &types.MsgSetPrimaryDomainResponse{}, nil
}
