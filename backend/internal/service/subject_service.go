package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"errors"
)

type SubjectService interface {
	CreateSubject(subject *domain.Subject) error
	GetAllSubjects(schoolID string) ([]*domain.Subject, error) // Tambahkan schoolID
	GetSubjectByCode(code string, schoolID string) (*domain.Subject, error)
	UpdateSubject(subject *domain.Subject) error
	DeleteSubject(id string) error
}

type subjectService struct {
	subjectRepo repository.SubjectRepository
	schoolRepo  repository.SchoolRepository
}

// constructor
func NewSubjectService(subRepo repository.SubjectRepository, schRepo repository.SchoolRepository) SubjectService {
	return &subjectService{subjectRepo: subRepo, schoolRepo: schRepo}
}

func (s *subjectService) CreateSubject(subject *domain.Subject) error {
	// Cek apakah kode sudah ada di sekolah yang sama
	existing, _ := s.subjectRepo.GetSubjectByCode(subject.Code, subject.SchoolID)
	if existing != nil && existing.ID != "" {
		return errors.New("subject code already exists in this school")
	}
	return s.subjectRepo.CreateSubject(subject)
}

func (s *subjectService) GetAllSubjects(schoolID string) ([]*domain.Subject, error) {
	// Di sini bisa ditambahkan validasi apakah schoolID ada, jika perlu
	return s.subjectRepo.GetAllSubjects(schoolID)
}

func (s *subjectService) GetSubjectByCode(code string, schoolID string) (*domain.Subject, error) {
	return s.subjectRepo.GetSubjectByCode(code, schoolID)

}

func (s *subjectService) UpdateSubject(subject *domain.Subject) error {
	return s.subjectRepo.UpdateSubject(subject)
}

func (s *subjectService) DeleteSubject(id string) error {
	return s.subjectRepo.DeleteSubject(id)
}