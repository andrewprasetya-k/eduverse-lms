package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	service service.SubjectService
}

func NewSubjectHandler(service service.SubjectService) *SubjectHandler {
	return &SubjectHandler{service: service}
}

func (h *SubjectHandler) Create(c *gin.Context) {
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

	if err := h.service.Create(&subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, h.mapToResponse(&subject))
}

func (h *SubjectHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")

	subjects, total, err := h.service.FindAll(search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.SubjectResponseDTO
	for _, s := range subjects {
		response = append(response, h.mapToResponse(s))
	}

	totalPages := (total + int64(limit) - 1) / int64(limit)

	paginatedResponse := dto.PaginatedResponse{
		Data:       response,
		TotalItems: total,
		Page:       page,
		Limit:      limit,
		TotalPages: int(totalPages),
	}
	c.JSON(http.StatusOK, paginatedResponse)
}

func (h *SubjectHandler) GetBySchool(c *gin.Context) {
	schoolCode := c.Param("schoolCode")
	subjects, err := h.service.GetBySchool(schoolCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.SubjectResponseDTO
	for _, s := range subjects {
		response = append(response, h.mapToResponse(s))
	}

	c.JSON(http.StatusOK, response)
}

func (h *SubjectHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	subject, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Subject not found"})
		return
	}
	c.JSON(http.StatusOK, h.mapToResponse(subject))
}

func (h *SubjectHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateSubjectDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subject, err := h.service.GetByID(id)
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

	if err := h.service.Update(subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.mapToResponse(subject))
}

func (h *SubjectHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}

func (h *SubjectHandler) mapToResponse(s *domain.Subject) dto.SubjectResponseDTO {
	return dto.SubjectResponseDTO{
		ID:         s.ID,
		SchoolID:   s.SchoolID,
		SchoolName: s.School.Name,
		SchoolCode: s.School.Code,
		Name:       s.Name,
		Code:       s.Code,
		CreatedAt:  s.CreatedAt.Format("02-01-2006 15:04:05"),
	}
}
