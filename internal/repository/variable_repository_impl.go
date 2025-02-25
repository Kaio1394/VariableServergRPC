package repository

import (
	"VariableServergRPC/internal/dto"
	"VariableServergRPC/internal/model"
	"VariableServergRPC/pb"
	"context"
	"gorm.io/gorm"
)

type VariableRepositoryImpl struct {
	db *gorm.DB
}

func NewVariableRepositoryImpl(db *gorm.DB) *VariableRepositoryImpl {
	return &VariableRepositoryImpl{db: db}
}

func (v *VariableRepositoryImpl) CreateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	variableAux, err := dto.VariableProtoToVariableModel(variable)
	if err != nil {
		return nil, err
	}
	if err := v.db.WithContext(ctx).Create(variableAux).Error; err != nil {
		return nil, err
	}
	return variable, nil
}

func (v *VariableRepositoryImpl) DeleteVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	varModel, err := dto.VariableProtoToVariableModel(variable)
	if err != nil {
		return nil, err
	}
	if err := v.db.WithContext(ctx).First(varModel).Error; err != nil {
		return nil, err
	}
	if err := v.db.WithContext(ctx).Delete(varModel).Error; err != nil {
		return nil, err
	}
	return variable, nil
}

func (v *VariableRepositoryImpl) GetVariable(ctx context.Context, variable *pb.VariableRequest) (*pb.Variable, error) {
	var varExists model.Variable
	if err := v.db.WithContext(ctx).First(&varExists, "Key = ?", variable.VariableKey).Error; err != nil {
		return nil, err
	}
	varProto, err := dto.VariableModelToVariableProto(varExists)
	if err != nil {
		return nil, err
	}
	return varProto, nil
}

func (v *VariableRepositoryImpl) GetVariables(ctx context.Context, empty *pb.Empty) (*pb.VariablesList, error) {
	var variables []model.Variable
	err := v.db.WithContext(ctx).Find(&variables).Error
	if err != nil {
		return &pb.VariablesList{}, err
	}
	variablesAux, err := dto.GetVariablesToVariableList(variables)
	return &variablesAux, nil
}

func (v *VariableRepositoryImpl) UpdateVariable(ctx context.Context, variable *pb.Variable) (*pb.Variable, error) {
	var varExists model.Variable
	if err := v.db.WithContext(ctx).First(&varExists, "key = ?", variable.VariableKey).Error; err != nil {
		return nil, err
	}
	varExists.Key = variable.VariableKey
	varExists.Value = variable.VariableValue
	if err := v.db.WithContext(ctx).Save(&varExists).Error; err != nil {
		return nil, err
	}
	variable, err := dto.VariableModelToVariableProto(varExists)
	if err != nil {
		return nil, err
	}
	return variable, nil
}
