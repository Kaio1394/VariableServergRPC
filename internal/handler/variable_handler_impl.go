package handler

import (
	"VariableServergRPC/pb"
	"context"
)

type VariableHandlerImpl struct{}

func (v VariableHandlerImpl) CreateVariable(ctx context.Context, req *pb.Variable) error {
	//TODO implement me
	panic("implement me")
}

func (v VariableHandlerImpl) DeleteVariable(ctx context.Context, req *pb.Variable) error {
	//TODO implement me
	panic("implement me")
}

func (v VariableHandlerImpl) UpdateVariable(ctx context.Context, req *pb.Variable) error {
	//TODO implement me
	panic("implement me")
}

func (v VariableHandlerImpl) GetVariable(ctx context.Context, req *pb.Variable) (pb.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (v VariableHandlerImpl) GetVariables(ctx context.Context) (pb.VariablesList, error) {
	//TODO implement me
	panic("implement me")
}
