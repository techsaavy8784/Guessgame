package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/game module sentinel errors
var (
	ErrSample         = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrWrongGameState = sdkerrors.Register(ModuleName, 2, "wrong game state")
)
