package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type AttachmentRepository interface {
	Create(att *domain.Attachment) error
	GetBySource(sourceType domain.SourceType, sourceID string) ([]*domain.Attachment, error)
	Delete(id string) error
	DeleteBySource(sourceType domain.SourceType, sourceID string) error
}

type attachmentRepository struct {
	db *gorm.DB
}

func NewAttachmentRepository(db *gorm.DB) AttachmentRepository {
	return &attachmentRepository{db: db}
}

func (r *attachmentRepository) Create(att *domain.Attachment) error {
	return r.db.Create(att).Error
}

func (r *attachmentRepository) GetBySource(sourceType domain.SourceType, sourceID string) ([]*domain.Attachment, error) {
	var results []*domain.Attachment
	err := r.db.Preload("Media").
		Where("att_source_type = ? AND att_source_id = ?", sourceType, sourceID).
		Find(&results).Error
	return results, err
}

func (r *attachmentRepository) Delete(id string) error {
	result := r.db.Delete(&domain.Attachment{}, "att_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *attachmentRepository) DeleteBySource(sourceType domain.SourceType, sourceID string) error {
	return r.db.Where("att_source_type = ? AND att_source_id = ?", sourceType, sourceID).
		Delete(&domain.Attachment{}).Error
}
