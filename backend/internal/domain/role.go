package domain

import (
	"time"
)

type Role struct {
	ID          string       `gorm:"primaryKey;column:rol_id;default:gen_random_uuid()" json:"roleId"`
	SchoolID    string       `gorm:"column:rol_sch_id;type:uuid" json:"schoolId"`
	School      School       `gorm:"foreignKey:SchoolID;references:ID" json:"school,omitempty"`
	Name        string       `gorm:"column:rol_name" json:"roleName"`
	CreatedAt   time.Time    `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	Permissions []Permission `gorm:"many2many:edv.role_permissions;foreignKey:ID;joinForeignKey:rp_rol_id;References:ID;joinReferences:rp_prm_id" json:"permissions,omitempty"`
}

func (Role) TableName() string {
	return "edv.roles"
}

type RolePermission struct {
	ID           string `gorm:"primaryKey;column:rp_id;default:gen_random_uuid()" json:"rolePermissionId"`
	RoleID       string `gorm:"column:rp_rol_id;type:uuid"`
	PermissionID string `gorm:"column:rp_prm_id;type:uuid"`
}

func (RolePermission) TableName() string {
	return "edv.role_permissions"
}
