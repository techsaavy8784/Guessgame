package keeper_test

import (
	"testing"

	testkeeper "game/testutil/keeper"
	"game/x/game/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GameKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.MaxPlayersPerGame, k.MaxPlayersPerGame(ctx))
	require.EqualValues(t, params.MinDistanceToWin, k.MinDistanceToWin(ctx))
}
