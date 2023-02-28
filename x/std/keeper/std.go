package keeper

import (
	"coslms/x/std/types"
	"fmt"
	"log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/ruthishvitiwt/coslms/x/std/types"
	// "google.golang.org/genproto/googleapis/actions/sdk/v2"
)

// func NewStudentKeeper(cdc codec.BinaryCodec, storekey storetypes.StoreKey) Keeper {
// 	if _, err := sdk.AccAddressFromBech32("h"); err != nil {
// 		panic(fmt.Errorf("invalid  authority address: %w", err))
// 	}
// 	return Keeper{
// 		cdc:      cdc,
// 		storeKey: storekey,
// 	}
// }
func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (k Keeper) RegisterAdmin(ctx sdk.Context, registerAdmin *types.RegisterAdminRequest) error {

	if registerAdmin.Name == "" {
		return types.ErrAdminNameNil
	} else if registerAdmin.Address == "" {
		return types.ErrAdminAddressNil
	} else {
		store := ctx.KVStore(k.storeKey)
		marshalRegisterAdmin, err := k.cdc.Marshal(registerAdmin)
		handleError(err)
		store.Set(types.GetAdminKey(registerAdmin.Address), marshalRegisterAdmin)
		return nil
	}
}

func (k Keeper) GetAdmin(ctx sdk.Context, id string) {
	store := ctx.KVStore(k.storeKey)
	fmt.Println(types.GetAdminKey(id))
	v := store.Get(types.GetAdminKey(id))
	fmt.Println(v)
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

// func (k Keeper) GetStudent(ctx sdk.Context, id string) {
// 	store := ctx.KVStore(k.storeKey)
// 	m := store.Get(types.StdKey(id))
// 	fmt.Println(m)

// }
func (k Keeper) GetStudent(ctx sdk.Context, getStudents *types.GetStudentsRequest) {
	store := ctx.KVStore(k.storeKey)
	var t types.Student
	itr := store.Iterator(types.StdId, nil)
	for ; itr.Valid(); itr.Next() {
		k.cdc.Unmarshal(itr.Value(), &t)
		fmt.Println(t)
	}
}
