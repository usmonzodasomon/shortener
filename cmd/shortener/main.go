package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/usmonzodasomon/shortener/internal/config"
	"github.com/usmonzodasomon/shortener/internal/routes"
	"github.com/usmonzodasomon/shortener/pkg/logger"
	"github.com/usmonzodasomon/shortener/pkg/postgres"
	"github.com/usmonzodasomon/shortener/pkg/server"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.Logger(cfg.Env)

	logger.Info("starting url-shortener",
		slog.String("env", cfg.Env))

	logger.Debug("debug messages are enabled")

	connection, err := postgres.GetConnection(*cfg)
	if err != nil {
		logger.Error("Failed to connect to database: ", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer postgres.CloseConnection(connection)

	logger.Info("connected to database")

	logger.Info("starting server", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	r := chi.NewRouter()
	routes.SetUpRoutes(r, connection, logger)

	srv := server.Server{}
	go func() {
		if err := srv.Run(*cfg, r); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("failed to start server")
		}
	}()

	logger.Info(fmt.Sprintf("server started on %s", cfg.Address))

	<-done
	logger.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("failed to stop server", slog.String("error", err.Error()))
		return
	}

	logger.Info("server stopped")
}
