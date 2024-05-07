package logger

import (
	"github.com/usmonzodasomon/shortener/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Logger(cfg config.Config) *slog.Logger {
	switch cfg.Env {
	case envLocal:
		opts := PrettyHandlerOptions{
			SlogOpts: &slog.HandlerOptions{
				Level: slog.LevelDebug,
			},
		}
		return slog.New(opts.NewPrettyHandler(os.Stdout))
	case envDev:
		return slog.New(slog.NewJSONHandler(os.Stdout,
			&slog.HandlerOptions{
				Level: slog.LevelDebug,
			}))
	case envProd:
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	return nil
}
