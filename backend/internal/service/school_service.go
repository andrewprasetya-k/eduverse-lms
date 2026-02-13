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
	GetSchools(search string, status string, page int, limit int) ([]*domain.School, int64, error)
	GetSchoolByCode(schoolCode string) (*domain.School, error)
	GetSchoolByID(schoolID string) (*domain.School, error)
	RestoreDeletedSchool(schoolCode string) error
	UpdateSchool(school *domain.School) error
	DeleteSchool(schoolCode string) error

	//functional methods
	convertCodeToID(schoolCode string) (string, error)
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

func (s *schoolService) GetSchools(search string, status string, page int, limit int) ([]*domain.School, int64, error) {
	return s.repo.GetSchools(search, status, page, limit)
}

func (s *schoolService) GetSchoolByCode(schoolCode string) (*domain.School, error) {
	return s.repo.GetSchoolByCode(schoolCode)
}

func (s *schoolService) GetSchoolByID(schoolID string) (*domain.School, error) {
	return s.repo.GetSchoolByID(schoolID)
}

func (s *schoolService) UpdateSchool(school *domain.School) error {
    existing, err := s.repo.GetSchoolByCode(school.Code)
    
    // Kalau kodenya ketemu, cek apakah itu milik sekolah LAIN?
    if err == nil && existing != nil {
        // Jika ID yang di DB beda dengan ID yang mau kita update, berarti DUPLIKAT
        if existing.ID != school.ID {
            return fmt.Errorf("kode sekolah '%s' sudah dipakai oleh sekolah lain", school.Code)
        }
    }

    return s.repo.UpdateSchool(school)
}

func (s *schoolService) RestoreDeletedSchool(schoolCode string) error {
	schoolID, err:= s.convertCodeToID(schoolCode)
	if err != nil {
		return err
	}
	return s.repo.RestoreDeletedSchool(schoolID)
}

func (s *schoolService) DeleteSchool(schoolCode string) error {
	schoolID, err:= s.convertCodeToID(schoolCode)
	if err != nil {
		return err
	}
	return s.repo.DeleteSchool(schoolID)
}

//functional methods
func (s *schoolService) convertCodeToID(schoolCode string) (string, error) {
	school, err := s.repo.GetSchoolByCode(schoolCode)
	if err != nil {
		return "", err
	}
	return school.ID, nil
}