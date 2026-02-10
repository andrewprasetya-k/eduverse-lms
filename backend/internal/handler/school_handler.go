package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SchoolHandler struct {
	service service.SchoolService
}

func NewSchoolHandler(service service.SchoolService) *SchoolHandler {
	return &SchoolHandler{service: service}
}

// Create
func (h *SchoolHandler) CreateSchool(c *gin.Context) {
	var input dto.CreateSchoolDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	school := domain.School{
		Name:    input.Name,
		Code:    input.Code,
		LogoID:  input.LogoID,
		Address: input.Address,
		Email:   input.Email,
		Phone:   input.Phone,
		Website: input.Website,
	}

	if err := h.service.CreateSchool(&school); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, school)
}

// Get All
func (h *SchoolHandler) GetAllSchools(c *gin.Context) {
	schools, err := h.service.GetAllSchools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}

// Get By Code
func (h *SchoolHandler) GetSchoolByCode(c *gin.Context) {
	schoolCode := c.Param("schoolCode")
	school, err := h.service.GetSchoolByCode(schoolCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}
	c.JSON(http.StatusOK, school)
}

// Update
func (h *SchoolHandler) UpdateSchool(c *gin.Context) {
	var input dto.UpdateSchoolDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schoolCode := c.Param("schoolCode")
	school, err := h.service.GetSchoolByCode(schoolCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

	if input.Name != nil {
		school.Name = *input.Name
	}
	if input.Code != nil {
		school.Code = *input.Code
	}
	if input.LogoID != nil {
		school.LogoID = input.LogoID
	}
	if input.Address != nil {
		school.Address = *input.Address
	}
	if input.Email != nil {
		school.Email = *input.Email
	}
	if input.Phone != nil {
		school.Phone = *input.Phone
	}
	if input.Website != nil {
		school.Website = input.Website
	}

	if err := h.service.UpdateSchool(school); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, school)
}

// Delete
func (h *SchoolHandler) DeleteSchool(c *gin.Context) {
	schoolCode := c.Param("schoolCode")
	if err := h.service.DeleteSchool(schoolCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "School deleted successfully"})
}
