package domain

type Permission struct {
	ID          string `gorm:"primaryKey;column:prm_id;default:gen_random_uuid()" json:"permissionId"`
	Key         string `gorm:"column:prm_key;unique" json:"permissionKey"`
	Description string `gorm:"column:prm_desc" json:"description"`
}

func (Permission) TableName() string {
	return "edv.permissions"
}
