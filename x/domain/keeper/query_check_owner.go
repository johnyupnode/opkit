package keeper

import (
	"context"

	"opkit/x/domain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CheckOwner(goCtx context.Context, req *types.QueryCheckOwnerRequest) (*types.QueryCheckOwnerResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	domains, err := k.GetDomainOwnerFromIndexer(ctx, req.Address)
	if err != nil {
		k.logger.Error("could not get domain from indexer", "error", err)
		return nil, err
	}

	return &types.QueryCheckOwnerResponse{Domain: domains}, nil
}
