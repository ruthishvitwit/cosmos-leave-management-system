package types

const (
	// ModuleName is the name of the module
	ModuleName = "std"

	// StoreKey is the store key string for slashing
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName
)

// Keys for slashing store
// Items are stored with the following key: values
//
// - 0x01<consAddrLen (1 Byte)><consAddress_Bytes>: ValidatorSigningInfo
//
// - 0x02<consAddrLen (1 Byte)><consAddress_Bytes><period_Bytes>: bool
//
// - 0x03<accAddrLen (1 Byte)><accAddr_Bytes>: cryptotypes.PubKey

var (
	AdminKey = []byte{0x00}
	StdKey   = []byte{0x01}
	LeaveKey = []byte{0x02}
)

// func GetadminKey(byte){
// }
