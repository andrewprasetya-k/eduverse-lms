package dto

import "time"

type SaveStudentNoteDTO struct {
	Content string `json:"content" binding:"required,max=10000"`
}

type StudentNoteResponseDTO struct {
	ID         string    `json:"noteId"`
	MaterialID string    `json:"materialId"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type StudentNoteEnvelopeDTO struct {
	Note *StudentNoteResponseDTO `json:"note"`
}

type StudentNoteCollectionItemDTO struct {
	ID            string    `json:"noteId"`
	MaterialID    string    `json:"materialId"`
	MaterialTitle string    `json:"materialTitle"`
	Content       string    `json:"content"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type StudentNoteCollectionDTO struct {
	Notes []StudentNoteCollectionItemDTO `json:"notes"`
}

type StudentGlobalNoteItemDTO struct {
	ID             string    `json:"noteId"`
	MaterialID     string    `json:"materialId"`
	MaterialTitle  string    `json:"materialTitle"`
	MaterialType   string    `json:"materialType"`
	SubjectClassID string    `json:"subjectClassId"`
	SubjectID      string    `json:"subjectId"`
	SubjectName    string    `json:"subjectName"`
	SubjectCode    string    `json:"subjectCode"`
	ClassID        string    `json:"classId"`
	ClassName      string    `json:"className"`
	ClassCode      string    `json:"classCode"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type StudentGlobalNotesDTO struct {
	Notes []StudentGlobalNoteItemDTO `json:"notes"`
}
