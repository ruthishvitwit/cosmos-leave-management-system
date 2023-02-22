package keeper

import (
	"context"
	"coslms/x/std/types"
	//"github/ruthishvitwit/coslms/x/std/types"
)

var _ types.QueryServer = Queryserver{}

type Queryserver struct {
	Keeper
	types.UnimplementedQueryServer
}

func NewQueryServerImpl(k Keeper) types.QueryServer {
	return &Queryserver{
		Keeper: k,
	}
}

func (k Queryserver) GetStudent(context.Context, *types.GetStudentRequest) (*types.GetStudentResponse, error) {
	return &types.GetStudentResponse{}, nil
}
