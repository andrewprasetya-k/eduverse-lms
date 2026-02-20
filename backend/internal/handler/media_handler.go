package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MediaHandler struct {
	service service.MediaService
}

func NewMediaHandler(service service.MediaService) *MediaHandler {
	return &MediaHandler{service: service}
}

// RecordMetadata records metadata of an already uploaded file (e.g., to Supabase/S3)
func (h *MediaHandler) RecordMetadata(c *gin.Context) {
	var input dto.RecordMediaDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	media := domain.Media{
		SchoolID:     input.SchoolID,
		Name:         input.Name,
		FileSize:     input.FileSize,
		MimeType:     input.MimeType,
		StoragePath:  input.StoragePath,
		FileURL:      input.FileURL,
		ThumbnailURL: input.ThumbnailURL,
		IsPublic:     input.IsPublic,
		OwnerType:    domain.OwnerType(input.OwnerType),
		OwnerID:      input.OwnerID,
	}

	if err := h.service.RecordMetadata(&media); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, media)
}

func (h *MediaHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	media, err := h.service.GetByID(id)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, media)
}

func (h *MediaHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Media record deleted"})
}
