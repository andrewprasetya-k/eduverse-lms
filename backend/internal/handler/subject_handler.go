package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	service service.SubjectService
}

func NewSubjectHandler(service service.SubjectService) *SubjectHandler {
	return &SubjectHandler{service: service}
}

// Create
func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var input dto.CreateSubjectDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject := domain.Subject{
		SchoolID: input.SchoolID,
		Name:     input.Name,
		Code:     input.Code,
	}

	if err := h.service.CreateSubject(&subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, subject)
}

// Get All
func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {
	schoolID := c.Query("school_id")
	if schoolID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "school_id query parameter is required"})
		return
	}

	subjects, err := h.service.GetAllSubjects(schoolID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}

// Get By ID
func (h *SubjectHandler) GetSubjectByID(c *gin.Context) {
	id := c.Param("id")
	subject, err := h.service.GetSubjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}
	c.JSON(http.StatusOK, subject)
}

// Update
func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	var input dto.UpdateSubjectDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	subject, err := h.service.GetSubjectByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}

	if input.Name != nil {
		subject.Name = *input.Name
	}
	if input.Code != nil {
		subject.Code = *input.Code
	}

	if err := h.service.UpdateSubject(subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subject)
}

// Delete
func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteSubject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}