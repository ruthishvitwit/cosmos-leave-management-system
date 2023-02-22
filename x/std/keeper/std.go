package keeper

import (
	"fmt"

	"coslms/x/std/types"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/ruthishvitiwt/coslms/x/std/types"
	// "google.golang.org/genproto/googleapis/actions/sdk/v2"
)

func NewStudentKeeper(cdc codec.BinaryCodec, storekey storetypes.StoreKey) Keeper {
	if _, err := sdk.AccAddressFromBech32("h"); err != nil {
		panic(fmt.Errorf("invalid  authority address: %w", err))
	}
	return Keeper{
		//cdc:      cdc,
		storeKey: storekey,
	}
}

func (k Keeper) AddStudent(ctx sdk.Context, req *types.AddStudentRequest) error {
	// if _, err := sdk.AccAddressFromBech32(req.Address); err != nil {
	// 	panic(fmt.Errorf("Invalid Authority Address:%w", err))
	// }
	store := ctx.KVStore(k.storeKey)
	bz, err := k.cdc.Marshal(req)
	if err != nil {
		return err
	} else {
		store.Set(types.StdKey, bz)
	}
	return nil
}
