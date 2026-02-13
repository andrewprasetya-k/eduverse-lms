package dto

type CreateSubjectDTO struct {
	SchoolID string `json:"schoolId" binding:"required,uuid"`
	Name     string `json:"subjectName" binding:"required"`
	Code     string `json:"subjectCode" binding:"required"`
}

type UpdateSubjectDTO struct {
	Name *string `json:"subjectName"`
	Code *string `json:"subjectCode"`
}

type SubjectResponseDTO struct {
	ID         string `json:"subjectId"`
	SchoolID   string `json:"schoolId"`
	SchoolName string `json:"schoolName,omitempty"`
	SchoolCode string `json:"schoolCode,omitempty"`
	Name       string `json:"subjectName"`
	Code       string `json:"subjectCode"`
	CreatedAt  string `json:"createdAt"`
}

type SchoolWithSubjectsDTO struct {
	School   SchoolHeaderDTO      `json:"school"`
	Subjects []SubjectResponseDTO `json:"subjects"`
}
