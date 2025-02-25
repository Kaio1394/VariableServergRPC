package handler

import (
	"VariableServergRPC/pb"
	"context"
)

type VariableHandler interface {
	CreateVariable(context.Context, *pb.Variable) (*pb.Variable, error)
	GetVariables(context.Context, *pb.Empty) (*pb.VariablesList, error)
	GetVariable(context.Context, *pb.Variable) (*pb.Variable, error)
	UpdateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error)
	DeleteVariable(context.Context, *pb.Variable) (*pb.Variable, error)
}
