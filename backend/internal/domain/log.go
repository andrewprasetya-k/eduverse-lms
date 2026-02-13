package domain

import (
	"time"
)

type Log struct {
	ID        string    `gorm:"primaryKey;column:log_id;default:gen_random_uuid()" json:"logId"`
	SchoolID  string    `gorm:"column:log_sch_id;type:uuid" json:"schoolId"`
	UserID    string    `gorm:"column:log_usr_id;type:uuid" json:"userId"`
	User      User      `gorm:"foreignKey:UserID;references:ID" json:"user,omitempty"`
	Action    string    `gorm:"column:log_action" json:"action"`
	Metadata  string    `gorm:"column:log_metadata;type:jsonb" json:"metadata"` // Stored as string for simplicity in basic impl
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (Log) TableName() string {
	return "edv.logs"
}
