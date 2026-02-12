package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type SchoolRepository interface{
	CreateSchool(school *domain.School) error
	GetSchools(search string, status string, page int, limit int) ([]*domain.School, error)
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

func (r *schoolRepository) GetSchools(search string, status string, page int, limit int) ([]*domain.School, error) {
	var schools []*domain.School
	var total int64

	query := r.db.Model(&domain.School{})

	// Filter by status
	switch status {
	case "active":
		query = query.Where("deleted_at IS NULL")
	case "deleted":
		query = query.Unscoped().Where("deleted_at IS NOT NULL")
	default:
		// "all" or empty -> Include soft-deleted records
		query = query.Unscoped()
	}

	// Filter by search term (Name or Code)
	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("sch_name ILIKE ? OR sch_code ILIKE ?", searchTerm, searchTerm)
	}
	//hitung total data
	query.Count(&total)

	//pagiatanion
	offset := (page - 1)*limit
	err := query.Limit(limit).Offset(offset).Find(&schools).Error
	return schools, err
}

func (r *schoolRepository) GetSchoolByCode(schoolCode string) (*domain.School, error) {
	var school domain.School
	err := r.db.Unscoped().Where("sch_code = ?", schoolCode).First(&school).Error
	return &school, err
}

func (r *schoolRepository) GetSchoolByID(schoolID string) (*domain.School, error) {
	var school domain.School
	err := r.db.Unscoped().Where("sch_id = ?", schoolID).First(&school).Error
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