package auth

import (
	"context"

	"github.com/ELRAS1/auth/internal/model"
	"github.com/ELRAS1/auth/internal/repository"
)

type service struct {
	authRepository repository.AuthRepository
}

func New(authRepo repository.AuthRepository) *service {
	return &service{authRepository: authRepo}
}

func (s *service) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	resp, err := s.authRepository.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, req *model.UpdateRequest) error {
	err := s.authRepository.Update(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, req *model.DeleteRequest) error {
	err := s.authRepository.Delete(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error) {
	res, err := s.authRepository.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
