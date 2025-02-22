package dto

import (
	"VariableServergRPC/internal/model"
	"VariableServergRPC/pb"
)

func VariableModelToVariableProto(model model.Variable) (*pb.Variable, error) {
	var varProto pb.Variable
	varProto.VariableKey = model.Key
	varProto.VariableValue = model.Value
	return &varProto, nil
}

func VariableProtoToVariableModel(variable *pb.Variable) (*model.Variable, error) {
	var varModel model.Variable
	varModel.Key = variable.VariableKey
	varModel.Value = variable.VariableValue
	return &varModel, nil
}

func GetVariablesToVariableList(variables []model.Variable) (pb.VariablesList, error) {
	var variablesList = []*pb.Variable{}
	for _, v := range variables {
		variablesList = append(variablesList, &pb.Variable{
			VariableKey:   v.Key,
			VariableValue: v.Value,
		})
	}
	return pb.VariablesList{Variables: variablesList}, nil
}
