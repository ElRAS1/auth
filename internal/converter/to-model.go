package converter

import (
	"github.com/ELRAS1/auth/internal/model"
	"github.com/ELRAS1/auth/pkg/userApi"
)

func ServiceCreateToModel(req *userApi.CreateRequest) *model.CreateRequest {
	return &model.CreateRequest{
		Name:            req.GetName(),
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
		PasswordConfirm: req.GetPasswordConfirm(),
		Role:            model.Role(req.GetRole()),
	}
}

func ServiceUpdateToModel(req *userApi.UpdateRequest) *model.UpdateRequest {
	return &model.UpdateRequest{
		Id:    req.GetId(),
		Name:  req.GetName().Value,
		Email: req.GetEmail().Value,
	}
}

func ServiceDeleteToModel(req *userApi.DeleteRequest) *model.DeleteRequest {
	return &model.DeleteRequest{
		Id: req.GetId(),
	}
}

func ServiceGetToModel(req *userApi.GetRequest) *model.GetRequest {
	return &model.GetRequest{
		Id: req.GetId(),
	}
}
