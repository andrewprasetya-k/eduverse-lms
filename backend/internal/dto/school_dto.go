package dto

type CreateSchoolDTO struct {
	Name      string         `json:"schoolName"`
	Code      string         ` json:"schoolCode"`
	LogoID    *string        `json:"schoolLogo,omitempty"`
	Address   string        `json:"schoolAddress"`
	Email     string        `json:"schoolEmail"`
	Phone     string        `json:"schoolPhone"`
	Website   *string        `json:"schoolWebsite,omitempty"`
}

type UpdateSchoolDTO struct {
	Name      *string         `json:"schoolName"`
	Code      *string         ` json:"schoolCode"`
	LogoID    *string        `json:"schoolLogo,omitempty"`
	Address   *string        `json:"schoolAddress"`
	Email     *string        `json:"schoolEmail"`
	Phone     *string        `json:"schoolPhone"`
	Website   *string        `json:"schoolWebsite,omitempty"`
}