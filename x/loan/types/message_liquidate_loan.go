package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgLiquidateinsurance = "liquidate_insurance"

var _ sdk.Msg = &MsgLiquidateinsurance{}

func NewMsgLiquidateinsurance(creator string, id uint64) *MsgLiquidateinsurance {
	return &MsgLiquidateinsurance{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgLiquidateinsurance) Route() string {
	return RouterKey
}

func (msg *MsgLiquidateinsurance) Type() string {
	return TypeMsgLiquidateinsurance
}

func (msg *MsgLiquidateinsurance) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgLiquidateinsurance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgLiquidateinsurance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
