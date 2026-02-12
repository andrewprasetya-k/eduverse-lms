package dto

type CreateSchoolDTO struct {
	Name      string         `json:"schoolName" binding:"required"`
	Code      string         `json:"schoolCode"`
	LogoID    *string        `json:"schoolLogo,omitempty"`
	Address   string        `json:"schoolAddress" binding:"required"`
	Email     string        `json:"schoolEmail" binding:"required"`
	Phone     string        `json:"schoolPhone" binding:"omitempty"`
	Website   *string        `json:"schoolWebsite,omitempty"`
}

type UpdateSchoolDTO struct {
	Name      *string         `json:"schoolName"`
	Code      *string         `json:"schoolCode"`
	LogoID    *string        `json:"schoolLogo,omitempty"`
	Address   *string        `json:"schoolAddress"`
	Email     *string        `json:"schoolEmail" binding:"required,email"`
	Phone     *string        `json:"schoolPhone" binding:"omitempty"`
	Website   *string        `json:"schoolWebsite" binding:"omitempty,url"`
}

type SchoolResponseDTO struct{
	ID string `json:"schoolId"`
	Name string `json:"schoolName"`
	Code string `json:"schoolCode"`
	LogoID    *string        `json:"schoolLogo,omitempty"`
	Address   string        `json:"schoolAddress"`
	Email     string        `json:"schoolEmail"`
	Phone     string        `json:"schoolPhone"`
	Website   *string        `json:"schoolWebsite,omitempty"`
	IsDeleted bool `json:"isDeleted"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}