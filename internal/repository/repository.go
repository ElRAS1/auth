package repository

import (
	"context"

	"github.com/ELRAS1/auth/internal/models/user/model"
)

type User interface {
	Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error)
	Update(ctx context.Context, req *model.UpdateRequest) error
	Delete(ctx context.Context, req *model.DeleteRequest) error
	Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error)
}

type Auth interface {
}

type Access interface {
}
