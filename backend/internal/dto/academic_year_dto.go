package dto

type CreateAcademicYearDTO struct {
	SchoolID string `json:"schoolId" binding:"required,uuid"`
	Name     string `json:"academicYearName" binding:"required"`
	IsActive bool   `json:"isActive"`
}

type UpdateAcademicYearDTO struct {
	Name     *string `json:"academicYearName"`
	IsActive *bool   `json:"isActive"`
}

type AcademicYearResponseDTO struct {
	ID        string `json:"academicYearId"`
	SchoolID  string `json:"schoolId"`
	Name      string `json:"academicYearName"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
}
