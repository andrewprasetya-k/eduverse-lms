package domain

import "time"

type ChatReadReceipt struct {
	ID                string    `gorm:"primaryKey;column:rct_id;default:gen_random_uuid()" json:"readReceiptId"`
	RoomID            string    `gorm:"column:rct_room_id;type:uuid" json:"roomId"`
	UserID            string    `gorm:"column:rct_usr_id;type:uuid" json:"userId"`
	LastReadMessageID *string   `gorm:"column:last_read_msg_id;type:uuid" json:"lastReadMessageId,omitempty"`
	LastReadAt        time.Time `gorm:"column:last_read_at;autoUpdateTime" json:"lastReadAt"`
}

func (ChatReadReceipt) TableName() string {
	return "edv.chat_read_receipts"
}
