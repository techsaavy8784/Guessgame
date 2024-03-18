package game

import (
	"game/x/game/keeper"
	"game/x/game/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the player
	for _, elem := range genState.PlayerList {
		k.SetPlayer(ctx, elem)
	}

	// Set player count
	k.SetPlayerCount(ctx, genState.PlayerCount)
	// Set all the game
	for _, elem := range genState.GameList {
		k.SetGame(ctx, elem)
	}

	// Set game count
	k.SetGameCount(ctx, genState.GameCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PlayerList = k.GetAllPlayer(ctx)
	genesis.PlayerCount = k.GetPlayerCount(ctx)
	genesis.GameList = k.GetAllGame(ctx)
	genesis.GameCount = k.GetGameCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
