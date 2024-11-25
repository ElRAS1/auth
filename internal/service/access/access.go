package access

import (
	"context"

	accessModel "github.com/ELRAS1/auth/internal/models/access/model"
	"github.com/ELRAS1/auth/internal/repository"
)

type service struct {
	accessRepo repository.Access
}

func New(accessRepo repository.Access) *service {
	return &service{
		accessRepo: accessRepo,
	}
}

func (s *service) Check(context.Context, *accessModel.CheckRequest) error {
	return nil
}
