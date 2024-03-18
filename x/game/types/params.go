package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	KeyMaxPlayersPerGame = []byte("MaxPlayersPerGame")
	// TODO: Determine the default value
	DefaultMaxPlayersPerGame uint64 = 0
)

var (
	KeyMinDistanceToWin = []byte("MinDistanceToWin")
	// TODO: Determine the default value
	DefaultMinDistanceToWin uint64 = 0
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	maxPlayersPerGame uint64,
	minDistanceToWin uint64,
) Params {
	return Params{
		MaxPlayersPerGame: maxPlayersPerGame,
		MinDistanceToWin:  minDistanceToWin,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultMaxPlayersPerGame,
		DefaultMinDistanceToWin,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxPlayersPerGame, &p.MaxPlayersPerGame, validateMaxPlayersPerGame),
		paramtypes.NewParamSetPair(KeyMinDistanceToWin, &p.MinDistanceToWin, validateMinDistanceToWin),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateMaxPlayersPerGame(p.MaxPlayersPerGame); err != nil {
		return err
	}

	if err := validateMinDistanceToWin(p.MinDistanceToWin); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// validateMaxPlayersPerGame validates the MaxPlayersPerGame param
func validateMaxPlayersPerGame(v interface{}) error {
	maxPlayersPerGame, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = maxPlayersPerGame

	return nil
}

// validateMinDistanceToWin validates the MinDistanceToWin param
func validateMinDistanceToWin(v interface{}) error {
	minDistanceToWin, ok := v.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", v)
	}

	// TODO implement validation
	_ = minDistanceToWin

	return nil
}
