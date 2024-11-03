package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
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
		if err = server.Serve(listener); err != nil {
			log.Fatalln(fmt.Sprintf("failed to grpc serve: %v", err))
		}
	}()
	logger.Info(fmt.Sprintf("grpc server is running on the port %v", cfg.GRPCPort))

	httpServer := InitHTTP(ctx, cfg.GRPCPort, cfg.HTTPPort)
	go func() {
		if err = httpServer.ListenAndServe(); err != nil {
			log.Fatalln(fmt.Sprintf("failed to http serve: %v", err))
		}
	}()
	logger.Info(fmt.Sprintf("http server is running on the port %v", cfg.HTTPPort))

	go func() {
		if err = http.ListenAndServe(cfg.HTTPSwagger, InitSwagger()); err != nil {
			log.Fatalln(fmt.Sprintf("failed to swagger serve: %v", err))
		}
	}()
	logger.Info(fmt.Sprintf("swagger ui is running on the url %v", "http://localhost:8090"))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}

func InitHTTP(ctx context.Context, GRPCport, HTTPport string) *http.Server {
	mux := runtime.NewServeMux()
	if err := userApi.RegisterUserApiHandlerFromEndpoint(ctx, mux, GRPCport, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}); err != nil {
		log.Fatalln(err)
	}

	return &http.Server{
		Handler: CorsMiddleware(mux),
		Addr:    HTTPport,
	}
}

func InitSwagger() *http.ServeMux {
	swaggerHTTP := http.NewServeMux()
	swaggerHTTP.HandleFunc("/api.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/swagger/api.swagger.json")
	})

	swaggerHTTP.Handle("/", http.FileServer(http.Dir("./swagger-ui/")))

	return swaggerHTTP
}

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, UPDATE, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	})
}
