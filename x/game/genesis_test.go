package game_test

import (
	"testing"

	keepertest "game/testutil/keeper"
	"game/testutil/nullify"
	"game/x/game"
	"game/x/game/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		PlayerList: []types.Player{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PlayerCount: 2,
		GameList: []types.Game{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		GameCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GameKeeper(t)
	game.InitGenesis(ctx, *k, genesisState)
	got := game.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.PlayerList, got.PlayerList)
	require.Equal(t, genesisState.PlayerCount, got.PlayerCount)
	require.ElementsMatch(t, genesisState.GameList, got.GameList)
	require.Equal(t, genesisState.GameCount, got.GameCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
