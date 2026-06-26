package domain

import (
	"time"

	"gorm.io/gorm"
)

type ChatRoom struct {
	ID        string         `gorm:"primaryKey;column:room_id;default:gen_random_uuid()" json:"roomId"`
	SchoolID  string         `gorm:"column:room_sch_id;type:uuid" json:"schoolId"`
	Name      string         `gorm:"column:room_name" json:"roomName"`
	Type      string         `gorm:"column:room_type" json:"roomType"`
	RefType   string         `gorm:"column:room_ref_type" json:"refType"`
	RefID     string         `gorm:"column:room_ref_id;type:uuid" json:"refId"`
	CreatedBy string         `gorm:"column:created_by;type:uuid" json:"createdBy"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (ChatRoom) TableName() string {
	return "edv.chat_rooms"
}
