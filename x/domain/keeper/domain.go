package keeper

import (
	"context"
	"encoding/binary"

	"opkit/x/domain/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetDomainCount get the total number of domain
func (k Keeper) GetDomainCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.DomainCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDomainCount set the total number of domain
func (k Keeper) SetDomainCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.DomainCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendDomain appends a domain in the store with a new id and update the count
func (k Keeper) AppendDomain(
	ctx context.Context,
	domain types.Domain,
) uint64 {
	// Create the domain
	count := k.GetDomainCount(ctx)

	// Set the ID of the appended value
	domain.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainKey))
	appendedValue := k.cdc.MustMarshal(&domain)
	store.Set(GetDomainIDBytes(domain.Id), appendedValue)

	// Update domain count
	k.SetDomainCount(ctx, count+1)

	return count
}

// SetDomain set a specific domain in the store
func (k Keeper) SetDomain(ctx context.Context, domain types.Domain) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainKey))
	b := k.cdc.MustMarshal(&domain)
	store.Set(GetDomainIDBytes(domain.Id), b)
}

// GetDomain returns a domain from its id
func (k Keeper) GetDomain(ctx context.Context, id uint64) (val types.Domain, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainKey))
	b := store.Get(GetDomainIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDomain removes a domain from the store
func (k Keeper) RemoveDomain(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainKey))
	store.Delete(GetDomainIDBytes(id))
}

// GetAllDomain returns all domain
func (k Keeper) GetAllDomain(ctx context.Context) (list []types.Domain) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.DomainKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Domain
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDomainIDBytes returns the byte representation of the ID
func GetDomainIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.DomainKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
