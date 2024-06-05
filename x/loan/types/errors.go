package types

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrWronginsuranceState = sdkerrors.Register(ModuleName, 2, "wrong insurance state")
	ErrDeadline            = sdkerrors.Register(ModuleName, 3, "deadline")
)
