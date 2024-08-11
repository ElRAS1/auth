package server

import (
	"context"

	"github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a AppServer) Get(ctx context.Context, req *userApi.Request) (*userApi.GetResponse, error) {
	return nil, nil
}

func (a AppServer) Create(ctx context.Context, req *userApi.Request) (*userApi.CreateResponse, error) {
	err := a.validations(req)
	if err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}
	id, err := a.db.SaveUser(req)
	if err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}
	return &userApi.CreateResponse{Id: id}, nil
}

func (a AppServer) Update(ctx context.Context, req *userApi.UpdateRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (a AppServer) Delete(ctx context.Context, req *userApi.DeleteRequest) (*emptypb.Empty, error) {
	err := a.db.DeleteUser(req)
	if err != nil {
		a.logger.Info(err.Error())
		return nil, err
	}
	return nil, nil
}
