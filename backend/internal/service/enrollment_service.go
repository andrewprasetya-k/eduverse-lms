package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type EnrollmentService interface {
	Enroll(schoolID string, classID string, schoolUserIDs []string, role string) error
	GetByClass(classID string) ([]*domain.Enrollment, error)
	GetByMember(schoolUserID string) ([]*domain.Enrollment, error)
	Unenroll(id string) error
}

type enrollmentService struct {
	repo repository.EnrollmentRepository
}

func NewEnrollmentService(repo repository.EnrollmentRepository) EnrollmentService {
	return &enrollmentService{repo: repo}
}

func (s *enrollmentService) Enroll(schoolID string, classID string, schoolUserIDs []string, role string) error {
	for _, scuID := range schoolUserIDs {
		// 1. Validasi: Apakah sudah terdaftar di kelas ini?
		already, err := s.repo.CheckExists(classID, scuID)
		if err != nil {
			return err
		}
		if already {
			continue // Jika sudah ada, lewati user ini
		}

		enr := domain.Enrollment{
			SchoolID:     schoolID,
			ClassID:      classID,
			SchoolUserID: scuID,
			Role:         role,
		}

		if err := s.repo.Create(&enr); err != nil {
			return err
		}
	}
	return nil
}

func (s *enrollmentService) GetByClass(classID string) ([]*domain.Enrollment, error) {
	return s.repo.GetByClass(classID)
}

func (s *enrollmentService) GetByMember(schoolUserID string) ([]*domain.Enrollment, error) {
	return s.repo.GetByMember(schoolUserID)
}

func (s *enrollmentService) Unenroll(id string) error {
	return s.repo.Delete(id)
}
