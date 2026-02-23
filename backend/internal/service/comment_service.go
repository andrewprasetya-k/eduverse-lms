package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type CommentService interface {
	Create(comment *domain.Comment) error
	GetBySource(sourceType string, sourceID string) ([]*domain.Comment, error)
	GetByID(id string) (*domain.Comment, error)
	Update(id string, comment *domain.Comment) error
	Delete(id string) error
	CountBySource(sourceType string, sourceID string) (int, error)
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) Create(comment *domain.Comment) error {
	return s.repo.Create(comment)
}

func (s *commentService) GetBySource(sourceType string, sourceID string) ([]*domain.Comment, error) {
	return s.repo.GetBySource(domain.SourceType(sourceType), sourceID)
}

func (s *commentService) GetByID(id string) (*domain.Comment, error) {
	return s.repo.GetByID(id)
}

func (s *commentService) Update(id string, comment *domain.Comment) error {
	comment.ID = id
	return s.repo.Update(comment)
}

func (s *commentService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *commentService) CountBySource(sourceType string, sourceID string) (int, error) {
	return s.repo.CountBySource(domain.SourceType(sourceType), sourceID)
}
