package dto

type AdminSchoolMemberImportRowDTO struct {
	RowNumber int      `json:"rowNumber"`
	FullName  string   `json:"fullName"`
	Email     string   `json:"email"`
	Role      string   `json:"role"`
	ClassCode string   `json:"classCode,omitempty"`
	Status    string   `json:"status"`
	Errors    []string `json:"errors"`
}

type AdminSchoolMemberImportPreviewResponseDTO struct {
	Rows         []AdminSchoolMemberImportRowDTO `json:"rows"`
	ValidCount   int                             `json:"validCount"`
	InvalidCount int                             `json:"invalidCount"`
}

type AdminSchoolMemberImportCommitRequestDTO struct {
	DefaultPassword string                          `json:"defaultPassword" binding:"required,min=6"`
	Rows            []AdminSchoolMemberImportRowDTO `json:"rows" binding:"required"`
}

type AdminSchoolMemberImportResultDTO struct {
	RowNumber int    `json:"rowNumber"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	ClassCode string `json:"classCode,omitempty"`
	Status    string `json:"status"`
	Reason    string `json:"reason,omitempty"`
}

type AdminSchoolMemberImportCommitResponseDTO struct {
	ImportedCount int                                `json:"importedCount"`
	SkippedCount  int                                `json:"skippedCount"`
	FailedCount   int                                `json:"failedCount"`
	Results       []AdminSchoolMemberImportResultDTO `json:"results"`
}
