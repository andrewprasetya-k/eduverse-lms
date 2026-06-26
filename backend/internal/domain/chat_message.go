package domain

import (
	"time"

	"gorm.io/gorm"
)

type ChatMessage struct {
	ID        string         `gorm:"primaryKey;column:msg_id;default:gen_random_uuid()" json:"messageId"`
	RoomID    string         `gorm:"column:msg_room_id;type:uuid" json:"roomId"`
	UserID    string         `gorm:"column:msg_usr_id;type:uuid" json:"senderId"`
	Content   string         `gorm:"column:msg_content" json:"content"`
	Type      string         `gorm:"column:msg_type" json:"messageType"`
	ReplyTo   *string        `gorm:"column:msg_reply_to;type:uuid" json:"replyTo,omitempty"`
	MediaID   *string        `gorm:"column:msg_med_id;type:uuid" json:"mediaId,omitempty"`
	RefType   *string        `gorm:"column:msg_ref_type" json:"refType,omitempty"`
	RefID     *string        `gorm:"column:msg_ref_id;type:uuid" json:"refId,omitempty"`
	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}

func (ChatMessage) TableName() string {
	return "edv.chat_messages"
}
