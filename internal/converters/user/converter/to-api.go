package converter

import (
	"github.com/ELRAS1/auth/internal/models/user/model"
	"github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetToApi(req *model.GetResponse) *userApi.GetResponse {
	var updatedAt *timestamppb.Timestamp

	if req.UpdatedAt.Valid {
		updatedAt = timestamppb.New(req.UpdatedAt.Time)
	}

	return &userApi.GetResponse{
		Id:        req.Id,
		Name:      req.Name,
		Email:     req.Email,
		Role:      userApi.Role(req.Role),
		CreatedAt: timestamppb.New(req.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func CreateToApi(req *model.CreateResponse) *userApi.CreateResponse {
	return &userApi.CreateResponse{
		Id: req.Id,
	}
}
