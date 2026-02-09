package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type SchoolRepository interface{
	CreateSchool(school *domain.School) error
	GetAllSchools() ([]*domain.School, error)
	GetSchoolByID(id string) (*domain.School, error)
	GetSchoolByCode(code string) (*domain.School, error)
	UpdateSchool(school *domain.School) error
	DeleteSchool(id string) error
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
	err := r.db.Find(&schools).Error
	return schools, err
}

func (r *schoolRepository) GetSchoolByID(id string) (*domain.School, error) {
	var school domain.School
	err := r.db.First(&school, "sch_id = ?", id).Error
	return &school, err
}

func (r *schoolRepository) GetSchoolByCode(code string) (*domain.School, error) {
	var school domain.School
	err := r.db.First(&school, "sch_code = ?", code).Error
	return &school, err
}

func (r *schoolRepository) UpdateSchool(school *domain.School) error {
	return r.db.Save(school).Error
}

func (r *schoolRepository) DeleteSchool(code string) error {
	return r.db.Delete(&domain.School{}, "sch_code = ?", code).Error
}