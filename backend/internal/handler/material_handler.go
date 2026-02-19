package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	service service.MaterialService
}

func NewMaterialHandler(service service.MaterialService) *MaterialHandler {
	return &MaterialHandler{service: service}
}

func (h *MaterialHandler) Create(c *gin.Context) {
	var input dto.CreateMaterialDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mat := domain.Material{
		SchoolID:    input.SchoolID,
		ClassID:     input.ClassID,
		Title:       input.Title,
		Description: input.Description,
		Type:        domain.MaterialType(input.Type),
		CreatedBy:   input.CreatedBy,
	}

	if err := h.service.Create(&mat, input.MediaIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Material created successfully"})
}

func (h *MaterialHandler) FindAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	classID := c.Query("classId")

	materials, total, err := h.service.FindAll(search, classID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.MaterialResponseDTO
	for _, m := range materials {
		response = append(response, h.mapToResponse(m))
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

func (h *MaterialHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	mat, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}
	c.JSON(http.StatusOK, h.mapToResponse(mat))
}

func (h *MaterialHandler) UpdateProgress(c *gin.Context) {
	var input dto.UpdateProgressDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.UpdateProgress(input.UserID, input.MaterialID, input.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Progress updated"})
}

func (h *MaterialHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input dto.UpdateMaterialDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mat, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}

	if input.Title != nil {
		mat.Title = *input.Title
	}
	if input.Description != nil {
		mat.Description = *input.Description
	}
	if input.Type != nil {
		mat.Type = domain.MaterialType(*input.Type)
	}

	if err := h.service.Update(mat, input.MediaIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Material updated successfully"})
}

func (h *MaterialHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Material deleted successfully"})
}

func (h *MaterialHandler) mapToResponse(m *domain.Material) dto.MaterialResponseDTO {
	var atts []dto.MediaResponseDTO
	for _, a := range m.Attachments {
		atts = append(atts, dto.MediaResponseDTO{
			ID:       a.Media.ID,
			Name:     a.Media.Name,
			FileSize: a.Media.FileSize,
			MimeType: a.Media.MimeType,
			FileURL:  a.Media.FileURL,
		})
	}

	return dto.MaterialResponseDTO{
		ID:          m.ID,
		ClassID:     m.ClassID,
		ClassTitle:  m.Class.Title,
		Title:       m.Title,
		Description: m.Description,
		Type:        string(m.Type),
		CreatorName: m.Creator.FullName,
		CreatedAt:   m.CreatedAt.Format("02-01-2006 15:04:05"),
		Attachments: atts,
	}
}
