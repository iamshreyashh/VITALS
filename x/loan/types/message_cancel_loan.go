package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelinsurance = "cancel_insurance"

var _ sdk.Msg = &MsgCancelinsurance{}

func NewMsgCancelinsurance(creator string, id uint64) *MsgCancelinsurance {
	return &MsgCancelinsurance{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCancelinsurance) Route() string {
	return RouterKey
}

func (msg *MsgCancelinsurance) Type() string {
	return TypeMsgCancelinsurance
}

func (msg *MsgCancelinsurance) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelinsurance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelinsurance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
