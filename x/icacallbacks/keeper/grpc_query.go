package keeper

import (
	"github.com/TessorNetwork/dredger/v4/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
