package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RBACHandler struct {
	service service.RBACService
}

func NewRBACHandler(service service.RBACService) *RBACHandler {
	return &RBACHandler{service: service}
}

// Role Handlers
func (h *RBACHandler) CreateRole(c *gin.Context) {
	var input dto.CreateRoleDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role := domain.Role{
		SchoolID: input.SchoolID,
		Name:     input.Name,
	}

	if err := h.service.CreateRole(&role, input.PermissionIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, h.mapRoleToResponse(&role))
}

func (h *RBACHandler) GetRolesBySchool(c *gin.Context) {
	schoolCode := c.Param("schoolCode")
	roles, err := h.service.GetRolesBySchool(schoolCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.RoleResponseDTO
	for _, r := range roles {
		response = append(response, h.mapRoleToResponse(r))
	}

	c.JSON(http.StatusOK, response)
}

func (h *RBACHandler) GetRoleByID(c *gin.Context) {
	id := c.Param("id")
	role, err := h.service.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	c.JSON(http.StatusOK, h.mapRoleToResponse(role))
}

func (h *RBACHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateRoleDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.service.GetRoleByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	if input.Name != nil {
		role.Name = *input.Name
	}

	if err := h.service.UpdateRole(role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.mapRoleToResponse(role))
}

func (h *RBACHandler) SetRolePermissions(c *gin.Context) {
	id := c.Param("id")
	var input dto.SetRolePermissionsDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SetRolePermissions(id, input.PermissionIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role permissions updated successfully"})
}

func (h *RBACHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteRole(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

// Permission Handlers
func (h *RBACHandler) GetAllPermissions(c *gin.Context) {
	perms, err := h.service.GetAllPermissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.PermissionResponseDTO
	for _, p := range perms {
		response = append(response, dto.PermissionResponseDTO{
			ID:          p.ID,
			Key:         p.Key,
			Description: p.Description,
		})
	}

	c.JSON(http.StatusOK, response)
}

// User-Role Handlers
func (h *RBACHandler) AssignRole(c *gin.Context) {
	var input dto.AssignRoleDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.AssignRoleToUser(input.SchoolUserID, input.RoleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role assigned successfully"})
}

func (h *RBACHandler) RemoveRole(c *gin.Context) {
	schoolUserID := c.Query("schoolUserId")
	roleID := c.Query("roleId")

	if err := h.service.RemoveRoleFromUser(schoolUserID, roleID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role removed successfully"})
}

func (h *RBACHandler) GetUserRoles(c *gin.Context) {
	schoolUserID := c.Param("schoolUserId")
	userRoles, err := h.service.GetUserRoles(schoolUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.RoleResponseDTO
	for _, ur := range userRoles {
		response = append(response, h.mapRoleToResponse(&ur.Role))
	}

	c.JSON(http.StatusOK, response)
}

// Helpers
func (h *RBACHandler) mapRoleToResponse(role *domain.Role) dto.RoleResponseDTO {
	var perms []dto.PermissionResponseDTO
	for _, p := range role.Permissions {
		perms = append(perms, dto.PermissionResponseDTO{
			ID:          p.ID,
			Key:         p.Key,
			Description: p.Description,
		})
	}

	return dto.RoleResponseDTO{
		ID:          role.ID,
		SchoolID:    role.SchoolID,
		SchoolName:  role.School.Name,
		Name:        role.Name,
		Permissions: perms,
		CreatedAt:   role.CreatedAt.Format("02-01-2006 15:04:05"),
	}
}
