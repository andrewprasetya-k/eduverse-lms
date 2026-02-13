package domain

import (
	"time"
)

type AcademicYear struct {
	ID        string    `gorm:"primaryKey;column:acy_id;default:gen_random_uuid()" json:"academicYearId"`
	SchoolID  string    `gorm:"column:acy_sch_id;type:uuid" json:"schoolId"`
	Name      string    `gorm:"column:acy_name" json:"academicYearName"`
	IsActive  bool      `gorm:"column:is_active;default:false" json:"isActive"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (AcademicYear) TableName() string {
	return "edv.academic_years"
}
