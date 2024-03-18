package keeper

import (
	"context"
	"math"

	"game/x/game/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EndGame(goCtx context.Context, msg *types.MsgEndGame) (*types.MsgEndGameResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	game, found := k.GetGame(ctx, msg.GameId)
	if !found {
		return nil, errorsmod.Wrapf(sdkerrors.ErrNotFound, "game %d doesn't exist", msg.GameId)
	}
	if game.Creator != msg.Creator {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "only the game creator can end the game")
	}
	if game.State != "created" {
		return nil, errorsmod.Wrapf(types.ErrWrongGameState, "game state %v is not created", game.State)
	}

	minDistanceToWin := uint64(100)
	secretNumber := game.SecretNumber
	creator, _ := sdk.AccAddressFromBech32(game.Creator)

	for _, player := range game.Players {
		distance := math.Abs(float64(secretNumber - player.Guess))
		playerAddr, err := sdk.AccAddressFromBech32(player.Address)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid player address")
		}

		deposit, err := sdk.ParseCoinsNormalized(game.EntryFee)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid entry fee format")
		}

		if uint64(distance) <= minDistanceToWin {
			// Player guessed correctly, reward them and return their deposit.
			reward, err := sdk.ParseCoinsNormalized(game.Reward)
			if err != nil {
				return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidCoins, "invalid reward format")
			}
			// Send reward to player's address.
			if err := k.bankKeeper.SendCoins(ctx, creator, playerAddr, reward); err != nil {
				return nil, err
			}
			// return player's deposit
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, playerAddr, deposit); err != nil {
				return nil, err
			}
		} else {
			// Player guessed incorrectly, forfeit their deposit to the game creator.
			if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, deposit); err != nil {
				return nil, err
			}
		}
	}

	// Update game state to "ended".
	game.State = "ended"
	k.SetGame(ctx, game)

	return &types.MsgEndGameResponse{}, nil
}
