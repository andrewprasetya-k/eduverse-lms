package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"
	"strings"
)

type RBACService interface {
	// Role management
	CreateRole(role *domain.Role, permissionIDs []string) error
	GetRolesBySchool(schoolCode string) ([]*domain.Role, error)
	GetRoleByID(id string) (*domain.Role, error)
	UpdateRole(role *domain.Role) error
	DeleteRole(id string) error
	SetRolePermissions(roleID string, permissionIDs []string) error

	// Permission management
	CreatePermission(permission *domain.Permission) error
	GetAllPermissions() ([]*domain.Permission, error)

	// User-Role management
	AssignRoleToUser(schoolUserID string, roleID string) error
	RemoveRoleFromUser(schoolUserID string, roleID string) error
	GetUserRoles(schoolUserID string) ([]*domain.UserRole, error)
	SyncUserRoles(schoolUserID string, roleIDs []string) error
}

type rbacService struct {
	repo          repository.RBACRepository
	schoolService SchoolService
}

func NewRBACService(repo repository.RBACRepository, schoolService SchoolService) RBACService {
	return &rbacService{
		repo:          repo,
		schoolService: schoolService,
	}
}

func (s *rbacService) CreateRole(role *domain.Role, permissionIDs []string) error {
	role.Name = strings.TrimSpace(role.Name)

	// 1. Validasi Duplikasi Nama Role per Sekolah
	exists, err := s.repo.CheckDuplicateRoleName(role.SchoolID, role.Name, "")
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("role '%s' sudah terdaftar di sekolah ini", role.Name)
	}

	// 2. Ambil permissions jika ada
	if len(permissionIDs) > 0 {
		perms, err := s.repo.GetPermissionsByIDs(permissionIDs)
		if err != nil {
			return err
		}
		role.Permissions = perms
	}

	return s.repo.CreateRole(role)
}

func (s *rbacService) GetRolesBySchool(schoolCode string) ([]*domain.Role, error) {
	schoolID, err := s.schoolService.ConvertCodeToID(schoolCode)
	if err != nil {
		return nil, err
	}
	return s.repo.GetRolesBySchool(schoolID)
}

func (s *rbacService) GetRoleByID(id string) (*domain.Role, error) {
	return s.repo.GetRoleByID(id)
}

func (s *rbacService) UpdateRole(role *domain.Role) error {
	role.Name = strings.TrimSpace(role.Name)

	// Validasi Duplikasi Nama
	exists, err := s.repo.CheckDuplicateRoleName(role.SchoolID, role.Name, role.ID)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("role '%s' sudah terdaftar di sekolah ini", role.Name)
	}

	return s.repo.UpdateRole(role)
}

func (s *rbacService) DeleteRole(id string) error {
	return s.repo.DeleteRole(id)
}

func (s *rbacService) SetRolePermissions(roleID string, permissionIDs []string) error {
	perms, err := s.repo.GetPermissionsByIDs(permissionIDs)
	if err != nil {
		return err
	}
	return s.repo.SetRolePermissions(roleID, perms)
}

func (s *rbacService) CreatePermission(permission *domain.Permission) error {
	permission.Key = strings.ToUpper(strings.TrimSpace(permission.Key))
	permission.Description = strings.TrimSpace(permission.Description)

	if permission.Key == "" {
		return fmt.Errorf("permission key cannot be empty")
	}

	return s.repo.CreatePermission(permission)
}

func (s *rbacService) GetAllPermissions() ([]*domain.Permission, error) {
	return s.repo.GetAllPermissions()
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
