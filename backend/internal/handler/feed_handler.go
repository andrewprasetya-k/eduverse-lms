package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FeedHandler struct {
	service        service.FeedService
	commentService service.CommentService
}

func NewFeedHandler(service service.FeedService, commentService service.CommentService) *FeedHandler {
	return &FeedHandler{
		service:        service,
		commentService: commentService,
	}
}

func (h *FeedHandler) Create(c *gin.Context) {
	var input dto.CreateFeedDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feed := domain.Feed{
		SchoolID:  input.SchoolID,
		ClassID:   input.ClassID,
		Content:   input.Content,
		CreatedBy: input.CreatedBy,
	}

	if err := h.service.Create(&feed, input.MediaIDs); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Feed posted"})
}

func (h *FeedHandler) GetByClass(c *gin.Context) {
	classID := c.Param("classId")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	feeds, total, err := h.service.GetByClass(classID, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.FeedResponseDTO
	for _, f := range feeds {
		count, _ := h.commentService.CountBySource(string(domain.SourceFeed), f.ID)
		response = append(response, h.mapToResponse(f, count))
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

func (h *FeedHandler) mapToResponse(f *domain.Feed, commentCount int) dto.FeedResponseDTO {
	var atts []dto.MediaResponseDTO
	for _, a := range f.Attachments {
		atts = append(atts, dto.MediaResponseDTO{
			ID:       a.Media.ID,
			Name:     a.Media.Name,
			FileSize: a.Media.FileSize,
			MimeType: a.Media.MimeType,
			FileURL:  a.Media.FileURL,
		})
	}

	return dto.FeedResponseDTO{
		ID:           f.ID,
		Content:      f.Content,
		CreatorName:  f.Creator.FullName,
		CreatedAt:    f.CreatedAt.Format("02-01-2006 15:04:05"),
		Attachments:  atts,
		CommentCount: commentCount,
	}
}
