package middleware

import (
	"backend/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var rbacRepo repository.RBACRepository

// InitRBAC initializes RBAC middleware with repository
func InitRBAC(repo repository.RBACRepository) {
	rbacRepo = repo
}

// RequireSchoolAccess checks if user belongs to the school in path
func RequireSchoolAccess(schoolService interface {
	ConvertCodeToID(code string) (string, error)
}) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := GetUserID(c)
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		schoolCode := c.Param("schoolCode")
		if schoolCode == "" {
			c.Next()
			return
		}

		schoolID, err := schoolService.ConvertCodeToID(schoolCode)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
			c.Abort()
			return
		}

		isMember, err := rbacRepo.IsUserInSchool(userID, schoolID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify school access"})
			c.Abort()
			return
		}

		if !isMember {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: not a member of this school"})
			c.Abort()
			return
		}

		// Store for later use
		c.Set("school_id", schoolID)
		c.Next()
	}
}

// RequireRole checks if user has any of the allowed roles in the school
func RequireRole(schoolService interface {
	ConvertCodeToID(code string) (string, error)
}, allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := GetUserID(c)
		if userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Get schoolID from context or param
		schoolID, exists := c.Get("school_id")
		if !exists {
			schoolCode := c.Param("schoolCode")
			if schoolCode != "" {
				sid, err := schoolService.ConvertCodeToID(schoolCode)
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
					c.Abort()
					return
				}
				schoolID = sid
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "School context required"})
				c.Abort()
				return
			}
		}

		roles, err := rbacRepo.GetUserRoleNamesInSchool(userID, schoolID.(string))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to verify roles"})
			c.Abort()
			return
		}

		// Check if user has any of the allowed roles
		hasRole := false
		for _, userRole := range roles {
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		if !hasRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("user_roles", roles)
		c.Next()
	}
}
