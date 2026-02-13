package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"strings"
	"time"
)

type MaterialService interface {
	Create(mat *domain.Material, mediaIDs []string) error
	FindAll(search string, classID string, page int, limit int) ([]*domain.Material, int64, error)
	GetByID(id string) (*domain.Material, error)
	Update(mat *domain.Material) error
	Delete(id string) error

	// Progress
	UpdateProgress(userID, matID string, status string) error
	GetProgress(userID, matID string) (*domain.MaterialProgress, error)
}

type materialService struct {
	repo          repository.MaterialRepository
	attService    AttachmentService
	mediaRepo     repository.MediaRepository
}

func NewMaterialService(repo repository.MaterialRepository, attService AttachmentService, mediaRepo repository.MediaRepository) MaterialService {
	return &materialService{
		repo:       repo,
		attService: attService,
		mediaRepo:  mediaRepo,
	}
}

func (s *materialService) Create(mat *domain.Material, mediaIDs []string) error {
	mat.Title = strings.TrimSpace(mat.Title)

	err := s.repo.Create(mat)
	if err != nil {
		return err
	}

	// Link attachments
	for _, mID := range mediaIDs {
		att := &domain.Attachment{
			SchoolID:   mat.SchoolID,
			SourceID:   mat.ID,
			SourceType: domain.SourceMaterial,
			MediaID:    mID,
		}
		s.attService.Link(att)
	}

	return nil
}

func (s *materialService) FindAll(search string, classID string, page int, limit int) ([]*domain.Material, int64, error) {
	return s.repo.FindAll(search, classID, page, limit)
}

func (s *materialService) GetByID(id string) (*domain.Material, error) {
	mat, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Load attachments metadata
	atts, _ := s.attService.GetBySource(string(domain.SourceMaterial), id)
	mat.Attachments = nil
	for _, a := range atts {
		mat.Attachments = append(mat.Attachments, *a)
	}

	return mat, nil
}

func (s *materialService) Update(mat *domain.Material) error {
	mat.Title = strings.TrimSpace(mat.Title)
	return s.repo.Update(mat)
}

func (s *materialService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *materialService) UpdateProgress(userID, matID string, status string) error {
	now := time.Now()
	prog := &domain.MaterialProgress{
		UserID:       userID,
		MaterialID:   matID,
		Status:       domain.StatusProgress(status),
		LastOpenedAt: &now,
	}
	return s.repo.UpsertProgress(prog)
}

func (s *materialService) GetProgress(userID, matID string) (*domain.MaterialProgress, error) {
	return s.repo.GetProgress(userID, matID)
}
