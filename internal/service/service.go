package service

import (
	"context"

	accessModel "github.com/ELRAS1/auth/internal/models/access/model"
	authModel "github.com/ELRAS1/auth/internal/models/auth/model"
	userModel "github.com/ELRAS1/auth/internal/models/user/model"
)

type User interface {
	Create(ctx context.Context, req *userModel.CreateRequest) (*userModel.CreateResponse, error)
	Update(ctx context.Context, req *userModel.UpdateRequest) error
	Delete(ctx context.Context, req *userModel.DeleteRequest) error
	Get(ctx context.Context, req *userModel.GetRequest) (*userModel.GetResponse, error)
}

type Access interface {
	Check(context.Context, *accessModel.CheckRequest) error
}

type Auth interface {
	Login(ctx context.Context, req *authModel.LoginRequest) (*authModel.LoginResponse, error)
	GetRefreshToken(ctx context.Context, req *authModel.GetRefreshTokenRequest) (*authModel.GetRefreshTokenResponse, error)
	GetAccessToken(ctx context.Context, req *authModel.GetAccessTokenRequest) (*authModel.GetAccessTokenResponse, error)
}
