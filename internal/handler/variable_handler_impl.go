package handler

import (
	"VariableServergRPC/internal/service"
	"VariableServergRPC/pb"
	"context"
)

type VariableHandlerImpl struct {
	pb.UnimplementedVariableServiceServer
	s *service.VariableServiceImpl
}

func (v VariableHandlerImpl) CreateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (v VariableHandlerImpl) GetVariables(ctx context.Context, empty *pb.Empty) (*pb.VariablesList, error) {
	//TODO implement me
	panic("implement me")
}

func (v VariableHandlerImpl) GetVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (v VariableHandlerImpl) UpdateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (v VariableHandlerImpl) DeleteVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func NewVariableHandlerImpl(service *service.VariableServiceImpl) *VariableHandlerImpl {
	return &VariableHandlerImpl{s: service}
}
