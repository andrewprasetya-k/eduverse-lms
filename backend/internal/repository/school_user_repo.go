package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type SchoolUserRepository interface {
	Create(scu *domain.SchoolUser) error
	GetBySchool(schoolID string) ([]*domain.SchoolUser, error)
	GetByUser(userID string) ([]*domain.SchoolUser, error)
	Delete(id string) error
	IsEnrolled(userID string, schoolID string) (bool, error)
}

type schoolUserRepository struct {
	db *gorm.DB
}

func NewSchoolUserRepository(db *gorm.DB) SchoolUserRepository {
	return &schoolUserRepository{db: db}
}

func (r *schoolUserRepository) Create(scu *domain.SchoolUser) error {
	return r.db.Create(scu).Error
}

func (r *schoolUserRepository) GetBySchool(schoolID string) ([]*domain.SchoolUser, error) {
	var members []*domain.SchoolUser
	err := r.db.Preload("User").Where("scu_sch_id = ?", schoolID).Find(&members).Error
	return members, err
}

func (r *schoolUserRepository) GetByUser(userID string) ([]*domain.SchoolUser, error) {
	var schools []*domain.SchoolUser
	err := r.db.Preload("School").Where("scu_usr_id = ?", userID).Find(&schools).Error
	return schools, err
}

func (r *schoolUserRepository) Delete(id string) error {
	result := r.db.Delete(&domain.SchoolUser{}, "scu_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *schoolUserRepository) IsEnrolled(userID string, schoolID string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.SchoolUser{}).
		Where("scu_usr_id = ? AND scu_sch_id = ?", userID, schoolID).
		Count(&count).Error
	return count > 0, err
}
