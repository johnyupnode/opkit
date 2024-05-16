package domain_test

import (
	"testing"

	keepertest "opkit/testutil/keeper"
	"opkit/testutil/nullify"
	domain "opkit/x/domain/module"
	"opkit/x/domain/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DomainList: []types.Domain{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DomainCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DomainKeeper(t)
	domain.InitGenesis(ctx, k, genesisState)
	got := domain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DomainList, got.DomainList)
	require.Equal(t, genesisState.DomainCount, got.DomainCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
