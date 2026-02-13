package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"
)

type EnrollmentService interface {
	Enroll(enr *domain.Enrollment) error
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

func (s *enrollmentService) Enroll(enr *domain.Enrollment) error {
	// 1. Validasi: Apakah sudah terdaftar di kelas ini?
	already, err := s.repo.CheckExists(enr.ClassID, enr.SchoolUserID)
	if err != nil {
		return err
	}
	if already {
		return fmt.Errorf("user sudah terdaftar di kelas ini")
	}

	return s.repo.Create(enr)
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
