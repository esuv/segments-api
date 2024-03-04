package repository

import (
	"log/slog"
	"segments-api/internal/model/segment"
	"segments-api/pkg/database"
)

type SegmentRepositoryImpl struct {
	client database.Client
	logger *slog.Logger
}

func New(client database.Client, logger *slog.Logger) *SegmentRepositoryImpl {
	return &SegmentRepositoryImpl{client: client, logger: logger}
}

func (r SegmentRepositoryImpl) GetByName(name string) segment.Segment {
	return segment.Segment{}
}

func (r SegmentRepositoryImpl) Create(name string) (int, error) {
	//TODO implement me
	panic("implement me")
}
