package service

import (
	"VariableServergRPC/pb"
	"context"
)

type VariableService interface {
	CreateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error)
	DeleteVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error)
	GetVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error)
	GetVariables(ctx context.Context, empty *pb.Empty) (*pb.VariablesList, error)
	UpdateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error)
}
