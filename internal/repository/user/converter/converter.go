package converter

import (
	"database/sql"

	"github.com/ELRAS1/auth/internal/models/user/model"
	modulRepo "github.com/ELRAS1/auth/internal/repository/user/model"
)

func RepoGetToModel(resp *modulRepo.GetResponse) *model.GetResponse {
	var updatedAt sql.NullTime

	if resp.UpdatedAt.Valid {
		updatedAt = resp.UpdatedAt
	}

	return &model.GetResponse{
		Id:        resp.Id,
		Name:      resp.Name,
		Email:     resp.Email,
		Role:      model.Role(resp.Role),
		CreatedAt: resp.CreatedAt,
		UpdatedAt: updatedAt,
	}
}

func RepoCreateToModel(resp *modulRepo.CreateResponse) *model.CreateResponse {
	return &model.CreateResponse{
		Id: resp.Id,
	}
}
