package repository

import (
	"backend/internal/domain"
	"fmt"

	"gorm.io/gorm"
)

type SchoolRepository interface{
	CreateSchool(school *domain.School) error
	GetSchools(search string, status string, page int, limit int, sortBy string, order string) ([]*domain.School, int64, error)
	GetSchoolByCode(schoolCode string) (*domain.School, error)
	GetSchoolByID(schoolID string) (*domain.School, error)
	RestoreDeletedSchool(schoolID string) error
	UpdateSchool(school *domain.School) error
	DeleteSchool(schoolID string) error
	HardDeleteSchool(schoolID string) error
	CheckEmailExists(email string, excludeID string) (bool, error)
	CheckPhoneExists(phone string, excludeID string) (bool, error)
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

func (r *schoolRepository) GetSchools(search string, status string, page int, limit int, sortBy string, order string) ([]*domain.School, int64, error) {
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
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Sorting
	if sortBy == "" {
		sortBy = "created_at"
	}
	if order == "" {
		order = "desc"
	}
	query = query.Order(fmt.Sprintf("%s %s", sortBy, order))

	//pagiatanion
	offset := (page - 1)*limit
	err := query.Limit(limit).Offset(offset).Find(&schools).Error
	return schools, total, err
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

func (r *schoolRepository) HardDeleteSchool(schoolID string) error {
	return r.db.Unscoped().Delete(&domain.School{}, "sch_id = ?", schoolID).Error
}

func (r *schoolRepository) CheckEmailExists(email string, excludeID string) (bool, error) {
	var count int64
	query := r.db.Model(&domain.School{}).Where("sch_email = ?", email)
	if excludeID != "" {
		query = query.Where("sch_id != ?", excludeID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}

func (r *schoolRepository) CheckPhoneExists(phone string, excludeID string) (bool, error) {
	var count int64
	query := r.db.Model(&domain.School{}).Where("sch_phone = ?", phone)
	if excludeID != "" {
		query = query.Where("sch_id != ?", excludeID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}