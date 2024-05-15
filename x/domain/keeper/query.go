package keeper

import (
	"opkit/x/domain/types"
)

var _ types.QueryServer = Keeper{}
