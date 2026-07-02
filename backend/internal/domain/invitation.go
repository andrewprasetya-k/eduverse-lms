package domain

import "time"

type Invitation struct {
	ID           string     `gorm:"primaryKey;column:inv_id;default:gen_random_uuid()" json:"invitationId"`
	SchoolID     string     `gorm:"column:inv_school_id;type:uuid" json:"schoolId"`
	School       School     `gorm:"foreignKey:SchoolID;references:ID" json:"school,omitempty"`
	Email        string     `gorm:"column:inv_email" json:"email"`
	Role         string     `gorm:"column:inv_role" json:"role"`
	FullName     *string    `gorm:"column:inv_full_name" json:"fullName,omitempty"`
	ClassID      *string    `gorm:"column:inv_class_id;type:uuid" json:"classId,omitempty"`
	Class        Class      `gorm:"foreignKey:ClassID;references:ID" json:"class,omitempty"`
	TokenHash    string     `gorm:"column:inv_token_hash" json:"-"`
	InvitedBy    string     `gorm:"column:inv_invited_by;type:uuid" json:"invitedBy"`
	TargetUserID *string    `gorm:"column:inv_target_user_id;type:uuid" json:"targetUserId,omitempty"`
	ExpiresAt    time.Time  `gorm:"column:inv_expires_at" json:"expiresAt"`
	AcceptedAt   *time.Time `gorm:"column:inv_accepted_at" json:"acceptedAt,omitempty"`
	RevokedAt    *time.Time `gorm:"column:inv_revoked_at" json:"revokedAt,omitempty"`
	CreatedAt    time.Time  `gorm:"column:created_at;autoCreateTime" json:"createdAt"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;autoUpdateTime" json:"updatedAt"`
}

func (Invitation) TableName() string {
	return "edv.invitations"
}
