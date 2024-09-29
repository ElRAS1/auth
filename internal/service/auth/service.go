package auth

import (
	"context"
	"fmt"

	"github.com/ELRAS1/auth/internal/model"
	"github.com/ELRAS1/auth/internal/repository"
	"github.com/ELRAS1/auth/internal/validations"
)

type service struct {
	authRepository repository.AuthRepository
}

func New(authRepo repository.AuthRepository) *service {
	return &service{authRepository: authRepo}
}

func (s *service) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	if err := validations.CheckCreate(req); err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}

	resp, err := s.authRepository.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, req *model.UpdateRequest) error {
	if err := validations.CheckUpdate(req); err != nil {
		return fmt.Errorf("update error: %w", err)
	}

	err := s.authRepository.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("update error: %w", err)
	}

	return nil
}

func (s *service) Delete(ctx context.Context, req *model.DeleteRequest) error {
	err := s.authRepository.Delete(ctx, req)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	return nil
}

func (s *service) Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error) {
	res, err := s.authRepository.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("get error: %w", err)
	}

	return res, nil
}
