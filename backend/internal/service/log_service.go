package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type LogService interface {
	Record(log *domain.Log) error
	GetBySchool(schoolID string, page int, limit int) ([]*domain.Log, int64, error)
	GetByUser(userID string, page int, limit int) ([]*domain.Log, int64, error)
}

type logService struct {
	repo repository.LogRepository
}

func NewLogService(repo repository.LogRepository) LogService {
	return &logService{repo: repo}
}

func (s *logService) Record(log *domain.Log) error {
	return s.repo.Create(log)
}

func (s *logService) GetBySchool(schoolID string, page int, limit int) ([]*domain.Log, int64, error) {
	return s.repo.GetBySchool(schoolID, page, limit)
}

func (s *logService) GetByUser(userID string, page int, limit int) ([]*domain.Log, int64, error) {
	return s.repo.GetByUser(userID, page, limit)
}
