package keeper

import (
	"context"

	"opkit/x/domain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Info(goCtx context.Context, req *types.QueryInfoRequest) (*types.QueryInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	domain, err := k.GetDomainInfoFromIndexer(ctx, req.Domain)
	if err != nil {
		return nil, err
	}

	return &types.QueryInfoResponse{Domain: domain}, nil
}
