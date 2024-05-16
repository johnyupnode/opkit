package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"opkit/x/domain/types"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		bankKeeper types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	bankKeeper types.BankKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		bankKeeper: bankKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetDomainsFromIndexer from indexer
func (k Keeper) GetDomainsFromIndexer(ctx sdk.Context, key, value string) ([]types.Domain, error) {
	return []types.Domain{
		{
			Id:        1,
			Domain:    "domain1",
			Owner:     "owner1",
			Timestamp: ctx.BlockTime().UTC().String(),
			Txhash:    "txhash1",
		},
		{
			Id:        2,
			Domain:    "domain2",
			Owner:     "owner2",
			Timestamp: ctx.BlockTime().UTC().String(),
			Txhash:    "txhash2",
		},
	}, nil
}

// SetPrimaryDomain sets the primary domain.
func (k Keeper) SetPrimaryDomain(ctx sdk.Context, domain string, sender string) error {
	domains, err := k.GetDomainsFromIndexer(ctx, "opkit", sender)
	if err != nil {
		k.logger.Error("failed to get domains", "error", err)
		return err
	}

	if len(domains) == 0 {
		k.logger.Error("no domains found", "sender", sender)
		return fmt.Errorf("no domains found for sender: %s", sender)
	}

	// Set the primary domain
	isSet := false
	for _, d := range domains {
		if d.Domain == domain {
			// Set the primary domain
			k.logger.Info("set primary domain", "domain", d.Domain)
			isSet = true
			k.SetDomain(ctx, d)
		}
	}

	if !isSet {
		k.logger.Error("domain not found", "domain", domain)
		return fmt.Errorf("domain not found: %s", domain)
	}

	return nil
}

// GetDomainInfoFromIndexer returns a domain from the indexer.
func (k Keeper) GetDomainInfoFromIndexer(ctx sdk.Context, domain string) (types.Domain, error) {
	return types.Domain{
		Id:        1,
		Domain:    "domainIndexer",
		Owner:     "owner1",
		Timestamp: ctx.BlockTime().UTC().String(),
		Txhash:    "txhash1",
	}, nil
}
