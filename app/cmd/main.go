package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ELRAS1/auth/app/config"
	"github.com/ELRAS1/auth/app/internal/logger"
	"github.com/ELRAS1/auth/app/internal/server"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}
	logger := logger.ConfigureLogger(cfg.Loglevel, cfg.Configlog)
	logger.Debug("the logger has been initialized")
	srv, lis, err := server.InitServer(cfg.Port)
	go func() {
		if err = srv.Serve(lis); err != nil {
			log.Fatalln(err)
		}
	}()
	logger.Info(fmt.Sprintf("the server is running on the port %v", cfg.Port))
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
}
