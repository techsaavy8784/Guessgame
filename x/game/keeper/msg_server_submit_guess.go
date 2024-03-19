package keeper

import (
	"context"
	"strconv"

	"game/x/game/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) SubmitGuess(goCtx context.Context, msg *types.MsgSubmitGuess) (*types.MsgSubmitGuessResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	currentTime := ctx.BlockTime().Unix()

	maxPlayersPerGame := uint64(3)

	var player = types.Player{
		Address: msg.Creator,
		Guess:   msg.Guess,
	}
	// Retrieve the game by ID.
	game, found := k.GetGame(ctx, msg.GameId)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "game %d doesn't exist", msg.GameId)
	}
	// Check if the game is still active.
	if game.State != "created" {
		return nil, errorsmod.Wrapf(types.ErrWrongGameState, "game is in %v state", game.State)
	}
	gameEndTime := game.Time + game.Duration
	if uint64(currentTime) >= gameEndTime {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "game time is over")
	}

	// Check if the player has already submitted a guess.
	if k.HasPlayerSubmittedGuess(ctx, msg.Creator, msg.GameId) {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "player has already submitted a guess")
	}

	// Contraint on max players has been moved to the params validation function.

	if uint64(len(game.Players)) >= maxPlayersPerGame {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "max number of players reached")
	}

	// Ensure that the creator isn't joining their own game.
	if msg.Creator == game.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "creator cannot join their own game")
	}

	// Charge the entry fee to the player's account.
	guesser, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid guesser address")
	}

	fee, err := sdk.ParseCoinsNormalized(game.EntryFee)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid entry fee format")
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, guesser, types.ModuleName, fee); err != nil {
		return nil, err
	}

	// Append the player to the game.
	game.Players = append(game.Players, &player)

	// Update the game in the store.
	k.SetGame(ctx, game)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"SubmittedGuess",
			sdk.NewAttribute("gameId", strconv.Itoa(int(msg.GameId))),
			sdk.NewAttribute("player", msg.Creator),
			sdk.NewAttribute("guessNumber", strconv.Itoa(int(msg.Guess))),
		),
	)

	return &types.MsgSubmitGuessResponse{}, nil
}

// Check if the player has already submitted a guess for the given game ID.
func (k Keeper) HasPlayerSubmittedGuess(ctx sdk.Context, address string, gameID uint64) bool {
	game, found := k.GetGame(ctx, gameID)
	if !found {
		return false
	}
	for _, player := range game.Players {
		if player.Address == address {
			return true
		}
	}
	return false
}
