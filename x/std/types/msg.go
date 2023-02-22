package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// verify interface at compile time
var (
	_ sdk.Msg = &AddStudentRequest{}
	_ sdk.Msg = &RegisterAdminRequest{}
)

// add student
func NewAddStudentRequest(admin string, students []*Student) *AddStudentRequest {
	return &AddStudentRequest{
		Admin: admin, Students: students,
	}
}

// GetSigners returns the expected signers for MsgUnjail.
func (msg AddStudentRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Admin)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg AddStudentRequest) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

// ValidateBasic does a sanity check on the provided message.
func (msg AddStudentRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Admin); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil
}

func NewRegisterAdminReq(accAddr sdk.AccAddress) *RegisterAdminRequest {
	return &RegisterAdminRequest{
		Address: accAddr.String(),
	}
}

func (msg RegisterAdminRequest) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

func (msg RegisterAdminRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Address) //ValAddressFromBech32(msg.AdminAddress)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg RegisterAdminRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil
}
func NewApplyLeaveReq(accountAddr sdk.AccAddress) *ApplyLeaveRequest {
	return &ApplyLeaveRequest{
		Address: accountAddr.String(),
	}
}

func (msg ApplyLeaveRequest) GetSignBytes() []byte {
	b := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(b)
}

func (msg ApplyLeaveRequest) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.AccAddressFromBech32(msg.Address)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

func (msg ApplyLeaveRequest) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Address); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("account input address: %s", err)
	}
	return nil
}
