package domain

import (
	"time"
)

type Subject struct {
	ID        string         `gorm:"primaryKey;column:sub_id;default:gen_random_uuid()" json:"id"`
	SchoolID string `gorm:"column:sub_sch_id;type:uuid" json:"school_id"`
	Name      string         `gorm:"column:sub_name" json:"name"`
	Code      string         `gorm:"column:sub_code" json:"code"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`

	School School `gorm:"foreignKey:SchoolID;references:ID" json:"school,omitempty"`
}

func (Subject) TableName() string {
	return "edv.subjects"
}