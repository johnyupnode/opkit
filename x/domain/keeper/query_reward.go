package keeper

import (
	"context"
	"cosmossdk.io/math"

	"opkit/x/domain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Reward(goCtx context.Context, req *types.QueryRewardRequest) (*types.QueryRewardResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	reward, found := k.GetReward(ctx, req.Domain)
	if !found {
		amountClaim := sdk.NewCoin("stake", math.NewInt(AmountClaim))
		reward = types.Reward{
			Domain:    req.Domain,
			IsClaimed: false,
			Amount:    amountClaim,
		}
	}

	return &types.QueryRewardResponse{Reward: reward}, nil
}
