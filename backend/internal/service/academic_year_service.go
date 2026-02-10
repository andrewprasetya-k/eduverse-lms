package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type AcademicYearService interface{
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

type academicYearService struct {
	academicYear repository.AcademicYearRepository
}

//constructor
func NewAcademicYearService(repo repository.AcademicYearRepository) AcademicYearService {
	return &academicYearService{academicYear: repo}
}

func (s *academicYearService) CreateAcademicYear(academicYear *domain.AcademicYear) error {
	return s.academicYear.CreateAcademicYear(academicYear)
}

func (s *academicYearService) GetAllAcademicYears(schoolID string) ([]*domain.AcademicYear, error)	 {		
	return s.academicYear.GetAllAcademicYears(schoolID)
}

func (s *academicYearService) GetAcademicYearByID(id string) (*domain.AcademicYear, error) {
	return s.academicYear.GetAcademicYearByID(id)
}

func (s *academicYearService) UpdateAcademicYear(academicYear *domain.AcademicYear) error {
	return s.academicYear.UpdateAcademicYear(academicYear)
}

func (s *academicYearService) DeleteAcademicYear(id string) error {
	return s.academicYear.DeleteAcademicYear(id)
}

func (s *academicYearService) DeactivateAllYears(schoolID string) error {
	return s.academicYear.DeactivateAllYears(schoolID)
}

func (s *academicYearService) CreateTerm(term *domain.Term) error {
	return s.academicYear.CreateTerm(term)
}

func (s *academicYearService) GetTermByID(id string) (*domain.Term, error) {
	return s.academicYear.GetTermByID(id)
}

func (s *academicYearService) GetTermsByYear(yearID string) ([]*domain.Term, error) {
	return s.academicYear.GetTermsByYear(yearID)
}


func (s *academicYearService) UpdateTerm(term *domain.Term) error {
	return s.academicYear.UpdateTerm(term)
}

func (s *academicYearService) DeleteTerm(id string) error {
	return s.academicYear.DeleteTerm(id)
}

func (s *academicYearService) DeactivateAllTerms(yearID string) error {
	return s.academicYear.DeactivateAllTerms(yearID)
}