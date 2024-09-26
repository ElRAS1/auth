package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ELRAS1/auth/internal/api"
	"github.com/ELRAS1/auth/internal/config"
	repoAuth "github.com/ELRAS1/auth/internal/repository/auth"
	serviceAuth "github.com/ELRAS1/auth/internal/service/auth"

	"github.com/ELRAS1/auth/pkg/logger"
	"github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.NewServerCfg()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logger.New(cfg.LogLevel, cfg.ConfigLog)

	listener, err := net.Listen(cfg.Network, cfg.Port)
	if err != nil {
		log.Fatalln(err)
	}

	dbClient, err := config.InitializeDatabaseClient(ctx)

	if err != nil {
		log.Fatalln(err)
	}

	repo := repoAuth.New(dbClient)
	service := serviceAuth.New(repo)

	server := grpc.NewServer()
	reflection.Register(server)
	userApi.RegisterUserApiServer(server, api.New(service))

	go func() {
		if err = server.Serve(listener); err != nil {
			log.Fatalln(err)
		}
	}()

	logger.Info(fmt.Sprintf("the server is running on the port %v", cfg.Port))
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

}
