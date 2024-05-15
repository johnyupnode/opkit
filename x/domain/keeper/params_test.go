package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "opkit/testutil/keeper"
	"opkit/x/domain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DomainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
