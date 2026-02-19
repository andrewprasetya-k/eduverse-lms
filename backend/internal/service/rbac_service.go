package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"
	"strings"
)

type RBACService interface {
	// Role management
	CreateRole(role *domain.Role) error
	GetAllRoles() ([]*domain.Role, error)
	GetRoleByID(id string) (*domain.Role, error)
	UpdateRole(role *domain.Role) error
	DeleteRole(id string) error

	// User-Role management
	AssignRoleToUser(schoolUserID string, roleID string) error
	RemoveRoleFromUser(schoolUserID string, roleID string) error
	GetUserRoles(schoolUserID string) ([]*domain.UserRole, error)
	SyncUserRoles(schoolUserID string, roleIDs []string) error
}

type rbacService struct {
	repo repository.RBACRepository
}

func NewRBACService(repo repository.RBACRepository) RBACService {
	return &rbacService{
		repo: repo,
	}
}

func (s *rbacService) CreateRole(role *domain.Role) error {
	role.Name = strings.TrimSpace(role.Name)

	// 1. Validasi Duplikasi Nama Role Global
	exists, err := s.repo.CheckDuplicateRoleName(role.Name, "")
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("role '%s' sudah terdaftar", role.Name)
	}

	return s.repo.CreateRole(role)
}

func (s *rbacService) GetAllRoles() ([]*domain.Role, error) {
	return s.repo.GetAllRoles()
}

func (s *rbacService) GetRoleByID(id string) (*domain.Role, error) {
	return s.repo.GetRoleByID(id)
}

func (s *rbacService) UpdateRole(role *domain.Role) error {
	role.Name = strings.TrimSpace(role.Name)

	// Validasi Duplikasi Nama
	exists, err := s.repo.CheckDuplicateRoleName(role.Name, role.ID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("role '%s' sudah terdaftar", role.Name)
	}

	return s.repo.UpdateRole(role)
}

func (s *rbacService) DeleteRole(id string) error {
	return s.repo.DeleteRole(id)
}

func (s *rbacService) AssignRoleToUser(schoolUserID string, roleID string) error {
	userRole := &domain.UserRole{
		SchoolUserID: schoolUserID,
		RoleID:       roleID,
	}
	return s.repo.AssignRole(userRole)
}

func (s *rbacService) RemoveRoleFromUser(schoolUserID string, roleID string) error {
	return s.repo.RemoveRoleFromUser(schoolUserID, roleID)
}

func (s *rbacService) GetUserRoles(schoolUserID string) ([]*domain.UserRole, error) {
	return s.repo.GetUserRoles(schoolUserID)
}

func (s *rbacService) SyncUserRoles(schoolUserID string, roleIDs []string) error {
	return s.repo.SyncUserRoles(schoolUserID, roleIDs)
}
