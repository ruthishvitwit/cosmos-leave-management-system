package keeper

import (
	"context"
	"coslms/x/std/types"

	//"proto/cosmos/x/std/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

// type msgServer struct {
// 	Keeper
// 	types.UnimplementedMsgServer
// }

// func NewMsgServerImpl(k Keeper) types.MsgServer {
// 	return &msgServer{
// 		Keeper: k,
// 	}
// }

func (k Keeper) MsgAddStudent(ctx context.Context, students *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	k.AddStudent(sdkCtx, students)
	return &types.AddStudentResponse{}, nil
}

func (k Keeper) MsgApplyLeave(context.Context, *types.ApplyLeaveRequest) (*types.ApplyLeaveResponse, error) {
	return &types.ApplyLeaveResponse{}, nil
}

func (k Keeper) MsgRegisterAdmin(context.Context, *types.RegisterAdminRequest) (*types.RegisterAdminResponse, error) {
	return &types.RegisterAdminResponse{}, nil
}

func (k Keeper) MsgAcceptLeave(context.Context, *types.AcceptLeaveRequest) (*types.AcceptLeaveResponse, error) {
	return &types.AcceptLeaveResponse{}, nil
}
