package repository

import (
	"VariableServergRPC/internal/model"
	"VariableServergRPC/pb"
	"context"
	"gorm.io/gorm"
)

type VariableRepositoryImpl struct {
	db *gorm.DB
}

func NewVariableRepositoryImpl(db *gorm.DB) *VariableRepositoryImpl {
	return &VariableRepositoryImpl{}
}

func (v *VariableRepositoryImpl) CreateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	if err := v.db.WithContext(ctx).Create(variable).Error; err != nil {
		return nil, err
	}
	return variable, nil
}

func (v *VariableRepositoryImpl) DeleteVariable(ctx context.Context, variable *pb.Variable) error {
	//TODO implement me
	panic("implement me")
}

func (v *VariableRepositoryImpl) GetVariable(ctx context.Context, variable *pb.Variable) (pb.Variable, error) {
	//TODO implement me
	panic("implement me")
}

func (v *VariableRepositoryImpl) GetVariables(ctx context.Context) (pb.VariablesList, error) {
	panic("implement me")
}

func (v *VariableRepositoryImpl) UpdateVariable(ctx context.Context, variable *pb.Variable) error {
	var varExists model.Variable
	if err := v.db.WithContext(ctx).First(&variable, varExists.ID).Error; err != nil {
		return err
	}
	varExists.Key = variable.VariableKey
	varExists.Value = variable.VariableValue
	varExists.EditDate = variable.EditDate.AsTime()
	return v.db.WithContext(ctx).Save(varExists).Error
}
