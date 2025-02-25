package service

import (
	"VariableServergRPC/internal/repository"
	"VariableServergRPC/pb"
	"context"
)

type VariableServiceImpl struct {
	vr *repository.VariableRepositoryImpl
}

func NewVariableServiceimpl(repo *repository.VariableRepositoryImpl) *VariableServiceImpl {
	return &VariableServiceImpl{vr: repo}
}

func (vs *VariableServiceImpl) CreateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	return vs.vr.CreateVariable(ctx, variable)
}

func (vs *VariableServiceImpl) DeleteVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	return vs.vr.DeleteVariable(ctx, variable)
}

func (vs *VariableServiceImpl) GetVariables(ctx context.Context, empty *pb.Empty) (*pb.VariablesList, error) {
	return vs.vr.GetVariables(ctx, empty)
}

func (vs *VariableServiceImpl) UpdateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	return vs.vr.UpdateVariable(ctx, variable)
}
