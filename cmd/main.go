package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ELRAS1/auth/app/internal/config"
	"github.com/ELRAS1/auth/app/internal/server"
	"github.com/ELRAS1/auth/pkg/logger"
)

func main() {
	cfg, err := config.ReadingConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	logger := logger.ConfigureLogger(cfg.Loglevel, cfg.Configlog)

	srv, lis, err := server.StartServer(ctx, logger, cfg.Port)
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err = srv.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()

	logger.Info(fmt.Sprintf("the server is running on the port %v", cfg.Port))
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	cancel()
}
