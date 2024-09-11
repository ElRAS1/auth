package server

import (
	"context"

	"github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a AppServer) Get(ctx context.Context, req *userApi.GetRequest) (*userApi.GetResponse, error) {
	var resp *userApi.GetResponse
	var err error

	if resp, err = a.db.GetUsers(ctx, req); err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}

	return resp, nil
}

func (a AppServer) Create(ctx context.Context, req *userApi.CreateRequest) (*userApi.CreateResponse, error) {
	err := a.CreateValidations(req)
	if err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}

	id, err := a.db.SaveUser(ctx, req)
	if err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}

	return &userApi.CreateResponse{Id: id}, nil
}

func (a AppServer) Update(ctx context.Context, req *userApi.UpdateRequest) (*emptypb.Empty, error) {
	err := a.UpdateValidation(req)
	if err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}

	err = a.db.UpdateUser(ctx, req)
	if err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}

	return nil, nil
}

func (a AppServer) Delete(ctx context.Context, req *userApi.DeleteRequest) (*emptypb.Empty, error) {
	err := a.db.DeleteUser(ctx, req)
	if err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}

	return nil, nil
}
