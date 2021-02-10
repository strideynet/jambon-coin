package keeper

import (
	"github.com/strideynet/jambon-coin/x/jamboncoin/types"
)

var _ types.QueryServer = Keeper{}
