package keeper

import (
	"context"
	"opkit/x/domain/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// SetReward set a specific reward in the store
func (k Keeper) SetReward(ctx context.Context, reward types.Reward) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RewardKey))
	b := k.cdc.MustMarshal(&reward)
	store.Set(GetRewardBytes(reward.Domain), b)
}

// GetReward returns a reward from its id
func (k Keeper) GetReward(ctx context.Context, domain string) (val types.Reward, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.RewardKey))
	b := store.Get(GetRewardBytes(domain))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func GetRewardBytes(domain string) []byte {
	bz := types.KeyPrefix(types.RewardKey)
	bz = append(bz, []byte("/")...)
	bz = append(bz, []byte(domain)...)
	return bz
}
