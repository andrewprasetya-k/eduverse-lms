package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *domain.Comment) error
	GetBySourceInSchool(sourceType domain.SourceType, sourceID string, schoolID string) ([]*domain.Comment, error)
	GetByIDInSchool(id string, schoolID string) (*domain.Comment, error)
	UpdateInSchool(comment *domain.Comment, schoolID string) error
	DeleteInSchool(id string, schoolID string) error
	CountBySourceInSchool(sourceType domain.SourceType, sourceID string, schoolID string) (int, error)
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

func (r *commentRepository) GetBySourceInSchool(sourceType domain.SourceType, sourceID string, schoolID string) ([]*domain.Comment, error) {
	var results []*domain.Comment
	err := r.db.Preload("User").
		Where("cmn_source_type = ? AND cmn_source_id = ? AND cmn_sch_id = ?", sourceType, sourceID, schoolID).
		Order("created_at asc").Find(&results).Error
	return results, err
}

func (r *commentRepository) GetByIDInSchool(id string, schoolID string) (*domain.Comment, error) {
	var comment domain.Comment
	err := r.db.Preload("User").Where("cmn_id = ? AND cmn_sch_id = ?", id, schoolID).First(&comment).Error
	return &comment, err
}

func (r *commentRepository) UpdateInSchool(comment *domain.Comment, schoolID string) error {
	result := r.db.Model(&domain.Comment{}).Where("cmn_id = ? AND cmn_sch_id = ?", comment.ID, schoolID).Updates(comment)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *commentRepository) DeleteInSchool(id string, schoolID string) error {
	result := r.db.Where("cmn_id = ? AND cmn_sch_id = ?", id, schoolID).Delete(&domain.Comment{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *commentRepository) CountBySourceInSchool(sourceType domain.SourceType, sourceID string, schoolID string) (int, error) {
	var count int64
	err := r.db.Model(&domain.Comment{}).
		Where("cmn_source_type = ? AND cmn_source_id = ? AND cmn_sch_id = ?", sourceType, sourceID, schoolID).
		Count(&count).Error
	return int(count), err
}
