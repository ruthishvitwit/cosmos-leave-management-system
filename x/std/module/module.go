package module

import (
	"context"
	"encoding/json"

	"coslms/x/std/keeper"
	"coslms/x/std/types"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/gov/simulation"

	// "github.com/cosmos/cosmos-sdk/x/nft/client/cli"
	"coslms/x/std/client/cli"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	"github.com/spf13/cobra"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
	// _ module.AppModuleSimulation = AppModule{}
)

type AppModuleBasic struct {
	cdc codec.Codec
}

var _ module.AppModuleBasic = AppModuleBasic{}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the mint module's types on the given LegacyAmino codec.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

func (b AppModuleBasic) RegisterInterfaces(_ cdctypes.InterfaceRegistry) {}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	return nil
}
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// GetTxCmd returns the root tx command for the feegrant module.
func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

// GetQueryCmd returns no root query command for the feegrant module.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

func (AppModule) ConsensusVersion() uint64 { return 1 }

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return nil
}

type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
	}
}

// Name returns the mint module's name.
func (AppModule) Name() string {
	return types.ModuleName
}

// RegisterInvariants registers the mint module invariants.
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Deprecated: Route returns the message routing key for the mint module.
func (AppModule) Route() sdk.Route { return sdk.Route{} }

// QuerierRoute returns the mint module's querier route name.
func (AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

// LegacyQuerierHandler returns the mint module sdk.Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a gRPC query service to respond to the
// module-specific gRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// InitGenesis performs genesis initialization for the mint module. It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	// var genesisState types.GenesisState
	// cdc.MustUnmarshalJSON(data, &genesisState)

	// am.keeper.InitGenesis(ctx, am.authKeeper, &genesisState)
	// return []abci.ValidatorUpdate{}
	return nil
}

// ExportGenesis returns the exported genesis state as raw bytes for the mint
// module.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	// gs := am.keeper.ExportGenesis(ctx)
	// return cdc.MustMarshalJSON(gs)
	return nil
}

func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}

// import (
// 	"context"
// 	"coslms/x/std/types"

// 	"coslms/x/std/keeper"
// 	"encoding/json"

// 	//abci "github.com/cometbft/cometbft/abci/types"
// 	gwruntime "github.com/grpc-ecosystem/grpc-gateway/runtime"
// 	"github.com/spf13/cobra"

// 	"cosmossdk.io/core/appmodule"

// 	//"cosmossdk.io/store/types"

// 	sdkclient "github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

// 	//"github.com/cosmos/cosmos-sdk/store/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
// 	"github.com/cosmos/cosmos-sdk/types/module"
// 	ab "github.com/tendermint/tendermint/abci/types"

// 	//simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

// 	"github.com/cosmos/cosmos-sdk/x/nft"
// 	"github.com/cosmos/cosmos-sdk/x/nft/client/cli"
// 	//"github.com/cosmos/cosmos-sdk/x/nft/keeper"
// 	//"github.com/cosmos/cosmos-sdk/x/gov/simulation"
// 	//"github.com/cosmos/cosmos-sdk/x/nft"
// 	//"github.com/cosmos/cosmos-sdk/x/nft/client/cli"
// 	//"github.com/cosmos/cosmos-sdk/x/nft/keeper"
// 	//"cosmossdk.io/x/nft"
// 	//"cosmossdk.io/x/nft/client/cli"
// 	//"cosmossdk.io/x/nft/keeper"
// 	//"cosmossdk.io/x/nft/simulation"
// )

// var (
// 	_ module.AppModule      = AppModule{}
// 	_ module.AppModuleBasic = AppModuleBasic{}
// 	//_ module.AppModuleSimulation = AppModule{}
// )

// // AppModuleBasic defines the basic application module used by the nft module.
// type AppModuleBasic struct {
// 	cdc codec.Codec
// }

// // Name returns the nft module's name.
// func (AppModuleBasic) Name() string {
// 	return types.ModuleName
// }

// // RegisterServices registers a gRPC query service to respond to the
// // module-specific gRPC queries.
// func (am AppModule) RegisterServices(cfg module.Configurator) {
// 	types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
// 	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
// }

// // RegisterLegacyAminoCodec registers the nft module's types for the given codec.
// func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// // RegisterInterfaces registers the nft module's interface types
// func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
// 	nft.RegisterInterfaces(registry)
// }

// // DefaultGenesis returns default genesis state as raw bytes for the nft
// // module.
// func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
// 	return cdc.MustMarshalJSON(nft.DefaultGenesisState())
// }

// // ValidateGenesis performs genesis state validation for the nft module.
// func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config sdkclient.TxEncodingConfig, bz json.RawMessage) error {
// 	var data nft.GenesisState
// 	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
// 		return sdkerrors.Wrapf(err, "failed to unmarshal %s genesis state", nft.ModuleName)
// 	}

// 	return nft.ValidateGenesis(data)
// }

// // RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the nft module.
// func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *gwruntime.ServeMux) {
// 	if err := nft.RegisterQueryHandlerClient(context.Background(), mux, nft.NewQueryClient(clientCtx)); err != nil {
// 		panic(err)
// 	}
// }

// // GetQueryCmd returns the cli query commands for the nft module
// func (AppModuleBasic) GetQueryCmd() *cobra.Command {
// 	return cli.GetQueryCmd()
// }

// // GetTxCmd returns the transaction commands for the nft module
// func (AppModuleBasic) GetTxCmd() *cobra.Command {
// 	return cli.GetTxCmd()
// }

// // AppModule implements the sdk.AppModule interface
// type AppModule struct {
// 	AppModuleBasic
// 	keeper keeper.Keeper
// 	// TODO accountKeeper,bankKeeper will be replaced by query service
// 	//sKeeper  keeper.Keeper
// 	//registry cdctypes.InterfaceRegistry
// }

// // NewAppModule creates a new AppModule object
// func NewAppModule(cdc codec.Codec, keeper keeper.Keeper) AppModule {
// 	return AppModule{
// 		AppModuleBasic: AppModuleBasic{cdc: cdc},
// 		keeper:         keeper,
// 		//	sKeeper:        ak,
// 		//	registry: registry,
// 	}
// }

// var _ appmodule.AppModule = AppModule{}

// // IsOnePerModuleType implements the depinject.OnePerModuleType interface.
// func (am AppModule) IsOnePerModuleType() {}

// // IsAppModule implements the appmodule.AppModule interface.
// func (am AppModule) IsAppModule() {}

// // Name returns the nft module's name.
// func (AppModule) Name() string {
// 	return nft.ModuleName
// }

// // RegisterInvariants does nothing, there are no invariants to enforce
// func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// func (am AppModule) NewHandler() sdk.Handler {
// 	return nil
// }

// // InitGenesis performs genesis initialization for the nft module. It returns
// // no validator updates.
// func (am AppModule) InitGenesis(sdk.Context, codec.JSONCodec, json.RawMessage) []ab.ValidatorUpdate {
// 	return nil
// }

// // ExportGenesis returns the exported genesis state as raw bytes for the nft
// // module.
// func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
// 	return nil
// }

// // ConsensusVersion implements AppModule/ConsensusVersion.
// func (AppModule) ConsensusVersion() uint64 { return 1 }

// // ____________________________________________________________________________
