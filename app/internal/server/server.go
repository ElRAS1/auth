package server

import (
	"log/slog"
	"net"

	userApi "github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userApi.UnimplementedUserApiServer
	logger *slog.Logger
}

func newServer(log *slog.Logger) *server {
	return &server{
		logger: log,
	}
}

func InitServer(port string, logger *slog.Logger) (*grpc.Server, net.Listener, error) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, nil, err
	}
	s := grpc.NewServer()
	reflection.Register(s)
	userApi.RegisterUserApiServer(s, newServer(logger))

	return s, lis, nil
}
