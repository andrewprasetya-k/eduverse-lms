package dto

type CreateAcademicYearDTO struct{
	SchoolID string `json:"school_id" binding:"required,uuid"`
	Name string `json:"name" binding:"required"`
	IsActive bool `json:"is_active"`
}

type CreateTermDTO struct{
	AcademicYearID string `json:"academic_year_id" binding:"required,uuid"`
	Name string `json:"name" binding:"required"`
	IsActive bool `json:"is_active"`
}

type UpdateAcademicYearDTO struct{
	Name *string `json:"name"`
	IsActive *bool `json:"is_active"`
}