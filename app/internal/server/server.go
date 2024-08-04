package server

import (
	"net"

	userApi "github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userApi.UnimplementedUserApiServer
}

func InitServer(port string) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, nil, err
	}
	s := grpc.NewServer()
	reflection.Register(s)
	userApi.RegisterUserApiServer(s, server{})

	return s, lis, nil
}
