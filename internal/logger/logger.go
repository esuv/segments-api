package logger

import (
	"log/slog"
	"os"
	"segments-api/internal/config"
)

func SetupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case config.LocalEnv:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case config.DevEnv:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelDebug,
			}),
		)
	case config.ProdEnv:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelWarn,
			}),
		)
	}

	return log
}
