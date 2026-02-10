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
	GetSchoolByCode(schoolCode string) (*domain.School, error)
	UpdateSchool(school *domain.School) error
	DeleteSchool(id string) error
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
		school.Code = s.generateRandomCode()
	} else {
		_, err := s.repo.GetSchoolByCode(school.Code)
		if err == nil {
			return fmt.Errorf("school code '%s' already exists", school.Code)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	return s.repo.CreateSchool(school)
}

func (s *schoolService) generateRandomCode() string {
    word := []rune("ABCDEFGHJKMNPQRSTUVWXYZ23456789")
    seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
    
    for range 10 { // Coba maksimal 10 kali
        code := make([]rune, 6)
        for j := range code {
            code[j] = word[seededRand.Intn(len(word))]
        }
        
        // Cek keunikan
        _, err := s.repo.GetSchoolByCode(string(code))
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return string(code)
        }
    }
    return "" // Atau handle error jika gagal dapet kode unik
}

func (s *schoolService) GetAllSchools() ([]*domain.School, error) {
	return s.repo.GetAllSchools()
}

func (s *schoolService) GetSchoolByCode(schoolCode string) (*domain.School, error) {
	return s.repo.GetSchoolByCode(schoolCode)
}

func (s *schoolService) UpdateSchool(school *domain.School) error {
	return s.repo.UpdateSchool(school)
}

func (s *schoolService) DeleteSchool(schoolCode string) error {
	return s.repo.DeleteSchool(schoolCode)
}