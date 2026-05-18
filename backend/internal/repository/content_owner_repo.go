package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type ContentOwnerRepository interface {
	GetOwnerUserID(sourceType domain.SourceType, sourceID string) (string, error)
}

type contentOwnerRepository struct {
	db *gorm.DB
}

func NewContentOwnerRepository(db *gorm.DB) ContentOwnerRepository {
	return &contentOwnerRepository{db: db}
}

func (r *contentOwnerRepository) GetOwnerUserID(sourceType domain.SourceType, sourceID string) (string, error) {
	var ownerID string
	var err error

	switch sourceType {
	case domain.SourceFeed:
		err = r.db.Model(&domain.Feed{}).Where("fds_id = ?", sourceID).Pluck("created_by", &ownerID).Error
	case domain.SourceMaterial:
		err = r.db.Model(&domain.Material{}).Where("mat_id = ?", sourceID).Pluck("created_by", &ownerID).Error
	case domain.SourceAssignment:
		err = r.db.Model(&domain.Assignment{}).Where("asg_id = ?", sourceID).Pluck("created_by", &ownerID).Error
	case domain.SourceSubmission:
		err = r.db.Model(&domain.Submission{}).Where("sbm_id = ?", sourceID).Pluck("sbm_usr_id", &ownerID).Error
	default:
		return "", nil
	}

	return ownerID, err
}
