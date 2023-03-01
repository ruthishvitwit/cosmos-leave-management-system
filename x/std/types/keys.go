package types

import "fmt"

const (
	// ModuleName is the name of the module
	ModuleName = "lms"
	StoreKey   = ModuleName
	// RouterKey is the message route
	RouterKey = ModuleName
	// RouterKey is the query route
	QuerierRoute = ModuleName
)

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
func StudentLeaveKey(id string, LeavesCount string) []byte {
	key := make([]byte, len(LeaveId)+len(id)+len(LeavesCount))
	copy(key, LeaveId)
	copy(key[len(LeaveId):], []byte(id))
	copy(key[len(LeaveId)+len(id):], []byte(LeavesCount))
	return key
}

func LeavesCounterKey(id string) []byte {
	key := make([]byte, len(CounterId)+len(id))
	copy(key, CounterId)
	copy(key[len(CounterId):], []byte(id))
	return key
}

func AcceptLeaveKey(leavedata string) []byte {
	key := make([]byte, len(LeaveAccept)+len(leavedata))
	copy(key, leavedata)
	copy(key[len(LeaveAccept):], []byte(leavedata))
	return key
}
