package access

import (
	"context"
	"github.com/ELRAS1/auth/internal/converters/access/converter"
	"github.com/ELRAS1/auth/internal/service"
	"github.com/ELRAS1/auth/pkg/access"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Api struct {
	*access.UnimplementedAccessServer
	serv service.Access
}

func New(srv service.Access) *Api {
	return &Api{
		serv:                      srv,
		UnimplementedAccessServer: &access.UnimplementedAccessServer{},
	}
}

func (a *Api) Check(ctx context.Context, req *access.CheckRequest) (*emptypb.Empty, error) {
	if err := a.serv.Check(ctx, converter.CheckToModel(req)); err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
