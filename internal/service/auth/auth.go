package auth

import (
	"context"
	"github.com/ELRAS1/auth/internal/models/auth/model"
	"github.com/ELRAS1/auth/internal/repository"
)

type service struct {
	authRepo repository.Auth
}

func New(authRepo repository.Auth) *service {
	return &service{
		authRepo: authRepo,
	}
}

func (s *service) Login(ctx context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	return &model.LoginResponse{}, nil
}
func (s *service) GetRefreshToken(ctx context.Context, req *model.GetRefreshTokenRequest) (*model.GetRefreshTokenResponse, error) {
	return &model.GetRefreshTokenResponse{}, nil
}
func (s *service) GetAccessToken(ctx context.Context, req *model.GetAccessTokenRequest) (*model.GetAccessTokenResponse, error) {
	return &model.GetAccessTokenResponse{}, nil
}
