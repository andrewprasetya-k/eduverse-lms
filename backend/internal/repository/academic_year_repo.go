package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type AcademicYearRepository interface {
	Create(acy *domain.AcademicYear) error
	FindAll(search string, page int, limit int) ([]*domain.AcademicYear, int64, error)
	GetBySchool(schoolID string) ([]*domain.AcademicYear, error)
	GetByID(id string) (*domain.AcademicYear, error)
	Update(acy *domain.AcademicYear) error
	Delete(id string) error
	DeactivateAllExcept(schoolID string, activeID string) error
	SetActiveStatus(id string, isActive bool) error
}

type academicYearRepository struct {
	db *gorm.DB
}

func NewAcademicYearRepository(db *gorm.DB) AcademicYearRepository {
	return &academicYearRepository{db: db}
}

func (r *academicYearRepository) Create(acy *domain.AcademicYear) error {
	return r.db.Create(acy).Error
}

func (r *academicYearRepository) FindAll(search string, page int, limit int) ([]*domain.AcademicYear, int64, error) {
	var years []*domain.AcademicYear
	var total int64

	query := r.db.Model(&domain.AcademicYear{})

	if search != "" {
		searchTerm := "%" + search + "%"
		query = query.Where("acy_name ILIKE ?", searchTerm)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Order("created_at desc").Find(&years).Error
	return years, total, err
}

func (r *academicYearRepository) GetBySchool(schoolID string) ([]*domain.AcademicYear, error) {
	var years []*domain.AcademicYear
	err := r.db.Where("acy_sch_id = ?", schoolID).Order("acy_name desc").Find(&years).Error
	return years, err
}

func (r *academicYearRepository) GetByID(id string) (*domain.AcademicYear, error) {
	var acy domain.AcademicYear
	err := r.db.Where("acy_id = ?", id).First(&acy).Error
	return &acy, err
}

func (r *academicYearRepository) Update(acy *domain.AcademicYear) error {
	return r.db.Save(acy).Error
}

func (r *academicYearRepository) Delete(id string) error {
	return r.db.Delete(&domain.AcademicYear{}, "acy_id = ?", id).Error
}

// DeactivateAllExcept memastikan hanya satu tahun ajaran yang aktif per sekolah
func (r *academicYearRepository) DeactivateAllExcept(schoolID string, activeID string) error {
	return r.db.Model(&domain.AcademicYear{}).
		Where("acy_sch_id = ? AND acy_id != ?", schoolID, activeID).
		Update("is_active", false).Error
}

func (r *academicYearRepository) SetActiveStatus(id string, isActive bool) error {
	return r.db.Model(&domain.AcademicYear{}).Where("acy_id = ?", id).Update("is_active", isActive).Error
}
