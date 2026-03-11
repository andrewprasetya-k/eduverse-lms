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

// RequireSchoolAccess checks if user belongs to the school
// Priority: SchoolId header > schoolCode URL param
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

		var schoolID string
		var err error

		// Priority 1: Check SchoolId header
		schoolID = c.GetHeader("SchoolId")

		// Priority 2: Check schoolCode in URL param
		if schoolID == "" {
			schoolCode := c.Param("schoolCode")
			if schoolCode == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "School context required (SchoolId header or schoolCode param)"})
				c.Abort()
				return
			}
			schoolID, err = schoolService.ConvertCodeToID(schoolCode)
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
				c.Abort()
				return
			}
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
// Priority: context > SchoolId header > schoolCode URL param
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

		var schoolID string
		var err error

		// Priority 1: Get from context (set by RequireSchoolAccess)
		if sid, exists := c.Get("school_id"); exists {
			schoolID = sid.(string)
		} else {
			// Priority 2: Check SchoolId header
			schoolID = c.GetHeader("SchoolId")

			// Priority 3: Check schoolCode in URL param
			if schoolID == "" {
				schoolCode := c.Param("schoolCode")
				if schoolCode != "" {
					schoolID, err = schoolService.ConvertCodeToID(schoolCode)
					if err != nil {
						c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
						c.Abort()
						return
					}
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "School context required (SchoolId header or schoolCode param)"})
					c.Abort()
					return
				}
			}
		}

		roles, err := rbacRepo.GetUserRoleNamesInSchool(userID, schoolID)
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
