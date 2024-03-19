package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateGame = "create_game"

var _ sdk.Msg = &MsgCreateGame{}

func NewMsgCreateGame(creator string, secretNumber uint64, reward string, entryFee string, duration uint64) *MsgCreateGame {
	return &MsgCreateGame{
		Creator:      creator,
		SecretNumber: secretNumber,
		Reward:       reward,
		EntryFee:     entryFee,
		Duration:     duration,
	}
}

func (msg *MsgCreateGame) Route() string {
	return RouterKey
}

func (msg *MsgCreateGame) Type() string {
	return TypeMsgCreateGame
}

func (msg *MsgCreateGame) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateGame) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateGame) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	reward, _ := sdk.ParseCoinsNormalized(msg.Reward)
	if !reward.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reward is not a valid Coins object")
	}
	if reward.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "reward is empty")
	}
	fee, _ := sdk.ParseCoinsNormalized(msg.EntryFee)
	if !fee.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee is not a valid Coins object")
	}

	if msg.Duration <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "deadline should be a positive integer")
	}

	if msg.SecretNumber <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "secretNumber should be a positive integer")
	}
	return nil
}
