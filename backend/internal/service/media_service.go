package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"backend/internal/storage"
	"context"
	"io"
)

type MediaService interface {
	RecordMetadata(media *domain.Media) error
	UploadAndRecord(ctx context.Context, media *domain.Media, content io.Reader) error
	GetByID(id string) (*domain.Media, error)
	GetByOwner(ownerType string, ownerID string) ([]*domain.Media, error)
	Delete(id string) error
}

type mediaService struct {
	repo    repository.MediaRepository
	storage storage.Provider
}

func NewMediaService(repo repository.MediaRepository, storageProvider storage.Provider) MediaService {
	if storageProvider == nil {
		storageProvider = storage.NewDisabledStorage()
	}
	return &mediaService{repo: repo, storage: storageProvider}
}

func (s *mediaService) RecordMetadata(media *domain.Media) error {
	return s.repo.Create(media)
}

func (s *mediaService) UploadAndRecord(ctx context.Context, media *domain.Media, content io.Reader) error {
	publicURL, err := s.storage.Upload(ctx, media.StoragePath, content, media.MimeType)
	if err != nil {
		return err
	}

	media.FileURL = publicURL
	if err := s.repo.Create(media); err != nil {
		_ = s.storage.Delete(ctx, media.StoragePath)
		return err
	}
	return nil
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
