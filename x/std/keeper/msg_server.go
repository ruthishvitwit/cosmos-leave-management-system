package keeper

import (
	"context"
	"coslms/x/std/types"
	//"github/ruthishvitwit/coslms/x/std/types"
)

var _ types.MsgServer = Msgserver{}

type Msgserver struct {
	Keeper
	types.UnimplementedMsgServer
}

func NewMsgServerImpl(k Keeper) types.MsgServer {
	return &Msgserver{
		Keeper: k,
	}
}

func (k Msgserver) RegisterAdminRequest(context.Context, *types.RegisterAdminRequest) (*types.RegisterAdminRequest, error) {
	return &types.RegisterAdminRequest{}, nil
}

func (k Msgserver) AddStudent(context.Context, *types.AddStudentRequest) (*types.AddStudentResponse, error) {
	return &types.AddStudentResponse{}, nil
}

func (k Msgserver) ApplyLeaveRequest(context.Context, *types.ApplyLeaveRequest) (*types.ApplyLeaveRequest, error) {
	return &types.ApplyLeaveRequest{}, nil
}
