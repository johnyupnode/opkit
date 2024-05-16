package keeper

import (
	"context"

	"opkit/x/domain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DomainAll(ctx context.Context, req *types.QueryAllDomainRequest) (*types.QueryAllDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var domains []types.Domain

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	domainStore := prefix.NewStore(store, types.KeyPrefix(types.DomainKey))

	pageRes, err := query.Paginate(domainStore, req.Pagination, func(key []byte, value []byte) error {
		var domain types.Domain
		if err := k.cdc.Unmarshal(value, &domain); err != nil {
			return err
		}

		domains = append(domains, domain)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDomainResponse{Domain: domains, Pagination: pageRes}, nil
}

func (k Keeper) Domain(ctx context.Context, req *types.QueryGetDomainRequest) (*types.QueryGetDomainResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	domain, found := k.GetDomain(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetDomainResponse{Domain: domain}, nil
}
