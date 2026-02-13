package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type MediaService interface {
	RecordMetadata(media *domain.Media) error
	GetByID(id string) (*domain.Media, error)
	GetByOwner(ownerType string, ownerID string) ([]*domain.Media, error)
	Delete(id string) error
}

type mediaService struct {
	repo repository.MediaRepository
}

func NewMediaService(repo repository.MediaRepository) MediaService {
	return &mediaService{repo: repo}
}

func (s *mediaService) RecordMetadata(media *domain.Media) error {
	return s.repo.Create(media)
}

func (s *mediaService) GetByID(id string) (*domain.Media, error) {
	return s.repo.GetByID(id)
}

func (s *mediaService) GetByOwner(ownerType string, ownerID string) ([]*domain.Media, error) {
	return s.repo.GetByOwner(domain.OwnerType(ownerType), ownerID)
}

func (s *mediaService) Delete(id string) error {
	// TODO: Integrate with actual file deletion (S3/Local)
	return s.repo.Delete(id)
}
