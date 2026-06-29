package domain

import "time"

type ChatAttachment struct {
	ID        string    `gorm:"primaryKey;column:cat_id;default:gen_random_uuid()" json:"attachmentId"`
	MessageID string    `gorm:"column:cat_msg_id;type:uuid" json:"messageId"`
	MediaID   string    `gorm:"column:cat_med_id;type:uuid" json:"mediaId"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
}

func (ChatAttachment) TableName() string {
	return "edv.chat_attachments"
}
