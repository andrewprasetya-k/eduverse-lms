package dto

type SchoolBootstrapRequestDTO struct {
	School    CreateSchoolDTO       `json:"school" binding:"required"`
	AdminUser BootstrapAdminUserDTO `json:"adminUser" binding:"required"`
}

type BootstrapAdminUserDTO struct {
	Mode     string `json:"mode" binding:"required,oneof=new existing"`
	UserID   string `json:"userId" binding:"omitempty,uuid"`
	FullName string `json:"fullName"`
	Email    string `json:"email" binding:"omitempty,email"`
	Password string `json:"password"`
}

type SchoolBootstrapResponseDTO struct {
	School        SchoolBootstrapSchoolDTO      `json:"school"`
	AdminUser     BootstrapAdminUserResponseDTO `json:"adminUser"`
	SchoolUserID  string                        `json:"schoolUserId"`
	AssignedRoles []string                      `json:"assignedRoles"`
}

type SchoolBootstrapSchoolDTO struct {
	ID   string `json:"schoolId"`
	Name string `json:"schoolName"`
	Code string `json:"schoolCode"`
}

type BootstrapAdminUserResponseDTO struct {
	ID       string `json:"userId"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	IsActive bool   `json:"isActive"`
}
