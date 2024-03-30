package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log/slog"
	"segments-api/internal/logger/sl"
)

func New(ctx context.Context, cfg PostgresConfig, logger *slog.Logger) *pgxpool.Pool {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Password,
		cfg.DatabaseName,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		cfg.SSLMode,
	)

	client, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		logger.Error("database connection failed", sl.Err(err))
	}

	return client
}
