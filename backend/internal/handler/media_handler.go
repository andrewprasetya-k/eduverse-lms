package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"backend/internal/storage"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MediaHandler struct {
	service service.MediaService
}

func NewMediaHandler(service service.MediaService) *MediaHandler {
	return &MediaHandler{service: service}
}

// Upload handles multipart file upload
func (h *MediaHandler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	schoolID := c.PostForm("schoolId")
	ownerType := c.PostForm("ownerType")
	ownerID := c.PostForm("ownerId")

	if schoolID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "schoolId is required"})
		return
	}
	if _, err := uuid.Parse(schoolID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "schoolId must be a valid UUID"})
		return
	}

	// Auto-detect file info
	const maxUploadSize = 10 * 1024 * 1024
	if file.Size > maxUploadSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds 10MB limit"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read uploaded file"})
		return
	}
	defer src.Close()

	fileName := filepath.Base(file.Filename)
	ext := filepath.Ext(fileName)
	objectPath := fmt.Sprintf("schools/%s/%s%s", schoolID, uuid.NewString(), ext)
	mimeType := file.Header.Get("Content-Type")
	if strings.TrimSpace(mimeType) == "" {
		mimeType = "application/octet-stream"
	}

	media := domain.Media{
		SchoolID:    schoolID,
		Name:        fileName,
		FileSize:    file.Size,
		MimeType:    mimeType,
		StoragePath: objectPath,
		IsPublic:    true,
		OwnerType:   domain.OwnerType(ownerType),
		OwnerID:     ownerID,
	}

	if err := h.service.UploadAndRecord(c.Request.Context(), &media, src); err != nil {
		if errors.Is(err, storage.ErrNotImplemented) || errors.Is(err, storage.ErrUnavailable) {
			c.JSON(http.StatusNotImplemented, gin.H{"error": "File upload to storage is not configured"})
			return
		}
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "File uploaded successfully",
		"mediaId":     media.ID,
		"fileName":    fileName,
		"fileSize":    file.Size,
		"mimeType":    mimeType,
		"storagePath": objectPath,
		"fileUrl":     media.FileURL,
		"ext":         ext,
	})
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
	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Media record deleted"})
}
