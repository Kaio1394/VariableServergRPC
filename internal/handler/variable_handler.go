package handler

import (
	"VariableServergRPC/pb"
	"context"
)

type VariableHandler interface {
	CreateVariable(ctx context.Context, req *pb.Variable) error
	DeleteVariable(ctx context.Context, req *pb.Variable) error
	UpdateVariable(ctx context.Context, req *pb.Variable) error
	GetVariable(ctx context.Context, req *pb.Variable) (pb.Variable, error)
	GetVariables(ctx context.Context) (pb.VariablesList, error)
}
