package main

import (
	"os"
	"os/signal"

	"github.com/jaeecheveste/r2-back/pkg/app"
	"github.com/jaeecheveste/r2-back/pkg/config"

	"syscall"

	handler "github.com/jaeecheveste/r2-back/pkg/http/handlers"
	"github.com/jaeecheveste/r2-back/pkg/http/server"
	"github.com/jaeecheveste/r2-back/pkg/logger"
)

type exitCode int

func main() {
	log := logger.NewLogger()
	log.Info("Starting R2 backend service")
	defer logger.LogSync(log)

	// load env configuration.
	config, err := config.LoadConfiguration()
	if err != nil {
		log.Errorf("LoadCondiguration: %v", err)
	}

	// init services
	fibonacciService := app.NewFibonacciService(log)

	// init server handlers
	fh := handler.NewFibonacciHandler(fibonacciService)

	s := server.NewServer(config.Server, fh)

	go func() {
		if err := s.Start(); err != nil {
			log.Error(err)
			panic(exitCode(1))
		}
	}()

	// Graceful shutdown
	sigQuit := make(chan os.Signal, 1)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	sig := <-sigQuit

	log.Infof("Shutting down server with signal [%v] ...", sig)
	if err = s.Stop(); err != nil {
		log.Error(err)
		panic(exitCode(1))
	}
	log.Info("Ternminating server!")
}
