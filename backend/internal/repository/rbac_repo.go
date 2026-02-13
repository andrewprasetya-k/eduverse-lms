package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type RBACRepository interface {
	// Role operations
	CreateRole(role *domain.Role) error
	GetRoleByID(id string) (*domain.Role, error)
	GetRolesBySchool(schoolID string) ([]*domain.Role, error)
	UpdateRole(role *domain.Role) error
	DeleteRole(id string) error
	CheckDuplicateRoleName(schoolID string, name string, excludeID string) (bool, error)

	// Permission operations
	GetAllPermissions() ([]*domain.Permission, error)
	GetPermissionsByIDs(ids []string) ([]domain.Permission, error)

	// Role-Permission association
	SetRolePermissions(roleID string, permissions []domain.Permission) error

	// User-Role association
	AssignRole(userRole *domain.UserRole) error
	RemoveRoleFromUser(schoolUserID string, roleID string) error
	GetUserRoles(schoolUserID string) ([]*domain.UserRole, error)
}

type rbacRepository struct {
	db *gorm.DB
}

func NewRBACRepository(db *gorm.DB) RBACRepository {
	return &rbacRepository{db: db}
}

func (r *rbacRepository) CreateRole(role *domain.Role) error {
	return r.db.Create(role).Error
}

func (r *rbacRepository) GetRoleByID(id string) (*domain.Role, error) {
	var role domain.Role
	err := r.db.Preload("Permissions").Preload("School").Where("rol_id = ?", id).First(&role).Error
	return &role, err
}

func (r *rbacRepository) GetRolesBySchool(schoolID string) ([]*domain.Role, error) {
	var roles []*domain.Role
	err := r.db.Preload("Permissions").Where("rol_sch_id = ?", schoolID).Find(&roles).Error
	return roles, err
}

func (r *rbacRepository) UpdateRole(role *domain.Role) error {
	result := r.db.Save(role)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *rbacRepository) DeleteRole(id string) error {
	// First delete associations in role_permissions (GORM usually handles this if configured, but explicit is safer)
	err := r.db.Exec("DELETE FROM edv.role_permissions WHERE rp_rol_id = ?", id).Error
	if err != nil {
		return err
	}
	
	result := r.db.Delete(&domain.Role{}, "rol_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *rbacRepository) CheckDuplicateRoleName(schoolID string, name string, excludeID string) (bool, error) {
	var count int64
	query := r.db.Model(&domain.Role{}).Where("rol_sch_id = ? AND rol_name = ?", schoolID, name)
	if excludeID != "" {
		query = query.Where("rol_id != ?", excludeID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}

func (r *rbacRepository) GetAllPermissions() ([]*domain.Permission, error) {
	var perms []*domain.Permission
	err := r.db.Find(&perms).Error
	return perms, err
}

func (r *rbacRepository) GetPermissionsByIDs(ids []string) ([]domain.Permission, error) {
	var perms []domain.Permission
	err := r.db.Where("prm_id IN ?", ids).Find(&perms).Error
	return perms, err
}

func (r *rbacRepository) SetRolePermissions(roleID string, permissions []domain.Permission) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Delete existing associations
		if err := tx.Exec("DELETE FROM edv.role_permissions WHERE rp_rol_id = ?", roleID).Error; err != nil {
			return err
		}
		
		// 2. Create new associations
		if len(permissions) > 0 {
			var rolePermissions []domain.RolePermission
			for _, p := range permissions {
				rolePermissions = append(rolePermissions, domain.RolePermission{
					RoleID:       roleID,
					PermissionID: p.ID,
				})
			}
			if err := tx.Create(&rolePermissions).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *rbacRepository) AssignRole(userRole *domain.UserRole) error {
	return r.db.Create(userRole).Error
}

func (r *rbacRepository) RemoveRoleFromUser(schoolUserID string, roleID string) error {
	result := r.db.Where("urol_scu_id = ? AND urol_rol_id = ?", schoolUserID, roleID).Delete(&domain.UserRole{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *rbacRepository) GetUserRoles(schoolUserID string) ([]*domain.UserRole, error) {
	var userRoles []*domain.UserRole
	err := r.db.Preload("Role.Permissions").Where("urol_scu_id = ?", schoolUserID).Find(&userRoles).Error
	return userRoles, err
}
