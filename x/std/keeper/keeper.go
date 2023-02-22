package keeper

import (
	//"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/gogoproto/codec"
)

type Keeper struct {
	storeKey storetypes.StoreKey

	cdc codec.Codec
}
