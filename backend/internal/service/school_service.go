package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type SchoolService interface{
	CreateSchool(school *domain.School) error
	GetAllSchools() ([]*domain.School, error)
	GetSchoolByID(id string) (*domain.School, error)
	UpdateSchool(school *domain.School) error
	DeleteSchool(id string) error
}

type schoolService struct {
	repo repository.SchoolRepository
}

//constructor
func NewSchoolService(repo repository.SchoolRepository) SchoolService {
	return &schoolService{repo: repo}
}

func (s *schoolService) CreateSchool(school *domain.School) error {
	return s.repo.CreateSchool(school)
}

func (s *schoolService) GetAllSchools() ([]*domain.School, error) {
	return s.repo.GetAllSchools()
}

func (s *schoolService) GetSchoolByID(id string) (*domain.School, error) {
	return s.repo.GetSchoolByID(id)
}

func (s *schoolService) UpdateSchool(school *domain.School) error {
	return s.repo.UpdateSchool(school)
}

func (s *schoolService) DeleteSchool(id string) error {
	return s.repo.DeleteSchool(id)
}