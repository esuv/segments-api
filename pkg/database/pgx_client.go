package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log/slog"
	"segments-api/internal/logger/sl"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
}

func New(ctx context.Context, cfg PostgresConfig, logger *slog.Logger) *pgxpool.Pool {
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.Password,
		cfg.DatabaseName,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		cfg.SSLMode,
	)

	client, err := pgxpool.Connect(ctx, databaseURL)
	if err != nil {
		logger.Error("database connection failed", sl.Err(err))
	}

	return client
}
