package domain

import (
	"time"

	"gorm.io/gorm"
)

type Feed struct {
	ID          string         `gorm:"primaryKey;column:fds_id;default:gen_random_uuid()" json:"feedId"`
	SchoolID    string         `gorm:"column:fds_sch_id;type:uuid" json:"schoolId"`
	School      School         `gorm:"foreignKey:SchoolID;references:ID" json:"school,omitempty"`
	ClassID     string         `gorm:"column:fds_cls_id;type:uuid" json:"classId"`
	Class       Class          `gorm:"foreignKey:ClassID;references:ID" json:"class,omitempty"`
	Content     string         `gorm:"column:fds_content" json:"content"`
	CreatedBy   string         `gorm:"column:created_by;type:uuid" json:"createdBy"`
	Creator     User           `gorm:"foreignKey:CreatedBy;references:ID" json:"creator,omitempty"`
	CreatedAt   time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	Attachments []Attachment   `gorm:"-" json:"attachments,omitempty"`
	Comments    []Comment      `gorm:"-" json:"comments,omitempty"`
}

func (Feed) TableName() string {
	return "edv.feeds"
}
