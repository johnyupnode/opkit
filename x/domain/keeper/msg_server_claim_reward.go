package keeper

import (
	"context"

	"opkit/x/domain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ClaimReward(goCtx context.Context, msg *types.MsgClaimReward) (*types.MsgClaimRewardResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.ClaimReward(ctx, msg.Domain, msg.Creator); err != nil {
		return nil, err
	}

	return &types.MsgClaimRewardResponse{}, nil
}
