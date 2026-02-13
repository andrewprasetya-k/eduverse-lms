package dto

type CreateMaterialDTO struct {
	SchoolID    string   `json:"schoolId" binding:"required,uuid"`
	ClassID     string   `json:"classId" binding:"required,uuid"`
	Title       string   `json:"materialTitle" binding:"required"`
	Description string   `json:"materialDescription"`
	Type        string   `json:"materialType" binding:"required,oneof=video pdf ppt other"`
	CreatedBy   string   `json:"createdBy" binding:"required,uuid"`
	MediaIDs    []string `json:"mediaIds"` // Files to attach
}

type UpdateMaterialDTO struct {
	Title       *string `json:"materialTitle"`
	Description *string `json:"materialDescription"`
}

type MaterialResponseDTO struct {
	ID          string             `json:"materialId"`
	ClassID     string             `json:"classId"`
	ClassTitle  string             `json:"classTitle,omitempty"`
	Title       string             `json:"materialTitle"`
	Description string             `json:"materialDescription"`
	Type        string             `json:"materialType"`
	CreatorName string             `json:"creatorName,omitempty"`
	CreatedAt   string             `json:"createdAt"`
	Attachments []MediaResponseDTO `json:"attachments,omitempty"`
}

type UpdateProgressDTO struct {
	UserID     string `json:"userId" binding:"required,uuid"`
	MaterialID string `json:"materialId" binding:"required,uuid"`
	Status     string `json:"status" binding:"required,oneof=not_started completed"`
}
