package dto

type CreateTermDTO struct {
	AcademicYearID string `json:"academicYearId" binding:"required,uuid"`
	Name           string `json:"termName" binding:"required"`
}

type UpdateTermDTO struct {
	Name *string `json:"termName"`
}

type TermResponseDTO struct {
	ID               string `json:"termId"`
	AcademicYearID   string `json:"academicYearId"`
	AcademicYearName string `json:"academicYearName,omitempty"`
	SchoolName       string `json:"schoolName,omitempty"`
	Name             string `json:"termName"`
	IsActive         bool   `json:"isActive"`
	CreatedAt        string `json:"createdAt"`
}
