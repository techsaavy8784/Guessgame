package game

import (
	"math/rand"

	"game/testutil/sample"
	gamesimulation "game/x/game/simulation"
	"game/x/game/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = gamesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateGame = "op_weight_msg_create_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateGame int = 100

	opWeightMsgSubmitGuess = "op_weight_msg_submit_guess"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitGuess int = 100

	opWeightMsgEndGame = "op_weight_msg_end_game"
	// TODO: Determine the simulation weight value
	defaultWeightMsgEndGame int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	gameGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&gameGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	gameParams := types.DefaultParams()
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMaxPlayersPerGame), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(gameParams.MaxPlayersPerGame))
		}),
		simulation.NewSimParamChange(types.ModuleName, string(types.KeyMinDistanceToWin), func(r *rand.Rand) string {
			return string(types.Amino.MustMarshalJSON(gameParams.MinDistanceToWin))
		}),
	}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateGame, &weightMsgCreateGame, nil,
		func(_ *rand.Rand) {
			weightMsgCreateGame = defaultWeightMsgCreateGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateGame,
		gamesimulation.SimulateMsgCreateGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSubmitGuess int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitGuess, &weightMsgSubmitGuess, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitGuess = defaultWeightMsgSubmitGuess
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitGuess,
		gamesimulation.SimulateMsgSubmitGuess(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgEndGame int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgEndGame, &weightMsgEndGame, nil,
		func(_ *rand.Rand) {
			weightMsgEndGame = defaultWeightMsgEndGame
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgEndGame,
		gamesimulation.SimulateMsgEndGame(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
