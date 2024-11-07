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

	"github.com/ELRAS1/auth/internal/api"
	"github.com/ELRAS1/auth/internal/config"
	repoAuth "github.com/ELRAS1/auth/internal/repository/auth"
	serviceAuth "github.com/ELRAS1/auth/internal/service/auth"
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

	repo := repoAuth.New(dbClient)
	service := serviceAuth.New(repo)

	server := grpc.NewServer()
	reflection.Register(server)
	userApi.RegisterUserApiServer(server, api.New(service))

	go func() {
		logger.Info(fmt.Sprintf("grpc server is running on the port %v", cfg.GRPCPort))
		if err = server.Serve(listener); err != nil {
			log.Fatalln(fmt.Sprintf("failed to grpc serve: %v", err))
		}
	}()

	httpServer := config.InitHTTP(ctx, cfg.GRPCPort, cfg.HTTPPort)
	go func() {
		logger.Info(fmt.Sprintf("http server is running on the port %v", cfg.HTTPPort))
		if err = httpServer.ListenAndServe(); err != nil {
			log.Fatalln(fmt.Sprintf("failed to http serve: %v", err))
		}
	}()

	httpSwagger := config.InitSwagger()
	go func() {
		logger.Info(fmt.Sprintf("swagger ui is running on the url %v", "http://localhost:8090"))
		if err = http.ListenAndServe(cfg.HTTPSwagger, httpSwagger); err != nil {
			log.Fatalln(fmt.Sprintf("failed to swagger serve: %v", err))
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
