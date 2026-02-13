package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	service service.AssignmentService
}

func NewAssignmentHandler(service service.AssignmentService) *AssignmentHandler {
	return &AssignmentHandler{service: service}
}

func (h *AssignmentHandler) CreateCategory(c *gin.Context) {
	var input dto.CreateAssignmentCategoryDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cat := domain.AssignmentCategory{
		SchoolID: input.SchoolID,
		Name:     input.Name,
	}

	if err := h.service.CreateCategory(&cat); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {
	var input dto.CreateAssignmentDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asg := domain.Assignment{
		SchoolID:    input.SchoolID,
		ClassID:     input.ClassID,
		CategoryID:  input.CategoryID,
		Title:       input.Title,
		Description: input.Description,
		Deadline:    input.Deadline,
		CreatedBy:   input.CreatedBy,
	}

	if err := h.service.CreateAssignment(&asg, input.MediaIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Assignment created"})
}

func (h *AssignmentHandler) GetByClass(c *gin.Context) {
	classID := c.Param("classId")
	results, err := h.service.GetAssignmentsByClass(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.AssignmentResponseDTO
	for _, r := range results {
		response = append(response, h.mapAsgToResponse(r))
	}

	c.JSON(http.StatusOK, response)
}

func (h *AssignmentHandler) Submit(c *gin.Context) {
	var input dto.CreateSubmissionDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sbm := domain.Submission{
		SchoolID:     input.SchoolID,
		AssignmentID: input.AssignmentID,
		UserID:       input.UserID,
	}

	if err := h.service.Submit(&sbm, input.MediaIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Submission received"})
}

func (h *AssignmentHandler) Assess(c *gin.Context) {
	var input dto.CreateAssessmentDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	asm := domain.Assessment{
		SubmissionID: input.SubmissionID,
		Score:        input.Score,
		Feedback:     input.Feedback,
		AssessedBy:   input.AssessedBy,
	}

	if err := h.service.Assess(&asm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assessment recorded"})
}

func (h *AssignmentHandler) mapAsgToResponse(a *domain.Assignment) dto.AssignmentResponseDTO {
	var atts []dto.MediaResponseDTO
	for _, att := range a.Attachments {
		atts = append(atts, dto.MediaResponseDTO{
			ID:       att.Media.ID,
			Name:     att.Media.Name,
			FileURL:  att.Media.FileURL,
			MimeType: att.Media.MimeType,
		})
	}

	return dto.AssignmentResponseDTO{
		ID:           a.ID,
		Title:        a.Title,
		Description:  a.Description,
		CategoryName: a.Category.Name,
		Deadline:     a.Deadline,
		CreatedAt:    a.CreatedAt.Format("02-01-2006 15:04:05"),
		Attachments:  atts,
	}
}
