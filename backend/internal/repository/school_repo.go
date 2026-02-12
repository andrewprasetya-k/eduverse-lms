package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type SchoolRepository interface{
	CreateSchool(school *domain.School) error
	GetAllSchools() ([]*domain.School, error)
	GetActiveSchools() ([]*domain.School, error)
	GetDeletedSchools() ([]*domain.School, error)
	GetSchoolByCode(schoolCode string) (*domain.School, error)
	GetSchoolByID(schoolID string) (*domain.School, error)
	RestoreDeletedSchool(schoolID string) error
	UpdateSchool(school *domain.School) error
	DeleteSchool(schoolID string) error
}

type schoolRepository struct {
	db *gorm.DB
}

//constructor
func NewSchoolRepository(db *gorm.DB) SchoolRepository {
	return &schoolRepository{db: db}
}

func (r *schoolRepository) CreateSchool(school *domain.School) error {
	return r.db.Create(school).Error
}

func (r *schoolRepository) GetAllSchools() ([]*domain.School, error) {
	var schools []*domain.School
	err := r.db.Unscoped().Find(&schools).Error
	return schools, err
}

func (r *schoolRepository) GetActiveSchools() ([]*domain.School, error) {
	var schools []*domain.School
	err := r.db.Find(&schools).Error
	return schools, err
}

func (r *schoolRepository) GetDeletedSchools() ([]*domain.School, error) {
	var schools []*domain.School
	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").Find(&schools).Error
	return schools, err
}

func (r *schoolRepository) GetSchoolByCode(schoolCode string) (*domain.School, error) {
	var school domain.School
	err := r.db.Where("sch_code = ?", schoolCode).First(&school).Error
	return &school, err
}

func (r *schoolRepository) GetSchoolByID(schoolID string) (*domain.School, error) {
	var school domain.School
	err := r.db.Where("sch_id = ?", schoolID).First(&school).Error
	return &school, err
}

func (r *schoolRepository) UpdateSchool(school *domain.School) error {
	return r.db.Updates(school).Error
}

func (r *schoolRepository) RestoreDeletedSchool(schoolID string) error {
	return r.db.Unscoped().Model(&domain.School{}).Where("sch_id = ?", schoolID).Update("deleted_at", nil).Error
}

func (r *schoolRepository) DeleteSchool(schoolID string) error {
	return r.db.Delete(&domain.School{}, "sch_id = ?", schoolID).Error
}