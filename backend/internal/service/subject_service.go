package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"
	"strings"
)

type SubjectService interface {
	Create(subject *domain.Subject) error
	FindAll(search string, page int, limit int) ([]*domain.Subject, int64, error)
	GetBySchool(schoolCode string) ([]*domain.Subject, error)
	GetByID(id string) (*domain.Subject, error)
	GetByCode(schoolCode string, subjectCode string) (*domain.Subject, error)
	Update(subject *domain.Subject) error
	Delete(id string) error
}

type subjectService struct {
	repo          repository.SubjectRepository
	schoolService SchoolService
}

func NewSubjectService(repo repository.SubjectRepository, schoolService SchoolService) SubjectService {
	return &subjectService{
		repo:          repo,
		schoolService: schoolService,
	}
}

func (s *subjectService) Create(subject *domain.Subject) error {
	subject.Name = strings.TrimSpace(subject.Name)
	subject.Code = strings.ToUpper(strings.TrimSpace(subject.Code))

	// 1. Validasi Duplikasi Kode di Sekolah yang sama
	exists, err := s.repo.CheckDuplicateCode(subject.SchoolID, subject.Code, "")
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("kode mata pelajaran '%s' sudah terdaftar di sekolah ini", subject.Code)
	}

	return s.repo.Create(subject)
}

func (s *subjectService) FindAll(search string, page int, limit int) ([]*domain.Subject, int64, error) {
	return s.repo.FindAll(search, page, limit)
}

func (s *subjectService) GetBySchool(schoolCode string) ([]*domain.Subject, error) {
	schoolID, err := s.schoolService.ConvertCodeToID(schoolCode)
	if err != nil {
		return nil, err
	}
	return s.repo.GetBySchool(schoolID)
}

func (s *subjectService) GetByID(id string) (*domain.Subject, error) {
	return s.repo.GetByID(id)
}

func (s *subjectService) GetByCode(schoolCode string, subjectCode string) (*domain.Subject, error) {
	schoolID, err := s.schoolService.ConvertCodeToID(schoolCode)
	if err != nil {
		return nil, err
	}
	return s.repo.GetByCode(schoolID, strings.ToUpper(subjectCode))
}

func (s *subjectService) Update(subject *domain.Subject) error {
	subject.Name = strings.TrimSpace(subject.Name)
	subject.Code = strings.ToUpper(strings.TrimSpace(subject.Code))

	// 1. Validasi Duplikasi Kode
	exists, err := s.repo.CheckDuplicateCode(subject.SchoolID, subject.Code, subject.ID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("kode mata pelajaran '%s' sudah terdaftar di sekolah ini", subject.Code)
	}

	return s.repo.Update(subject)
}

func (s *subjectService) Delete(id string) error {
	// TODO: Cek apakah mata pelajaran sedang digunakan di tabel subject_classes
	return s.repo.Delete(id)
}
