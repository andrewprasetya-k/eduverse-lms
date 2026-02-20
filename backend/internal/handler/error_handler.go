package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// HandleError centralizes error responses and masks internal details
func HandleError(c *gin.Context, err error) {
	if err == nil {
		return
	}

	// 1. Check for GORM Record Not Found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Data not found"})
		return
	}

	errStr := err.Error()

	// 2. Masking Database Constraint Errors (PostgreSQL patterns)
	
	// Foreign Key Violation (e.g., trying to use an ID that doesn't exist)
	if strings.Contains(errStr, "violates foreign key constraint") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reference"})
		return
	}

	// Unique Violation (e.g., duplicate code or email)
	if strings.Contains(errStr, "duplicate key value violates unique constraint") {
		c.JSON(http.StatusConflict, gin.H{"error": "Resource already exists"})
		return
	}

	// 3. Default Error (Internal Server Error)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
}

// HandleBindingError masks raw validation errors from Gin/Validator
func HandleBindingError(c *gin.Context, err error) {
	// Untuk saat ini kita buat pesan general, tapi bisa dikembangkan 
	// untuk parsing field mana yang error jika dibutuhkan.
	c.JSON(http.StatusBadRequest, gin.H{"error": "Input validation failed"})
}
