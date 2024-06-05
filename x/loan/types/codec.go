package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRequestinsurance{}, "insurance/Requestinsurance", nil)
	cdc.RegisterConcrete(&MsgApproveinsurance{}, "insurance/Approveinsurance", nil)
	cdc.RegisterConcrete(&MsgCancelinsurance{}, "insurance/Cancelinsurance", nil)
	cdc.RegisterConcrete(&MsgRepayinsurance{}, "insurance/Repayinsurance", nil)
	cdc.RegisterConcrete(&MsgLiquidateinsurance{}, "insurance/Liquidateinsurance", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestinsurance{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveinsurance{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelinsurance{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRepayinsurance{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgLiquidateinsurance{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
