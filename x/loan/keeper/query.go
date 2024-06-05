package keeper

import (
	"insurance/x/insurance/types"
)

var _ types.QueryServer = Keeper{}
