package repository

import (
	"database/sql"
	"log/slog"
)

type SegmentRepository struct {
	logger *slog.Logger
	db     *sql.DB
}

func New(logger *slog.Logger, db *sql.DB) SegmentRepository {
	return SegmentRepository{
		logger: logger,
		db:     db,
	}
}
