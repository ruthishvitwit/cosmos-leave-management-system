package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers concrete types on LegacyAmino codec
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	legacy.RegisterAminoMsg(cdc, &RegisterAdminRequest{}, "coslms/RegisterAdminRequest")
	legacy.RegisterAminoMsg(cdc, &AddStudentRequest{}, "coslms/AddStudentRequest")
	legacy.RegisterAminoMsg(cdc, &ApplyLeaveRequest{}, "coslms/ApplyLeaveRequest")
	//legacy.RegisterAminoMsg(cdc, &AcceptLeaveRequest{}, "coslms/AcceptLeaveRequest")

}

// RegisterInterfaces registers the interfaces types with the Interface Registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&RegisterAdminRequest{},
		&AddStudentRequest{},
		&ApplyLeaveRequest{},
		//&AcceptLeaveRequest{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &Msg_ServiceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	sdk.RegisterLegacyAminoCodec(amino)
}
