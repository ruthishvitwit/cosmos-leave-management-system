package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrNoAdminName    = sdkerrors.Register(ModuleName, 1, "Admin name cannot be empty")
	ErrNoAdminAddress = sdkerrors.Register(ModuleName, 2, "Admin address cannot be empty")
)
