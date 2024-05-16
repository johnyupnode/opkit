package domain

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"opkit/x/domain/keeper"
	"opkit/x/domain/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the domain
	for _, elem := range genState.DomainList {
		k.SetDomain(ctx, elem)
	}

	// Set domain count
	k.SetDomainCount(ctx, genState.DomainCount)
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.DomainList = k.GetAllDomain(ctx)
	genesis.DomainCount = k.GetDomainCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
