package api

import (
	"context"
	"fmt"

	"github.com/ELRAS1/auth/internal/converter"
	"github.com/ELRAS1/auth/internal/service"
	"github.com/ELRAS1/auth/internal/validations"
	"github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/protobuf/types/known/emptypb"
)

type AuthApi struct {
	*userApi.UnimplementedUserApiServer
	serv service.AuthService
}

func New(srv service.AuthService) *AuthApi {
	return &AuthApi{
		serv:                       srv,
		UnimplementedUserApiServer: &userApi.UnimplementedUserApiServer{},
	}
}

func (a *AuthApi) Create(ctx context.Context, req *userApi.CreateRequest) (*userApi.CreateResponse, error) {
	if err := validations.CheckCreate(converter.CreateToModel(req)); err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}

	resp, err := a.serv.Create(ctx, converter.CreateToModel(req))
	if err != nil {
		return nil, err
	}

	return converter.CreateToApi(resp), nil
}

func (a *AuthApi) Update(ctx context.Context, req *userApi.UpdateRequest) (*emptypb.Empty, error) {
	if err := validations.CheckUpdate(converter.UpdateToModel(req)); err != nil {
		return &emptypb.Empty{}, fmt.Errorf("update error: %w", err)
	}

	err := a.serv.Update(ctx, converter.UpdateToModel(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (a *AuthApi) Delete(ctx context.Context, req *userApi.DeleteRequest) (*emptypb.Empty, error) {
	err := a.serv.Delete(ctx, converter.DeleteToModel(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (a *AuthApi) Get(ctx context.Context, req *userApi.GetRequest) (*userApi.GetResponse, error) {
	resp, err := a.serv.Get(ctx, converter.GetToModel(req))
	if err != nil {
		return nil, err
	}

	return converter.GetToApi(resp), nil
}
