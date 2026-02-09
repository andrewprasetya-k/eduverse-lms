package dto

type CreateSubjectDTO struct {
	SchoolID string `json:"school_id" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	Code    string  `json:"code"` // Hapus binding:"required"
}

type UpdateSubjectDTO struct {
	Name    *string `json:"name,omitempty"`
	Code    *string `json:"code,omitempty"`
}