package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type EnrollmentRepository interface {
	Create(enr *domain.Enrollment) error
	GetByID(id string) (*domain.Enrollment, error)
	GetByClass(classID string) ([]*domain.Enrollment, error)
	GetByMember(schoolUserID string) ([]*domain.Enrollment, error)
	Delete(id string) error
	CheckExists(classID, schoolUserID string) (bool, error)
}

type enrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) EnrollmentRepository {
	return &enrollmentRepository{db: db}
}

func (r *enrollmentRepository) Create(enr *domain.Enrollment) error {
	return r.db.Create(enr).Error
}

func (r *enrollmentRepository) GetByID(id string) (*domain.Enrollment, error) {
	var enr domain.Enrollment
	err := r.db.Preload("SchoolUser.User").Preload("Class").
		Where("enr_id = ?", id).First(&enr).Error
	return &enr, err
}

func (r *enrollmentRepository) GetByClass(classID string) ([]*domain.Enrollment, error) {
	var results []*domain.Enrollment
	err := r.db.Preload("SchoolUser.User").
		Where("enr_cls_id = ?", classID).Find(&results).Error
	return results, err
}

func (r *enrollmentRepository) GetByMember(schoolUserID string) ([]*domain.Enrollment, error) {
	var results []*domain.Enrollment
	err := r.db.Preload("Class.School").Preload("Class.Term.AcademicYear").
		Where("enr_scu_id = ?", schoolUserID).Find(&results).Error
	return results, err
}

func (r *enrollmentRepository) Delete(id string) error {
	result := r.db.Delete(&domain.Enrollment{}, "enr_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *enrollmentRepository) CheckExists(classID, schoolUserID string) (bool, error) {
	var count int64
	err := r.db.Model(&domain.Enrollment{}).
		Where("enr_cls_id = ? AND enr_scu_id = ?", classID, schoolUserID).
		Count(&count).Error
	return count > 0, err
}
