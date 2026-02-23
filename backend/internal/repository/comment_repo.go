package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *domain.Comment) error
	GetBySource(sourceType domain.SourceType, sourceID string) ([]*domain.Comment, error)
	GetByID(id string) (*domain.Comment, error)
	Update(comment *domain.Comment) error
	Delete(id string) error
	CountBySource(sourceType domain.SourceType, sourceID string) (int, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *domain.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) GetBySource(sourceType domain.SourceType, sourceID string) ([]*domain.Comment, error) {
	var results []*domain.Comment
	err := r.db.Preload("User").
		Where("cmn_source_type = ? AND cmn_source_id = ?", sourceType, sourceID).
		Order("created_at asc").Find(&results).Error
	return results, err
}

func (r *commentRepository) GetByID(id string) (*domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Preload("User").Where("cmn_id = ?", id).First(&comment).Error
	return &comment, err
}

func (r *commentRepository) Update(comment *domain.Comment) error {
	result := r.db.Model(&domain.Comment{}).Where("cmn_id = ?", comment.ID).Updates(comment)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *commentRepository) Delete(id string) error {
	return r.db.Delete(&domain.Comment{}, "cmn_id = ?", id).Error
}

func (r *commentRepository) CountBySource(sourceType domain.SourceType, sourceID string) (int, error) {
	var count int64
	err := r.db.Model(&domain.Comment{}).
		Where("cmn_source_type = ? AND cmn_source_id = ?", sourceType, sourceID).
		Count(&count).Error
	return int(count), err
}
