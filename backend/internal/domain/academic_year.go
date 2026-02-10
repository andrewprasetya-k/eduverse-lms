package domain

import "time"

type AcademicYear struct {
	ID        string         `gorm:"primaryKey;column:acy_id;default:gen_random_uuid()" json:"id"`
	SchoolID string 	   `gorm:"column:acy_sch_id;type:uuid" json:"school_id"`
	Name      string         `gorm:"column:acy_name" json:"name"`
	IsActive  bool           `gorm:"column:acy_is_active;default:true" json:"is_active"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	Terms []Term		 `gorm:"foreignKey:AcademicYearID" json:"terms,omitempty"`
}

type Term struct {
	ID             string         `gorm:"primaryKey;column:trm_id;default:gen_random_uuid()" json:"id"`
	AcademicYearID string         `gorm:"column:trm_acy_id;type:uuid" json:"academic_year_id"`
	Name string `gorm:"column:trm_name" json:"name"`
	IsActive bool `gorm:"column:is_active;default:false" json:"is_active"`
	CreatedAt      time.Time      `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (AcademicYear) TableName() string {
	return "edv.academic_years"
}

func (Term) TableName() string {
	return "edv.terms"
}