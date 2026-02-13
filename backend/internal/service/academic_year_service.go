package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"strings"
)

type AcademicYearService interface {
	Create(acy *domain.AcademicYear) error
	GetBySchool(schoolCode string) ([]*domain.AcademicYear, error)
	GetByID(id string) (*domain.AcademicYear, error)
	Update(acy *domain.AcademicYear) error
	Delete(id string) error
}

type academicYearService struct {
	repo          repository.AcademicYearRepository
	schoolService SchoolService
}

func NewAcademicYearService(repo repository.AcademicYearRepository, schoolService SchoolService) AcademicYearService {
	return &academicYearService{
		repo:          repo,
		schoolService: schoolService,
	}
}

func (s *academicYearService) Create(acy *domain.AcademicYear) error {
	acy.Name = strings.TrimSpace(acy.Name)

	// Paksa aktif untuk tahun ajaran baru
	acy.IsActive = true

	err := s.repo.Create(acy)
	if err != nil {
		return err
	}

	// Otomatis nonaktifkan yang lain
	return s.repo.DeactivateAllExcept(acy.SchoolID, acy.ID)
}

func (s *academicYearService) GetBySchool(schoolCode string) ([]*domain.AcademicYear, error) {
	schoolID, err := s.schoolService.ConvertCodeToID(schoolCode)
	if err != nil {
		return nil, err
	}
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
