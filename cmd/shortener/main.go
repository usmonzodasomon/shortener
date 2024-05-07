package main

import (
	"github.com/usmonzodasomon/shortener/internal/config"
	"github.com/usmonzodasomon/shortener/pkg/logger"
	"github.com/usmonzodasomon/shortener/pkg/postgres"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.Logger(*cfg)

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

}
