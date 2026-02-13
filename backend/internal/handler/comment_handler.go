package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	service service.CommentService
}

func NewCommentHandler(service service.CommentService) *CommentHandler {
	return &CommentHandler{service: service}
}

func (h *CommentHandler) Create(c *gin.Context) {
	var input dto.CreateCommentDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment := domain.Comment{
		SchoolID:   input.SchoolID,
		SourceType: domain.SourceType(input.SourceType),
		SourceID:   input.SourceID,
		UserID:     input.UserID,
		Content:    input.Content,
	}

	if err := h.service.Create(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment posted"})
}

func (h *CommentHandler) GetBySource(c *gin.Context) {
	sourceType := c.Query("type")
	sourceID := c.Query("id")

	comments, err := h.service.GetBySource(sourceType, sourceID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.CommentResponseDTO
	for _, c := range comments {
		response = append(response, dto.CommentResponseDTO{
			ID:          c.ID,
			Content:     c.Content,
			CreatorName: c.User.FullName,
			CreatedAt:   c.CreatedAt.Format("02-01-2006 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *CommentHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
