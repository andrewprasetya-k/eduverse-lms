package domain

import (
	"time"

	"gorm.io/gorm"
)

type MaterialType string

const (
	MaterialVideo MaterialType = "video"
	MaterialPDF   MaterialType = "pdf"
	MaterialPPT   MaterialType = "ppt"
	MaterialOther MaterialType = "other"
)

type Material struct {
	ID          string         `gorm:"primaryKey;column:mat_id;default:gen_random_uuid()" json:"materialId"`
	SchoolID    string         `gorm:"column:mat_sch_id;type:uuid" json:"schoolId"`
	School      School         `gorm:"foreignKey:SchoolID;references:ID" json:"school,omitempty"`
	ClassID     string         `gorm:"column:mat_cls_id;type:uuid" json:"classId"`
	Class       Class          `gorm:"foreignKey:ClassID;references:ID" json:"class,omitempty"`
	Title       string         `gorm:"column:mat_title" json:"materialTitle"`
	Description string         `gorm:"column:mat_desc" json:"materialDescription"`
	Type        MaterialType   `gorm:"column:mat_types;type:material_type" json:"materialType"`
	CreatedBy   string         `gorm:"column:created_by;type:uuid" json:"createdBy"`
	Creator     User           `gorm:"foreignKey:CreatedBy;references:ID" json:"creator,omitempty"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	Attachments []Attachment   `gorm:"-" json:"attachments,omitempty"` // Loaded manually or via logic
}

func (Material) TableName() string {
	return "edv.materials"
}

type StatusProgress string

const (
	StatusNotStarted StatusProgress = "not_started"
	StatusCompleted  StatusProgress = "completed"
)

type MaterialProgress struct {
	ID           string         `gorm:"primaryKey;column:map_id;default:gen_random_uuid()" json:"progressId"`
	UserID       string         `gorm:"column:map_usr_id;type:uuid" json:"userId"`
	MaterialID   string         `gorm:"column:map_mat_id;type:uuid" json:"materialId"`
	Status       StatusProgress `gorm:"column:map_status;type:status_progress" json:"status"`
	LastOpenedAt *time.Time     `gorm:"column:last_opened_at" json:"lastOpenedAt"`
}

func (MaterialProgress) TableName() string {
	return "edv.material_progress"
}
