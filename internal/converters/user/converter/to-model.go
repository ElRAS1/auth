package converter

import (
	"github.com/ELRAS1/auth/internal/models/user/model"
	"github.com/ELRAS1/auth/pkg/userApi"
)

func CreateToModel(req *userApi.CreateRequest) *model.CreateRequest {
	return &model.CreateRequest{
		Name:            req.GetName(),
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
		Role:            model.Role(req.GetRole()),
	}
}

func UpdateToModel(req *userApi.UpdateRequest) *model.UpdateRequest {
	return &model.UpdateRequest{
		Id:    req.GetId(),
		Name:  req.GetName().GetValue(),
		Email: req.GetEmail().GetValue(),
	}
}

func DeleteToModel(req *userApi.DeleteRequest) *model.DeleteRequest {
	return &model.DeleteRequest{
		Id: req.GetId(),
	}
}

func GetToModel(req *userApi.GetRequest) *model.GetRequest {
	return &model.GetRequest{
		Id: req.GetId(),
	}
}
