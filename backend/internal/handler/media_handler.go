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

// isStorageAvailable checks if the file storage backend is available
// Returns false if storage is not yet implemented
func isStorageAvailable() bool {
	// TODO: Replace with actual storage availability check (e.g., Supabase/S3 health check)
	// For now, always return false to prevent fake success responses
	return false
}

// Upload handles multipart file upload
func (h *MediaHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	schoolID := c.PostForm("schoolId")

	if schoolID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "schoolId is required"})
		return
	}

	// Auto-detect file info
	fileSize := file.Size / (1024 * 1024) // Convert to MB
	if fileSize > 10 {                    // Example limit: 10MB
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 10MB limit"})
		return
	}

	// Check if storage is available
	if !isStorageAvailable() {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "File storage service is not available. Please try again later"})
		return
	}

	// TODO: Upload to Supabase Storage here
	// For now, return error since real storage is not implemented
	c.JSON(http.StatusNotImplemented, gin.H{"error": "File upload to storage is not yet implemented"})
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
