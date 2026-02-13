package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"strings"
)

type AcademicYearService interface {
	Create(acy *domain.AcademicYear) error
	GetBySchool(schoolID string) ([]*domain.AcademicYear, error)
	GetByID(id string) (*domain.AcademicYear, error)
	Update(acy *domain.AcademicYear) error
	Delete(id string) error
}

type academicYearService struct {
	repo repository.AcademicYearRepository
}

func NewAcademicYearService(repo repository.AcademicYearRepository) AcademicYearService {
	return &academicYearService{repo: repo}
}

func (s *academicYearService) Create(acy *domain.AcademicYear) error {
	acy.Name = strings.TrimSpace(acy.Name)
	
	err := s.repo.Create(acy)
	if err != nil {
		return err
	}

	// Jika di-set aktif, nonaktifkan yang lain
	if acy.IsActive {
		return s.repo.DeactivateAllExcept(acy.SchoolID, acy.ID)
	}
	return nil
}

func (s *academicYearService) GetBySchool(schoolID string) ([]*domain.AcademicYear, error) {
	return s.repo.GetBySchool(schoolID)
}

func (s *academicYearService) GetByID(id string) (*domain.AcademicYear, error) {
	return s.repo.GetByID(id)
}

func (s *academicYearService) Update(acy *domain.AcademicYear) error {
	acy.Name = strings.TrimSpace(acy.Name)
	
	err := s.repo.Update(acy)
	if err != nil {
		return err
	}

	if acy.IsActive {
		return s.repo.DeactivateAllExcept(acy.SchoolID, acy.ID)
	}
	return nil
}

func (s *academicYearService) Delete(id string) error {
	return s.repo.Delete(id)
}
