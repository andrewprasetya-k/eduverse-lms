package dto

// Role DTOs
type CreateRoleDTO struct {
	SchoolID string   `json:"schoolId" binding:"required,uuid"`
	Name     string   `json:"roleName" binding:"required"`
	PermissionIDs []string `json:"permissionIds"` // Optional initial permissions
}

type UpdateRoleDTO struct {
	Name *string `json:"roleName"`
}

type RoleResponseDTO struct {
	ID          string                  `json:"roleId"`
	Name        string                  `json:"roleName"`
	Permissions []PermissionResponseDTO `json:"permissions,omitempty"`
	CreatedAt   string                  `json:"createdAt"`
}

// Permission DTOs
type CreatePermissionDTO struct {
	Key         string `json:"permissionKey" binding:"required"`
	Description string `json:"permissionDesc" binding:"required"`
}

type PermissionResponseDTO struct {
	ID          string `json:"permissionId"`
	Key         string `json:"permissionKey"`
	Description string `json:"description"`
}

// Assignment DTOs
type AssignRoleDTO struct {
	SchoolUserID string `json:"schoolUserId" binding:"required,uuid"`
	RoleID       string `json:"roleId" binding:"required,uuid"`
}

type SyncUserRolesDTO struct {
	RoleIDs []string `json:"roleIds" binding:"required"`
}

type SetRolePermissionsDTO struct {
	PermissionIDs []string `json:"permissionIds" binding:"required"`
}
