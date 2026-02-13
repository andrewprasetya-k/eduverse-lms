package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubjectClassHandler struct {
	service service.SubjectClassService
}

func NewSubjectClassHandler(service service.SubjectClassService) *SubjectClassHandler {
	return &SubjectClassHandler{service: service}
}

func (h *SubjectClassHandler) Assign(c *gin.Context) {
	var input dto.CreateSubjectClassDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scl := domain.SubjectClass{
		ClassID:      input.ClassID,
		SubjectID:    input.SubjectID,
		SchoolUserID: input.SchoolUserID,
	}

	if err := h.service.Assign(&scl); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Subject and teacher assigned to class successfully"})
}

func (h *SubjectClassHandler) GetByClass(c *gin.Context) {
	classID := c.Param("classId")
	results, err := h.service.GetByClass(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.SubjectClassResponseDTO
	for _, r := range results {
		response = append(response, h.mapToResponse(r))
	}

	c.JSON(http.StatusOK, response)
}

func (h *SubjectClassHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	result, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject class assignment not found"})
		return
	}
	c.JSON(http.StatusOK, h.mapToResponse(result))
}

func (h *SubjectClassHandler) Unassign(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Unassign(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Assignment removed successfully"})
}

func (h *SubjectClassHandler) mapToResponse(scl *domain.SubjectClass) dto.SubjectClassResponseDTO {
	return dto.SubjectClassResponseDTO{
		ID:          scl.ID,
		ClassID:     scl.ClassID,
		SubjectID:   scl.SubjectID,
		SubjectName: scl.Subject.Name,
		SubjectCode: scl.Subject.Code,
		TeacherID:   scl.SchoolUserID,
		TeacherName: scl.Teacher.User.FullName,
	}
}
