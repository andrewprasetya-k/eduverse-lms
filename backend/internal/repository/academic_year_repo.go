package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type AcademicYearRepository interface {
	// Academic Year
	CreateAcademicYear(academicYear *domain.AcademicYear) error
	GetAllAcademicYears(schoolID string) ([]*domain.AcademicYear, error)
	GetAcademicYearByID(id string) (*domain.AcademicYear, error)
	UpdateAcademicYear(academicYear *domain.AcademicYear) error
	DeleteAcademicYear(id string) error
	DeactivateAllYears(schoolID string) error

	// Term
	CreateTerm(term *domain.Term) error
	GetTermByID(id string) (*domain.Term, error)
	GetTermsByYear(yearID string) ([]*domain.Term, error)
	UpdateTerm(term *domain.Term) error
	DeleteTerm(id string) error
	DeactivateAllTerms(yearID string) error
}

type academicYearRepository struct {
	db *gorm.DB
}

// constructor
func NewAcademicYearRepository(db *gorm.DB) AcademicYearRepository {
	return &academicYearRepository{db: db}
}

// --- Academic Year Implementation ---

func (r *academicYearRepository) CreateAcademicYear(academicYear *domain.AcademicYear) error {
	return r.db.Create(academicYear).Error
}

func (r *academicYearRepository) GetAllAcademicYears(schoolID string) ([]*domain.AcademicYear, error) {
	var years []*domain.AcademicYear
	err := r.db.Where("acy_sch_id = ?", schoolID).Preload("Terms").Find(&years).Error
	return years, err
}

func (r *academicYearRepository) GetAcademicYearByID(id string) (*domain.AcademicYear, error) {
	var year domain.AcademicYear
	err := r.db.Where("acy_id = ?", id).Preload("Terms").First(&year).Error
	return &year, err
}

func (r *academicYearRepository) UpdateAcademicYear(academicYear *domain.AcademicYear) error {
	return r.db.Save(academicYear).Error
}

func (r *academicYearRepository) DeleteAcademicYear(id string) error {
	return r.db.Delete(&domain.AcademicYear{}, "acy_id = ?", id).Error
}

func (r *academicYearRepository) DeactivateAllYears(schoolID string) error {
	return r.db.Model(&domain.AcademicYear{}).
		Where("acy_sch_id = ?", schoolID).
		Update("is_active", false).Error
}

// --- Term Implementation ---

func (r *academicYearRepository) CreateTerm(term *domain.Term) error {
	return r.db.Create(term).Error
}

func (r *academicYearRepository) GetTermByID(id string) (*domain.Term, error) {
	var term domain.Term
	err := r.db.Where("trm_id = ?", id).First(&term).Error
	return &term, err
}

func (r *academicYearRepository) GetTermsByYear(yearID string) ([]*domain.Term, error) {
	var terms []*domain.Term
	err := r.db.Where("trm_acy_id = ?", yearID).Find(&terms).Error
	return terms, err
}

func (r *academicYearRepository) UpdateTerm(term *domain.Term) error {
	return r.db.Save(term).Error
}

func (r *academicYearRepository) DeleteTerm(id string) error {
	return r.db.Delete(&domain.Term{}, "trm_id = ?", id).Error
}

func (r *academicYearRepository) DeactivateAllTerms(yearID string) error {
	return r.db.Model(&domain.Term{}).
		Where("trm_acy_id = ?", yearID).
		Update("is_active", false).Error
}
