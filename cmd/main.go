package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	userImpl "github.com/ELRAS1/auth/internal/api/user"

	"github.com/ELRAS1/auth/internal/config"
	userRepo "github.com/ELRAS1/auth/internal/repository/user"
	userService "github.com/ELRAS1/auth/internal/service/user"
	lgr "github.com/ELRAS1/auth/pkg/logger"
	"github.com/ELRAS1/auth/pkg/userApi"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	logger := lgr.New(cfg.LogLevel, cfg.ConfigLog)
	slog.SetDefault(logger)

	listener, err := net.Listen(cfg.Network, cfg.GRPCPort)
	if err != nil {
		log.Fatalln(err)
	}

	dbClient, err := config.InitializeDatabaseClient(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	repo := userRepo.New(dbClient)
	service := userService.New(repo)

	server := grpc.NewServer()
	reflection.Register(server)
	userApi.RegisterUserApiServer(server, userImpl.New(service))

	go func() {
		logger.Info(fmt.Sprintf("grpc server is running on: %v", net.JoinHostPort(cfg.Host, cfg.GRPCPort)))
		if err = server.Serve(listener); err != nil {
			log.Fatalln(fmt.Sprintf("failed to grpc serve: %v", err))
		}
	}()

	httpServer := config.InitHTTP(ctx, cfg.GRPCPort, cfg.HTTPPort)
	go func() {
		logger.Info(fmt.Sprintf("http server is running on: %v", net.JoinHostPort(cfg.Host, cfg.HTTPPort)))
		if err = httpServer.ListenAndServe(); err != nil {
			log.Fatalln(fmt.Sprintf("failed to http serve: %v", err))
		}
	}()

	httpSwagger := config.InitSwagger()
	go func() {
		logger.Info(fmt.Sprintf("swagger ui is running on: %v", net.JoinHostPort(cfg.Host, cfg.SwaggerPort)))
		if err = http.ListenAndServe(cfg.SwaggerPort, httpSwagger); err != nil {
			log.Fatalln(fmt.Sprintf("failed to swagger serve: %v", err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
