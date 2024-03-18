package keeper

import (
	"game/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MaxPlayersPerGame(ctx),
		k.MinDistanceToWin(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MaxPlayersPerGame returns the MaxPlayersPerGame param
func (k Keeper) MaxPlayersPerGame(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMaxPlayersPerGame, &res)
	return
}

// MinDistanceToWin returns the MinDistanceToWin param
func (k Keeper) MinDistanceToWin(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinDistanceToWin, &res)
	return
}
