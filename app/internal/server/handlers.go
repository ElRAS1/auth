package server

import (
	"context"

	"github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s server) Get(ctx context.Context, req *userApi.Request) (*userApi.GetResponse, error) {
	return nil, nil
}

func (s server) Create(ctx context.Context, req *userApi.Request) (*userApi.CreateResponse, error) {
	return nil, nil
}

func (s server) Update(ctx context.Context, req *userApi.UpdateRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s server) Delete(ctx context.Context, req *userApi.DeleteRequest) (*emptypb.Empty, error) {
	return nil, nil
}
