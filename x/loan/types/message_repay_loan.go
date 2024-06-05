package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRepayinsurance = "repay_insurance"

var _ sdk.Msg = &MsgRepayinsurance{}

func NewMsgRepayinsurance(creator string, id uint64) *MsgRepayinsurance {
	return &MsgRepayinsurance{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRepayinsurance) Route() string {
	return RouterKey
}

func (msg *MsgRepayinsurance) Type() string {
	return TypeMsgRepayinsurance
}

func (msg *MsgRepayinsurance) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRepayinsurance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRepayinsurance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
