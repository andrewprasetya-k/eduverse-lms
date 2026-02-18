package handler

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EnrollmentHandler struct {
	service service.EnrollmentService
}

func NewEnrollmentHandler(service service.EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{service: service}
}

func (h *EnrollmentHandler) Enroll(c *gin.Context) {
	var input dto.CreateEnrollmentDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Enroll(input.SchoolID, input.ClassID, input.SchoolUserIDs, input.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Users enrolled to class successfully"})
}

func (h *EnrollmentHandler) GetByClass(c *gin.Context) {
	classID := c.Param("classId")
	results, err := h.service.GetByClass(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.EnrollmentResponseDTO
	for _, r := range results {
		response = append(response, dto.EnrollmentResponseDTO{
			ID:           r.ID,
			SchoolID:     r.SchoolID,
			SchoolUserID: r.SchoolUserID,
			UserFullName: r.SchoolUser.User.FullName,
			UserEmail:    r.SchoolUser.User.Email,
			ClassID:      r.ClassID,
			Role:         r.Role,
			JoinedAt:     r.JoinedAt.Format("02-01-2006 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *EnrollmentHandler) GetByMember(c *gin.Context) {
	schoolUserID := c.Param("schoolUserId")
	results, err := h.service.GetByMember(schoolUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.EnrollmentResponseDTO
	for _, r := range results {
		response = append(response, dto.EnrollmentResponseDTO{
			ID:           r.ID,
			SchoolID:     r.SchoolID,
			SchoolUserID: r.SchoolUserID,
			ClassID:      r.ClassID,
			ClassTitle:   r.Class.Title,
			Role:         r.Role,
			JoinedAt:     r.JoinedAt.Format("02-01-2006 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *EnrollmentHandler) Unenroll(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Unenroll(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Enrollment removed successfully"})
}
