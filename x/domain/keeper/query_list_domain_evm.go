package keeper

import (
	"context"

	"opkit/x/domain/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListDomainEvm(goCtx context.Context, req *types.QueryListDomainEvmRequest) (*types.QueryListDomainEvmResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	if !common.IsHexAddress(req.Address) {
		return nil, status.Error(codes.InvalidArgument, "invalid address evm")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	domains, err := k.GetDomainOwnerFromIndexer(ctx, req.Address)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListDomainEvmResponse{Domain: domains}, nil
}
