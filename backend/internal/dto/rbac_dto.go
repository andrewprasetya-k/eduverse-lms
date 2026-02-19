package dto

// Role DTOs
type CreateRoleDTO struct {
	Name string `json:"roleName" binding:"required"`
}

type UpdateRoleDTO struct {
	Name *string `json:"roleName"`
}

type RoleResponseDTO struct {
	ID        string `json:"roleId"`
	Name      string `json:"roleName"`
	CreatedAt string `json:"createdAt"`
}

// Assignment DTOs
type AssignRoleDTO struct {
	SchoolUserID string `json:"schoolUserId" binding:"required,uuid"`
	RoleID       string `json:"roleId" binding:"required,uuid"`
}

type SyncUserRolesDTO struct {
	RoleIDs []string `json:"roleIds" binding:"required"`
}
