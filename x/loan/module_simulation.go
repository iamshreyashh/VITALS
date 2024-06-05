package insurance

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"insurance/testutil/sample"
	insurancesimulation "insurance/x/insurance/simulation"
	"insurance/x/insurance/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = insurancesimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgRequestinsurance = "op_weight_msg_request_insurance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRequestinsurance int = 100

	opWeightMsgApproveinsurance = "op_weight_msg_approve_insurance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgApproveinsurance int = 100

	opWeightMsgCancelinsurance = "op_weight_msg_cancel_insurance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelinsurance int = 100

	opWeightMsgRepayinsurance = "op_weight_msg_repay_insurance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRepayinsurance int = 100

	opWeightMsgLiquidateinsurance = "op_weight_msg_liquidate_insurance"
	// TODO: Determine the simulation weight value
	defaultWeightMsgLiquidateinsurance int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	insuranceGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&insuranceGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRequestinsurance int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRequestinsurance, &weightMsgRequestinsurance, nil,
		func(_ *rand.Rand) {
			weightMsgRequestinsurance = defaultWeightMsgRequestinsurance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRequestinsurance,
		insurancesimulation.SimulateMsgRequestinsurance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgApproveinsurance int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgApproveinsurance, &weightMsgApproveinsurance, nil,
		func(_ *rand.Rand) {
			weightMsgApproveinsurance = defaultWeightMsgApproveinsurance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgApproveinsurance,
		insurancesimulation.SimulateMsgApproveinsurance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelinsurance int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelinsurance, &weightMsgCancelinsurance, nil,
		func(_ *rand.Rand) {
			weightMsgCancelinsurance = defaultWeightMsgCancelinsurance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelinsurance,
		insurancesimulation.SimulateMsgCancelinsurance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRepayinsurance int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRepayinsurance, &weightMsgRepayinsurance, nil,
		func(_ *rand.Rand) {
			weightMsgRepayinsurance = defaultWeightMsgRepayinsurance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRepayinsurance,
		insurancesimulation.SimulateMsgRepayinsurance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgLiquidateinsurance int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgLiquidateinsurance, &weightMsgLiquidateinsurance, nil,
		func(_ *rand.Rand) {
			weightMsgLiquidateinsurance = defaultWeightMsgLiquidateinsurance
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgLiquidateinsurance,
		insurancesimulation.SimulateMsgLiquidateinsurance(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgRequestinsurance,
			defaultWeightMsgRequestinsurance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				insurancesimulation.SimulateMsgRequestinsurance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgApproveinsurance,
			defaultWeightMsgApproveinsurance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				insurancesimulation.SimulateMsgApproveinsurance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgCancelinsurance,
			defaultWeightMsgCancelinsurance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				insurancesimulation.SimulateMsgCancelinsurance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgRepayinsurance,
			defaultWeightMsgRepayinsurance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				insurancesimulation.SimulateMsgRepayinsurance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgLiquidateinsurance,
			defaultWeightMsgLiquidateinsurance,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				insurancesimulation.SimulateMsgLiquidateinsurance(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
