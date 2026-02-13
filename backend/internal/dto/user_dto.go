package dto

type CreateUserDTO struct {
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserDTO struct {
	FullName *string `json:"fullName"`
	Email    *string `json:"email" binding:"omitempty,email"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required,min=6"`
}

type UserResponseDTO struct {
	ID        string `json:"userId"`
	FullName  string `json:"fullName"`
	Email     string `json:"email"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
}
