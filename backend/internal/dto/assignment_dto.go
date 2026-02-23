package dto

import "time"

// Category
type CreateAssignmentCategoryDTO struct {
	SchoolID string `json:"schoolId" binding:"required,uuid"`
	Name     string `json:"categoryName" binding:"required"`
}

type AssignmentCategoryResponseDTO struct {
	ID        string `json:"categoryId"`
	SchoolID  string `json:"schoolId"`
	Name      string `json:"categoryName"`
	CreatedAt string `json:"createdAt"`
}

type SchoolWithAssignmentCategoriesDTO struct {
	School     SchoolHeaderDTO                 `json:"school"`
	Categories []AssignmentCategoryResponseDTO `json:"categories"`
}

// Assignment
type CreateAssignmentDTO struct {
	SchoolID            string     `json:"schoolId" binding:"required,uuid"`
	SubjectClassID      string     `json:"subjectClassId" binding:"required,uuid"`
	CategoryID          string     `json:"categoryId" binding:"required,uuid"`
	Title               string     `json:"assignmentTitle" binding:"required"`
	Description         string     `json:"assignmentDescription"`
	Deadline            *time.Time `json:"deadline"`
	AllowLateSubmission bool       `json:"allowLateSubmission"`
	CreatedBy           string     `json:"createdBy" binding:"required,uuid"`
	MediaIDs            []string   `json:"mediaIds"`
}

type UpdateAssignmentDTO struct {
	CategoryID          *string    `json:"categoryId" binding:"omitempty,uuid"`
	Title               *string    `json:"assignmentTitle"`
	Description         *string    `json:"assignmentDescription"`
	Deadline            *time.Time `json:"deadline"`
	AllowLateSubmission *bool      `json:"allowLateSubmission"`
	MediaIDs            []string   `json:"mediaIds"`
}

type AssignmentResponseDTO struct {
	ID                  string             `json:"assignmentId"`
	Title               string             `json:"assignmentTitle"`
	Description         string             `json:"assignmentDescription"`
	CategoryName        string             `json:"categoryName"`
	Deadline            *time.Time         `json:"deadline,omitempty"`
	AllowLateSubmission bool               `json:"allowLateSubmission"`
	CreatedAt           string             `json:"createdAt"`
	Attachments         []MediaResponseDTO `json:"attachments,omitempty"`
}

type AssignmentPerSubjectClassResponseDTO struct {
	SubjectClass  SubjectClassHeaderDTO `json:"subjectClass"`
	Assignments   []AssignmentResponseDTO `json:"assignments"`
}

type AssignmentHeaderDTO struct {
	ID           string     `json:"assignmentId"`
	Title        string     `json:"assignmentTitle"`
	SubjectName  string     `json:"subjectName"`
	CategoryName string     `json:"categoryName"`
	Deadline     *time.Time `json:"deadline,omitempty"`
}

type AssignmentWithSubmissionsDTO struct {
	Assignment  AssignmentHeaderDTO     `json:"assignment"`
	Submissions []SubmissionResponseDTO `json:"submissions"`
}

// Submission
type CreateSubmissionDTO struct {
	SchoolID     string   `json:"schoolId" binding:"required,uuid"`
	UserID       string   `json:"userId" binding:"required,uuid"`
	MediaIDs     []string `json:"mediaIds"`
}

type SubmissionResponseDTO struct {
	ID          string                 `json:"submissionId"`
	UserName    string                 `json:"studentName"`
	SubmittedAt string                 `json:"submittedAt"`
	IsLate      bool                   `json:"isLate"`
	Attachments []MediaResponseDTO     `json:"attachments,omitempty"`
	Assessment  *AssessmentResponseDTO `json:"assessment,omitempty"`
}

// Assessment
type CreateAssessmentDTO struct {
	Score        float64 `json:"score" binding:"required"`
	Feedback     string  `json:"feedback"`
	AssessedBy   string  `json:"assessedBy" binding:"required,uuid"`
}

type AssessmentResponseDTO struct {
	Score      float64 `json:"score"`
	Feedback   string  `json:"feedback"`
	Assessor   string  `json:"assessorName"`
	AssessedAt string  `json:"assessedAt"`
}

// Weight
type SetAssessmentWeightDTO struct {
	SubjectID  string  `json:"subjectId" binding:"required,uuid"`
	CategoryID string  `json:"categoryId" binding:"required,uuid"`
	Weight     float64 `json:"weight" binding:"required"`
}
