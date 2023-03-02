package keeper

import (
	"coslms/x/std/types"
	"fmt"
	"log"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/ruthishvitiwt/coslms/x/std/types"
	// "google.golang.org/genproto/googleapis/actions/sdk/v2"
)

//	func NewStudentKeeper(cdc codec.BinaryCodec, storekey storetypes.StoreKey) Keeper {
//		if _, err := sdk.AccAddressFromBech32("h"); err != nil {
//			panic(fmt.Errorf("invalid  authority address: %w", err))
//		}
//		return Keeper{
//			cdc:      cdc,
//			storeKey: storekey,
//		}
//	}
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (k Keeper) RegisterAdmin(ctx sdk.Context, registerAdmin *types.RegisterAdminRequest) error {

	if registerAdmin.Name == "" {
		return types.ErrNoAdminName
	} else if registerAdmin.Address == "" {
		return types.ErrNoAdminAddress
	} else {
		store := ctx.KVStore(k.storeKey)
		marshalRegisterAdmin, err := k.cdc.Marshal(registerAdmin)
		handleError(err)
		store.Set(types.GetAdminKey(registerAdmin.Address), marshalRegisterAdmin)
		return nil
	}
}

func (k Keeper) AddStudent(ctx sdk.Context, req *types.AddStudentRequest) error {
	store := ctx.KVStore(k.storeKey)
	m, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.StdId, m)
		fmt.Println("completed")
	}
	return nil
}

func (k Keeper) GetStudent(ctx sdk.Context, getStudents *types.GetStudentsRequest) {
	store := ctx.KVStore(k.storeKey)
	var t types.Student
	fmt.Println("the students are ", t)
	itr := store.Iterator(types.StdId, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		fmt.Println("address:", t.Address, "  Id:", t.Id, "  name:", t.Name)
	}
}
func (k Keeper) ApplyLeave(ctx sdk.Context, applyLeave *types.ApplyLeaveRequest) string {
	store := ctx.KVStore(k.storeKey)
	marshalApplyLeave, err := k.cdc.Marshal(applyLeave)
	handleError(err)
	addr := types.LeavesCounterKey(sdk.AccAddress(string(applyLeave.Address)).String())
	counter := store.Get(addr)
	if counter == nil {
		store.Set(addr, []byte("1"))
	} else {
		c, err := strconv.Atoi(string(counter))
		handleError(err)
		c = c + 1
		store.Set(addr, []byte(fmt.Sprint(c)))
	}
	counter = store.Get(addr)
	handleError(err)
	store.Set(types.StudentLeaveKey(applyLeave.Address, string(counter)), marshalApplyLeave)
	// id := types.LeaveStoreKey(applyLeave.Address, string(counter))
	return "Leave Applied Successfully"
}

func (k Keeper) AcceptLeave(ctx sdk.Context, acceptLeave *types.AcceptLeaveRequest) string {
	store := ctx.KVStore(k.storeKey)
	marshalAcceptLeave, err := k.cdc.Marshal(acceptLeave)
	handleError(err)
	leaveid := store.Get(types.LeavesCounterKey(acceptLeave.LeaveId))
	store.Set(types.AcceptLeaveKey(string(leaveid)), marshalAcceptLeave)
	return "Leave Status Updated"
}

// func (k Keeper) GetLeaveStatus(ctx sdk.Context, getLeaveStatus *types.GetLeaveStatusRequest) string {
// 	store := ctx.KVStore(k.storeKey)

func (k Keeper) GetLeaveRequests(ctx sdk.Context, getLeaveRequests *types.GetLeaveRequestsRequest) {
	store := ctx.KVStore(k.storeKey)
	var t types.ApplyLeaveRequest
	itr := store.Iterator(types.LeaveId, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		fmt.Println(t)
	}
}

func (k Keeper) GetLeaveApprovedRequests(ctx sdk.Context, getLeaveApprovedRequests *types.GetLeaveApprovedRequestsRequest) {
	store := ctx.KVStore(k.storeKey)
	var t types.ApplyLeaveRequest
	itr := store.Iterator(types.LeaveAccept, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		fmt.Println(t)
	}
}
