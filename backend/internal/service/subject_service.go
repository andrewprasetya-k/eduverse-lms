package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type SubjectService interface {
	CreateSubject(subject *domain.Subject) error
	GetAllSubjects(schoolID string) ([]*domain.Subject, error) // Tambahkan schoolID
	GetSubjectByID(id string) (*domain.Subject, error)
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
	return s.subjectRepo.CreateSubject(subject)
}

func (s *subjectService) GetAllSubjects(schoolID string) ([]*domain.Subject, error) {
	// Di sini bisa ditambahkan validasi apakah schoolID ada, jika perlu
	return s.subjectRepo.GetAllSubjects(schoolID)
}

func (s *subjectService) GetSubjectByID(id string) (*domain.Subject, error){
	return s.subjectRepo.GetSubjectByID(id)
}

func (s *subjectService) GetSubjectByCode(code string, schoolID string) (*domain.Subject, error) {
	return s.subjectRepo.GetSubjectByCode(code, schoolID)

}

func (s *subjectService) UpdateSubject(subject *domain.Subject) error {
	return s.subjectRepo.UpdateSubject(subject)
}

func (s *subjectService) DeleteSubject(code string) error {
	return s.subjectRepo.DeleteSubject(code)
}