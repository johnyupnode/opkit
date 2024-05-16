package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"opkit/x/domain/types"
)

func (k Keeper) ListDomainEvm(goCtx context.Context, req *types.QueryListDomainEvmRequest) (*types.QueryListDomainEvmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	domains, err := k.GetIndexerDomains(ctx, "evm", req.Address)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListDomainEvmResponse{Domain: domains}, nil
}
