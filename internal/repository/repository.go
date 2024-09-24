package repository

import (
	"context"

	"github.com/ELRAS1/auth/internal/model"
)

type AuthRepository interface {
	SaveUser(ctx context.Context, req *model.CreateRequest) (int64, error)
	DeleteUser(ctx context.Context, req *model.DeleteRequest) error
	GetUsers(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error)
	UpdateUser(ctx context.Context, req *model.UpdateRequest) error
}
