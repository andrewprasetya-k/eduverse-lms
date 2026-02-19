package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type AttachmentService interface {
	Link(att *domain.Attachment) error
	GetBySource(sourceType string, sourceID string) ([]*domain.Attachment, error)
	Unlink(id string) error
	UnlinkBySource(sourceType string, sourceID string) error
}

type attachmentService struct {
	repo repository.AttachmentRepository
}

func NewAttachmentService(repo repository.AttachmentRepository) AttachmentService {
	return &attachmentService{repo: repo}
}

func (s *attachmentService) Link(att *domain.Attachment) error {
	return s.repo.Create(att)
}

func (s *attachmentService) GetBySource(sourceType string, sourceID string) ([]*domain.Attachment, error) {
	return s.repo.GetBySource(domain.SourceType(sourceType), sourceID)
}

func (s *attachmentService) Unlink(id string) error {
	return s.repo.Delete(id)
}

func (s *attachmentService) UnlinkBySource(sourceType string, sourceID string) error {
	return s.repo.DeleteBySource(domain.SourceType(sourceType), sourceID)
}
