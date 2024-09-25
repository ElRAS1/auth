package service

import (
	"context"

	"github.com/ELRAS1/auth/internal/model"
)

type AuthService interface {
	Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error)
	Update(ctx context.Context, req *model.UpdateRequest) error
	Delete(ctx context.Context, req *model.DeleteRequest) error
	Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error)
}
