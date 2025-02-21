package repository

import (
	"VariableServergRPC/pb"
	"context"
)

type VariableRepository interface {
	CreateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error)
	DeleteVariable(ctx context.Context, variable *pb.Variable) error
	UpdateVariable(ctx context.Context, variable *pb.Variable) error
	GetVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error)
	GetVariables(ctx context.Context) (*pb.VariablesList, error)
}
