package dto

type CreateSchoolDTO struct {
	Name    string  `json:"name" binding:"required"`
	Code    string  `json:"code"` // Hapus binding:"required"
	LogoID  *string `json:"logo_id,omitempty"`
	Address string  `json:"address" binding:"required"`
	Email   string  `json:"email" binding:"required,email"`
	Phone   string  `json:"phone" binding:"required"`
	Website *string `json:"website,omitempty"`
}

type UpdateSchoolDTO struct {
	Name    *string `json:"name,omitempty"`
	Code    *string `json:"code,omitempty"`
	LogoID  *string `json:"logo_id,omitempty"`
	Address *string `json:"address,omitempty"`
	Email   *string `json:"email,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Website *string `json:"website,omitempty"`
}