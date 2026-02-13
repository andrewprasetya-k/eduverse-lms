package dto

type RecordMediaDTO struct {
	SchoolID     string `json:"schoolId" binding:"required,uuid"`
	Name         string `json:"mediaName" binding:"required"`
	FileSize     int64  `json:"fileSize" binding:"required"`
	MimeType     string `json:"mimeType" binding:"required"`
	StoragePath  string `json:"storagePath" binding:"required"`
	FileURL      string `json:"fileUrl" binding:"required"`
	ThumbnailURL string `json:"thumbnailUrl"`
	IsPublic     bool   `json:"isPublic"`
	OwnerType    string `json:"ownerType" binding:"required"`
	OwnerID      string `json:"ownerId" binding:"required,uuid"`
}

type MediaResponseDTO struct {
	ID           string `json:"mediaId"`
	Name         string `json:"mediaName"`
	FileSize     int64  `json:"fileSize"`
	MimeType     string `json:"mimeType"`
	FileURL      string `json:"fileUrl"`
	ThumbnailURL string `json:"thumbnailUrl,omitempty"`
	OwnerType    string `json:"ownerType"`
	CreatedAt    string `json:"createdAt"`
}

type CreateAttachmentDTO struct {
	SchoolID   string `json:"schoolId" binding:"required,uuid"`
	SourceID   string `json:"sourceId" binding:"required,uuid"`
	SourceType string `json:"sourceType" binding:"required"`
	MediaID    string `json:"mediaId" binding:"required,uuid"`
}
