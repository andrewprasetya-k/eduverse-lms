package dto

type CreateEnrollmentDTO struct {
	SchoolID      string   `json:"schoolId" binding:"required,uuid"`
	SchoolUserIDs []string `json:"schoolUserIds" binding:"required,dive,uuid"`
	ClassID       string   `json:"classId" binding:"required,uuid"`
	Role          string   `json:"role" binding:"required,oneof=teacher student"`
}

type EnrollmentResponseDTO struct {
	ID           string `json:"enrollmentId"`
	SchoolID     string `json:"schoolId"`
	SchoolUserID string `json:"schoolUserId"`
	UserFullName string `json:"userFullName,omitempty"`
	UserEmail    string `json:"userEmail,omitempty"`
	ClassID      string `json:"classId"`
	ClassTitle   string `json:"classTitle,omitempty"`
	Role         string `json:"role"`
	JoinedAt     string `json:"joinedAt"`
}

type ClassWithMembersDTO struct {
	School  SchoolHeaderDTO         `json:"school"`
	Members []EnrollmentResponseDTO `json:"members"`
}
