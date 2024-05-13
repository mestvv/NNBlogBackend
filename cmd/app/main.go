package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	apiHttp "github.com/mestvv/NNBlogBackend/internal/api/http"
	"github.com/mestvv/NNBlogBackend/internal/config"
	"github.com/mestvv/NNBlogBackend/internal/db"
	"github.com/mestvv/NNBlogBackend/internal/log"
	"github.com/mestvv/NNBlogBackend/internal/repository"
	"github.com/mestvv/NNBlogBackend/internal/server"
	"github.com/mestvv/NNBlogBackend/internal/service"
)

const configPath = "config/config.yaml"

func main() {
	// Init cfg
	cfg := config.MustLoad(configPath)

	// Init logger
	logger := log.SetupLogger(cfg.Env)

	logger.Info("starting backend api", "env", cfg.Env)
	logger.Debug("debug messages are enabled")

	// Init database
	dbMySQL, err := db.New(cfg.Database)
	if err != nil {
		logger.Error("mysql connect problem", "error", err)
		os.Exit(1)
	}
	defer func() {
		if err := dbMySQL.Close(); err != nil {
			logger.Error("error when closing", "error", err)
		}
	}()
	logger.Info("mysql connection done")

	// Services, Repos & API Handlers
	repos := repository.NewRepositories(dbMySQL)
	services := service.NewServices(service.Deps{
		Logger: logger,
		Config: cfg,
		Repos:  repos,
	})
	handlers := apiHttp.NewHandlers(services, logger)

	// HTTP Server
	srv := server.NewServer(cfg, handlers.Init(cfg))
	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Error("error occurred while running http server", "error", err)
		}
	}()
	logger.Info("server started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Error("failed to stop server", "error", err)
	}

	logger.Info("app stopped")
}
