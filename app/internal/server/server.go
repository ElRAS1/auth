package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"

	userApi "github.com/ELRAS1/auth/pkg/userApi"
	"github.com/ELRAS1/auth/app/internal/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	userApi.UnimplementedUserApiServer
}

type AppServer struct {
	*server
	db     *storage.Storage
	logger *slog.Logger
}

func newApp(logger slog.Logger) *AppServer {
	s := &server{}
	return &AppServer{
		server: s,
		db:     &storage.Storage{},
		logger: &logger,
	}
}

func StartServer(ctx context.Context, logger *slog.Logger, port string) (*grpc.Server, net.Listener, error) {
	const nm = "[StartServer]"
	var err error

	app := newApp(*logger)
	app.db, err = storage.ConfigureStorage(ctx)
	if err != nil {
		logger.Error(err.Error())
		return nil, nil, fmt.Errorf("%s %v", nm, err)
	}

	logger.Info("the database is running")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, nil, fmt.Errorf("%s %v", nm, err)
	}
	
	srv := grpc.NewServer()
	reflection.Register(srv)
	userApi.RegisterUserApiServer(srv, app)

	return srv, lis, nil
}
