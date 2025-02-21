package dto

import (
	"VariableServergRPC/internal/model"
	"VariableServergRPC/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func VariableModelToVariableProto(model model.Variable) (*pb.Variable, error) {
	var varProto pb.Variable
	varProto.VariableKey = model.Key
	varProto.VariableValue = model.Value
	varProto.EditDate = timestamppb.New(model.EditDate)
	return &varProto, nil
}

func VariableProtoToVariableModel(variable *pb.Variable) (*model.Variable, error) {
	var varModel model.Variable
	varModel.Key = variable.VariableKey
	varModel.Value = variable.VariableValue
	varModel.EditDate = variable.EditDate.AsTime()
	return &varModel, nil
}

func GetVariablesToVariableList(variables []model.Variable) (pb.VariablesList, error) {
	var variablesList = []*pb.Variable{}
	for _, v := range variables {
		variablesList = append(variablesList, &pb.Variable{
			VariableKey:   v.Key,
			VariableValue: v.Value,
			EditDate:      timestamppb.New(v.EditDate),
		})
	}
	return pb.VariablesList{Variables: variablesList}, nil
}
