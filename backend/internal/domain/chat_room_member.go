package domain

import "time"

type ChatRoomMember struct {
	ID       string     `gorm:"primaryKey;column:crm_id;default:gen_random_uuid()" json:"chatRoomMemberId"`
	RoomID   string     `gorm:"column:crm_room_id;type:uuid" json:"roomId"`
	UserID   string     `gorm:"column:crm_usr_id;type:uuid" json:"userId"`
	EnrollID *string    `gorm:"column:crm_enr_id;type:uuid" json:"enrollmentId,omitempty"`
	Role     string     `gorm:"column:crm_role" json:"role"`
	JoinedAt time.Time  `gorm:"column:joined_at;autoCreateTime" json:"joinedAt"`
	LeftAt   *time.Time `gorm:"column:left_at" json:"leftAt,omitempty"`
}

func (ChatRoomMember) TableName() string {
	return "edv.chat_room_members"
}
