package handler

import (
	"errors"
	"fmt"
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

	// Print raw error to server logs for debugging
	fmt.Printf("[Error Log] %s\n", err.Error())

	// 1. Check for GORM Record Not Found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": "The requested data was not found"})
		return
	}

	errStr := err.Error()

	// 2. Masking Database Constraint Errors (PostgreSQL patterns)
	
	// Foreign Key Violation
	if strings.Contains(errStr, "violates foreign key constraint") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or non-existent data reference"})
		return
	}

	// Unique Violation
	if strings.Contains(errStr, "duplicate key value violates unique constraint") {
		c.JSON(http.StatusConflict, gin.H{"error": "This data already exists in the system"})
		return
	}

	// 3. Default Error (Internal Server Error)
	c.JSON(http.StatusInternalServerError, gin.H{"error": "An internal server error occurred"})
}

// HandleBindingError masks raw validation errors from Gin/Validator
func HandleBindingError(c *gin.Context, err error) {
	// Print raw binding error to server logs for debugging
	fmt.Printf("[Binding Error Log] %s\n", err.Error())

	errStr := err.Error()

	if strings.Contains(errStr, "failed on the 'required' tag") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Required fields are missing"})
		return
	}
	if strings.Contains(errStr, "failed on the 'uuid' tag") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}
	if strings.Contains(errStr, "failed on the 'email' tag") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}
	
	// Type mismatch errors (e.g. sending string instead of number)
	if strings.Contains(errStr, "unmarshal") || strings.Contains(errStr, "type mismatch") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data type mismatch. Please check your input values"})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data format. Please check your request"})
}
