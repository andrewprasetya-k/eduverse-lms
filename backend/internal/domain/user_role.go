package domain

import (
	"time"
)

type UserRole struct {
	ID           string     `gorm:"primaryKey;column:urol_id;default:gen_random_uuid()" json:"userRoleId"`
	SchoolUserID string     `gorm:"column:urol_scu_id;type:uuid" json:"schoolUserId"`
	SchoolUser   SchoolUser `gorm:"foreignKey:SchoolUserID;references:ID" json:"schoolUser,omitempty"`
	RoleID       string     `gorm:"column:urol_rol_id;type:uuid" json:"roleId"`
	Role         Role       `gorm:"foreignKey:RoleID;references:ID" json:"role,omitempty"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (UserRole) TableName() string {
	return "edv.user_roles"
}
