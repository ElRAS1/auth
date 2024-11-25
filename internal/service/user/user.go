package user

import (
	"context"
	"fmt"

	"github.com/ELRAS1/auth/internal/models/user/model"
	"github.com/ELRAS1/auth/internal/repository"
)

type service struct {
	userRepo repository.User
}

func New(userRepo repository.User) *service {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Create(ctx context.Context, req *model.CreateRequest) (*model.CreateResponse, error) {
	resp, err := s.userRepo.Create(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, req *model.UpdateRequest) error {
	err := s.userRepo.Update(ctx, req)
	if err != nil {
		return fmt.Errorf("update error: %w", err)
	}

	return nil
}

func (s *service) Delete(ctx context.Context, req *model.DeleteRequest) error {
	err := s.userRepo.Delete(ctx, req)
	if err != nil {
		return fmt.Errorf("delete error: %w", err)
	}

	return nil
}

func (s *service) Get(ctx context.Context, req *model.GetRequest) (*model.GetResponse, error) {
	res, err := s.userRepo.Get(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("get error: %w", err)
	}

	return res, nil
}
