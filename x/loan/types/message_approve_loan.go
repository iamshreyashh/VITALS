package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveinsurance = "approve_insurance"

var _ sdk.Msg = &MsgApproveinsurance{}

func NewMsgApproveinsurance(creator string, id uint64) *MsgApproveinsurance {
	return &MsgApproveinsurance{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgApproveinsurance) Route() string {
	return RouterKey
}

func (msg *MsgApproveinsurance) Type() string {
	return TypeMsgApproveinsurance
}

func (msg *MsgApproveinsurance) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveinsurance) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveinsurance) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
