package keeper

import (
	"context"

	"opkit/x/domain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListDomainByString(goCtx context.Context, req *types.QueryListDomainByStringRequest) (*types.QueryListDomainByStringResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	domains, err := k.GetDomainsFromIndexer(ctx, req.Key, req.Value)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListDomainByStringResponse{Domain: domains}, nil
}
