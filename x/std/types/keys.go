package types

import "fmt"

const (
	// ModuleName is the name of the module
	ModuleName = "lms"

	// StoreKey is the store key string for slashing
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
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
	AdminId     = []byte{0x00}
	StdId       = []byte{0x01}
	LeaveId     = []byte{0x02}
	CounterId   = []byte{0x03}
	LeaveAccept = []byte{0x04}
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
	fmt.Println("adminkeyprefix: ", AdminId)
	fmt.Println("adminkey: ", key)
	return key
}
func LeaveStoreKey(sid string, studentLeavesCount string) []byte {
	key := make([]byte, len(LeaveId)+len(sid)+len(studentLeavesCount))
	copy(key, LeaveId)
	copy(key[len(LeaveId):], []byte(sid))
	copy(key[len(LeaveId)+len(sid):], []byte(studentLeavesCount))
	return key
}

func StudentLeavesCounterKey(sid string) []byte {
	key := make([]byte, len(CounterId)+len(sid))
	copy(key, CounterId)
	copy(key[len(CounterId):], []byte(sid))
	return key
}

func AcceptLeaveStoreKey(leaveid string) []byte {
	key := make([]byte, len(LeaveAccept)+len(leaveid))
	copy(key, leaveid)
	copy(key[len(LeaveAccept):], []byte(leaveid))
	return key
}
