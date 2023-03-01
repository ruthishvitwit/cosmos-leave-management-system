package keeper

import (
	"context"
	"coslms/x/std/types"
	//"github/ruthishvitwit/coslms/x/std/types"
)

var _ types.QueryServer = Keeper{}

// type Queryserver struct {
// 	Keeper
// 	types.UnimplementedQueryServer
// }

// func NewQueryServerImpl(k Keeper) types.QueryServer {
// 	return &Queryserver{
// 		Keeper: k,
// 	}
// }

func (k Keeper) QueryGetStudent(context.Context, *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	return &types.GetStudentsResponse{}, nil
}
func (k Keeper) QueryGetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	return &types.GetLeaveRequestsResponse{}, nil
}

func (k Keeper) QueryGetLeaveApprovedRequests(context.Context, *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	return &types.GetLeaveApprovedRequestsResponse{}, nil
}

func (k Keeper) QueryGetLeaveStatus(context.Context, *types.GetLeaveStatusRequest) (*types.GetLeaveStatusResponse, error) {
	return &types.GetLeaveStatusResponse{}, nil
}
