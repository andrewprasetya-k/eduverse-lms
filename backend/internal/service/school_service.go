package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

type SchoolService interface {
	CreateSchool(school *domain.School) error
	GetAllSchools() ([]*domain.School, error)
	GetSchoolByID(id string) (*domain.School, error)
	GetSchoolByCode(code string) (*domain.School, error)
	UpdateSchool(school *domain.School) error
	DeleteSchool(code string) error
}

type schoolService struct {
	repo repository.SchoolRepository
}

// constructor
func NewSchoolService(repo repository.SchoolRepository) SchoolService {
	return &schoolService{repo: repo}
}

func (s *schoolService) CreateSchool(school *domain.School) error {
	// 1. Jika code kosong, generate otomatis dengan pengecekan keunikan
	if school.Code == "" {
		for {
			newCode := s.generateRandomCode()
			codeExist, err := s.repo.GetSchoolByCode(newCode)
			if codeExist != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					school.Code = newCode
					break
				}
				return err
			}
		}
	} else {
		codeExist, err := s.repo.GetSchoolByCode(school.Code)
		if codeExist != nil {
			return fmt.Errorf("school code '%s' already exists", school.Code)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return s.repo.CreateSchool(school)
}

func (s *schoolService) generateRandomCode() string {
	word := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	code := make([]rune, 6)
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range code {
		code[i] = word[seededRand.Intn(len(word))]
	}
	return string(code)
}

func (s *schoolService) GetAllSchools() ([]*domain.School, error) {
	return s.repo.GetAllSchools()
}


func (s *schoolService) GetSchoolByID(id string) (*domain.School, error) {
	return s.repo.GetSchoolByID(id)
}

func (s *schoolService) GetSchoolByCode(code string) (*domain.School, error) {
	return s.repo.GetSchoolByCode(code)
}

func (s *schoolService) UpdateSchool(school *domain.School) error {
	return s.repo.UpdateSchool(school)
}

func (s *schoolService) DeleteSchool(code string) error {
	return s.repo.DeleteSchool(code)
}