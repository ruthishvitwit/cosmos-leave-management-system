package keeper

import (
	"context"
	"coslms/x/std/types"
	//"github/ruthishvitwit/coslms/x/std/types"
)

var _ types.QueryServer = &Queryserver{}

type Queryserver struct {
	Keeper
	types.UnimplementedQueryServer
}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return &Queryserver{
		Keeper: k,
	}
}

func (k Queryserver) GetStudent(context.Context, *types.GetStudentsRequest) (*types.GetStudentsResponse, error) {
	return &types.GetStudentsResponse{}, nil
}
func (k Queryserver) GetLeaveRequests(context.Context, *types.GetLeaveRequestsRequest) (*types.GetLeaveRequestsResponse, error) {
	return &types.GetLeaveRequestsResponse{}, nil
}

func (k Queryserver) GetLeaveApprovedRequests(context.Context, *types.GetLeaveApprovedRequestsRequest) (*types.GetLeaveApprovedRequestsResponse, error) {
	return &types.GetLeaveApprovedRequestsResponse{}, nil
}
