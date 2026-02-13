package domain

import (
	"time"

	"gorm.io/gorm"
)

type OwnerType string

const (
	OwnerUser       OwnerType = "user"
	OwnerMaterial   OwnerType = "material"
	OwnerAssignment OwnerType = "assignment"
	OwnerFeed       OwnerType = "feed"
	OwnerSubmission OwnerType = "submission"
	OwnerComment    OwnerType = "comment"
	OwnerSchool     OwnerType = "school"
	OwnerSystem     OwnerType = "system"
)

type Media struct {
	ID           string         `gorm:"primaryKey;column:med_id;default:gen_random_uuid()" json:"mediaId"`
	SchoolID     string         `gorm:"column:med_sch_id;type:uuid" json:"schoolId"`
	School       School         `gorm:"foreignKey:SchoolID;references:ID" json:"school,omitempty"`
	Name         string         `gorm:"column:med_name" json:"mediaName"`
	FileSize     int64          `gorm:"column:med_file_size" json:"fileSize"`
	MimeType     string         `gorm:"column:med_mime_type" json:"mimeType"`
	StoragePath  string         `gorm:"column:med_storage_path" json:"storagePath"`
	FileURL      string         `gorm:"column:med_file_url" json:"fileUrl"`
	ThumbnailURL string         `gorm:"column:med_thumbnail_url" json:"thumbnailUrl,omitempty"`
	IsPublic     bool           `gorm:"column:is_public;default:true" json:"isPublic"`
	OwnerType    OwnerType      `gorm:"column:med_owner_type;type:owner_type" json:"ownerType"`
	OwnerID      string         `gorm:"column:med_owner_id;type:uuid" json:"ownerId"`
	CreatedAt    time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (Media) TableName() string {
	return "edv.medias"
}
