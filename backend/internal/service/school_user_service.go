package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"
)

type SchoolUserService interface {
	Enroll(scu *domain.SchoolUser) error
	GetMembersBySchool(schoolID string) ([]*domain.SchoolUser, error)
	GetSchoolsByUser(userID string) ([]*domain.SchoolUser, error)
	Unenroll(id string) error
}

type schoolUserService struct {
	repo repository.SchoolUserRepository
}

func NewSchoolUserService(repo repository.SchoolUserRepository) SchoolUserService {
	return &schoolUserService{repo: repo}
}

func (s *schoolUserService) Enroll(scu *domain.SchoolUser) error {
	// 1. Validasi: Apakah sudah terdaftar di sekolah ini?
	already, err := s.repo.IsEnrolled(scu.UserID, scu.SchoolID)
	if err != nil {
		return err
	}
	if already {
		return fmt.Errorf("user sudah terdaftar sebagai anggota di sekolah ini")
	}

	return s.repo.Create(scu)
}

func (s *schoolUserService) GetMembersBySchool(schoolID string) ([]*domain.SchoolUser, error) {
	return s.repo.GetBySchool(schoolID)
}

func (s *schoolUserService) GetSchoolsByUser(userID string) ([]*domain.SchoolUser, error) {
	return s.repo.GetByUser(userID)
}

func (s *schoolUserService) Unenroll(id string) error {
	return s.repo.Delete(id)
}
