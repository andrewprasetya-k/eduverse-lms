package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"strings"
)

type AcademicYearService interface {
	Create(acy *domain.AcademicYear) error
	FindAll(search string, page int, limit int) ([]*domain.AcademicYear, int64, error)
	GetBySchool(schoolCode string) ([]*domain.AcademicYear, error)
	GetByID(id string) (*domain.AcademicYear, error)
	Update(acy *domain.AcademicYear) error
	Delete(id string) error
	Activate(id string) error
	Deactivate(id string) error
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
	acy.IsActive = false // Selalu default false saat baru dibuat
	return s.repo.Create(acy)
}

func (s *academicYearService) FindAll(search string, page int, limit int) ([]*domain.AcademicYear, int64, error) {
	return s.repo.FindAll(search, page, limit)
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
	return s.repo.Update(acy)
}

func (s *academicYearService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *academicYearService) Activate(id string) error {
	acy, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	// 1. Aktifkan tahun ajaran ini
	if err := s.repo.SetActiveStatus(id, true); err != nil {
		return err
	}

	// 2. Nonaktifkan yang lainnya di sekolah yang sama
	return s.repo.DeactivateAllExcept(acy.SchoolID, id)
}

func (s *academicYearService) Deactivate(id string) error {
	return s.repo.SetActiveStatus(id, false)
}
