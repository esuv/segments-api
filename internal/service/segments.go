package service

import (
	"log/slog"
	"segments-api/internal/model/segment"
)

type SegmentRepository interface {
	GetByName(name string) segment.Segment
	Create(name string) (int, error)
}

type SegmentServiceImpl struct {
	repo SegmentRepository
	log  *slog.Logger
}

func New(repo SegmentRepository, logger *slog.Logger) *SegmentServiceImpl {
	return &SegmentServiceImpl{
		repo: repo,
		log:  logger,
	}
}

func (s SegmentServiceImpl) Create(name string) (int, error) {
	segmentID, err := s.repo.Create(name)
	if err != nil {
		return 0, err
	}

	return segmentID, nil
}

func (s SegmentServiceImpl) Delete(slug string) error {
	//TODO implement me
	panic("implement me")
}

func (s SegmentServiceImpl) AddUser(add []string, remove []string, userId int) error {
	//TODO implement me
	panic("implement me")
}

func (s SegmentServiceImpl) GetAllByUser(userID int) ([]segment.Segment, error) {
	//TODO implement me
	panic("implement me")
}
