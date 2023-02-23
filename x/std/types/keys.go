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
	AdminId = []byte{0x00}
	StdId   = []byte{0x01}
	LeaveId = []byte{0x02}
)

func GetstdKey(Stud string) []byte {
	key := make([]byte, len(StdId)+len(Stud))
	copy(key, StdId)
	copy(key[len(StdId):], []byte(Stud))
	return key
}
func GetAdminKey(admin string) []byte {
	key := make([]byte, len(AdminId)+len(admin))
	copy(key, AdminId)
	copy(key[len(AdminId):], []byte(admin))
	return key
}
