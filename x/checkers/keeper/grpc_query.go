package keeper

import (
	"github.com/mwackowski/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
