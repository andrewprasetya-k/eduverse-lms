package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type SubjectService interface{
	CreateSubject(subject *domain.Subject) error
	GetAllSubjects() ([]*domain.Subject, error)
	GetSubjectByID(id string) (*domain.Subject, error)
	GetSubjectByCode(code string) (*domain.Subject, error)
	UpdateSubject(subject *domain.Subject) error
	DeleteSubject(id string) error
}

type subjectService struct {
	subjectRepo repository.SubjectRepository
	schoolRepo repository.SchoolRepository
}

// constructor
func NewSubjectService(subRepo repository.SubjectRepository, schRepo repository.SchoolRepository) SubjectService {
	return &subjectService{subjectRepo: subRepo, schoolRepo: schRepo}
}


func (s *subjectService) CreateSubject(subject *domain.Subject) error {
	return s.subjectRepo.CreateSubject(subject)
}

func (s *subjectService) GetAllSubjects() ([]*domain.Subject, error) {
	return s.subjectRepo.GetAllSubjects()
}

func (s *subjectService) GetSubjectByID(id string) (*domain.Subject, error){
	return s.subjectRepo.GetSubjectByID(id)
}

func (s *subjectService) GetSubjectByCode(code string) (*domain.Subject, error) {
	return s.subjectRepo.GetSubjectByCode(code)

}

func (s *subjectService) UpdateSubject(subject *domain.Subject) error {
	return s.subjectRepo.UpdateSubject(subject)
}

func (s *subjectService) DeleteSubject(code string) error {
	return s.subjectRepo.DeleteSubject(code)
}