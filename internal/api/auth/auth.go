package auth

import (
	"context"
	"github.com/ELRAS1/auth/internal/converters/auth/converter"
	"github.com/ELRAS1/auth/internal/service"
	"github.com/ELRAS1/auth/pkg/auth"
)

type Api struct {
	*auth.UnimplementedAuthServer
	serv service.Auth
}

func New(serv service.Auth) *Api {
	return &Api{
		serv:                    serv,
		UnimplementedAuthServer: &auth.UnimplementedAuthServer{},
	}
}

func (a *Api) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	resp, err := a.serv.Login(ctx, converter.LoginToModel(req))
	if err != nil {
		return nil, err
	}

	return converter.LoginToApi(resp), nil
}
func (a *Api) GetRefreshToken(
	ctx context.Context, req *auth.GetRefreshTokenRequest) (*auth.GetRefreshTokenResponse, error) {
	resp, err := a.serv.GetRefreshToken(ctx, converter.GetRefreshTokenToModel(req))
	if err != nil {
		return nil, err
	}

	return converter.GetRefreshTokenToApi(resp), nil
}
func (a *Api) GetAccessToken(ctx context.Context, req *auth.GetAccessTokenRequest) (*auth.GetAccessTokenResponse, error) {
	resp, err := a.serv.GetAccessToken(ctx, converter.GetAccessTokenToModel(req))
	if err != nil {
		return nil, err
	}

	return converter.GetAccessTokenToApi(resp), nil
}
