package domain

type SubjectClass struct {
	ID           string     `gorm:"primaryKey;column:scl_id;default:gen_random_uuid()" json:"subjectClassId"`
	ClassID      string     `gorm:"column:scl_cls_id;type:uuid" json:"classId"`
	Class        Class      `gorm:"foreignKey:ClassID;references:ID" json:"class,omitempty"`
	SubjectID    string     `gorm:"column:scl_sub_id;type:uuid" json:"subjectId"`
	Subject      Subject    `gorm:"foreignKey:SubjectID;references:ID" json:"subject,omitempty"`
	SchoolUserID string     `gorm:"column:scl_scu_id;type:uuid" json:"teacherId"`
	Teacher      SchoolUser `gorm:"foreignKey:SchoolUserID;references:ID" json:"teacher,omitempty"`
}

func (SubjectClass) TableName() string {
	return "edv.subject_classes"
}
