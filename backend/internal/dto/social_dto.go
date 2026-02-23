package dto

// Feed DTOs
type CreateFeedDTO struct {
	SchoolID  string   `json:"schoolId" binding:"required,uuid"`
	ClassID   string   `json:"classId" binding:"required,uuid"`
	Content   string   `json:"content" binding:"required"`
	CreatedBy string   `json:"createdBy" binding:"required,uuid"`
	MediaIDs  []string `json:"mediaIds"`
}

type FeedResponseDTO struct {
	ID           string             `json:"feedId"`
	Content      string             `json:"content"`
	CreatorName  string             `json:"creatorName,omitempty"`
	CreatedAt    string             `json:"createdAt"`
	Attachments  []MediaResponseDTO `json:"attachments,omitempty"`
	CommentCount int                `json:"commentCount"`
}

type ClassWithFeedsDTO struct {
	Class ClassHeaderDTO    `json:"class"`
	Data  PaginatedResponse `json:"data"`
}

// Comment DTOs
type CreateCommentDTO struct {
	SchoolID   string `json:"schoolId" binding:"required,uuid"`
	SourceType string `json:"sourceType" binding:"required"`
	SourceID   string `json:"sourceId" binding:"required,uuid"`
	UserID     string `json:"userId" binding:"required,uuid"`
	Content    string `json:"content" binding:"required"`
}

type CommentResponseDTO struct {
	ID          string `json:"commentId"`
	Content     string `json:"content"`
	CreatorName string `json:"creatorName"`
	CreatedAt   string `json:"createdAt"`
}
