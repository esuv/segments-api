package service

import (
	"log/slog"
	"segments-api/internal/model/segment"
)

type SegmentRepository interface {
	Create(name string) (segment.Segment, error)
	Delete(name string) error
	AddUser(add []string, remove []string, userId int) error
	GetAllByUser(userID int) ([]segment.Segment, error)
}

type SegmentServiceImpl struct {
	repo SegmentRepository
	log  *slog.Logger
}

func New(repo SegmentRepository, log *slog.Logger) *SegmentServiceImpl {
	return &SegmentServiceImpl{repo: repo, log: log}
}

func (s SegmentServiceImpl) Create(slug string) (segment.Segment, error) {
	sgm, err := s.repo.Create(slug)
	if err != nil {
		return segment.Segment{}, err
	}

	return sgm, nil
}

func (s SegmentServiceImpl) Delete(slug string) error {
	return s.repo.Delete(slug)
}

func (s SegmentServiceImpl) AddUser(add []string, remove []string, userId int) error {
	return s.repo.AddUser(add, remove, userId)
}

func (s SegmentServiceImpl) GetAllByUser(userID int) ([]segment.Segment, error) {
	return s.repo.GetAllByUser(userID)
}
