package domain

import (
	"time"
)

type Term struct {
	ID             string       `gorm:"primaryKey;column:trm_id;default:gen_random_uuid()" json:"termId"`
	AcademicYearID string       `gorm:"column:trm_acy_id;type:uuid" json:"academicYearId"`
	AcademicYear   AcademicYear `gorm:"foreignKey:AcademicYearID;references:ID" json:"academicYear,omitempty"`
	Name           string       `gorm:"column:trm_name" json:"termName"`
	IsActive       bool         `gorm:"column:is_active;default:false" json:"isActive"`
	CreatedAt      time.Time    `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (Term) TableName() string {
	return "edv.terms"
}
