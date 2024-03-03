package service

import (
	"log/slog"
	"segments-api/internal/repository"
)

type SegmentServiceImpl struct {
	log  *slog.Logger
	repo repository.SegmentRepository
}

func New(logger *slog.Logger, repo repository.SegmentRepository) *SegmentServiceImpl {
	return &SegmentServiceImpl{
		log:  logger,
		repo: repo,
	}
}

func (s *SegmentServiceImpl) Create(slug string) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (s *SegmentServiceImpl) Delete(slug string) error {
	//TODO implement me
	panic("implement me")
}

func (s *SegmentServiceImpl) AddUser(add []string, remove []string, userId int) error {
	//TODO implement me
	panic("implement me")
}

func (s *SegmentServiceImpl) GetAllByUser(userID int) {
	//TODO implement me
	panic("implement me")
}
