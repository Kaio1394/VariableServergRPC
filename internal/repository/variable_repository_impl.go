package repository

import (
	"VariableServergRPC/pb"
	"context"
	"gorm.io/gorm"
)

type VariableRepositoryImpl struct {
	db *gorm.DB
}

func (v *VariableRepositoryImpl) CreateVariable(ctx context.Context, variable *pb.Variable) error {
	if err := v.db.WithContext(ctx).Create(variable).Error; err != nil {
		return err
	}
	return nil
}

func (v *VariableRepositoryImpl) DeleteVariable(ctx context.Context, variable *pb.Variable) error {
	//TODO implement me
	panic("implement me")
}

func (v *VariableRepositoryImpl) UpdateVariable(ctx context.Context, variable *pb.Variable) error {
	//TODO implement me
	panic("implement me")
}

func (v *VariableRepositoryImpl) GetVariable(ctx context.Context, variable *pb.Variable) (pb.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VariableRepositoryImpl) GetVariables(ctx context.Context) (pb.VariablesList, error) {
	//TODO implement me
	panic("implement me")
}
