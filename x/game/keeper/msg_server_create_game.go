package keeper

import (
	"context"
	"strconv"

	"game/x/game/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateGame(goCtx context.Context, msg *types.MsgCreateGame) (*types.MsgCreateGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTime := ctx.BlockTime().Unix()

	// Create a new Game instance.
	game := types.Game{
		Creator:      msg.Creator,
		SecretNumber: msg.SecretNumber,
		Reward:       msg.Reward,
		EntryFee:     msg.EntryFee,
		Duration:     msg.Duration,
		State:        "created",
		Time:         uint64(currentTime),
	}

	// Append the game to the store and retrieve its ID.
	id := k.AppendGame(ctx, game)

	// Emit an event for game creation.
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"GameCreated", // Event type
			sdk.NewAttribute("gameId", strconv.Itoa(int(id))), // Attribute - game ID
			sdk.NewAttribute("creator", msg.Creator),          // Attribute - game creator
		))

	// Return the ID of the newly created game in the response.
	return &types.MsgCreateGameResponse{GameId: id}, nil

}
