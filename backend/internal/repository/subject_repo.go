package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type SubjectRepository interface{
	CreateSubject(subject *domain.Subject) error
	GetAllSubjects(schoolID string) ([]*domain.Subject, error) // Tambahkan schoolID
	GetSubjectByID(id string) (*domain.Subject, error)
	GetSubjectByCode(code string, schoolID string) (*domain.Subject, error)
	UpdateSubject(subject *domain.Subject) error
	DeleteSubject(id string) error
}

type subjectRepository struct {
	db *gorm.DB
}

//constructor
func NewSubjectRepository(db *gorm.DB) SubjectRepository {
	return &subjectRepository{db: db}
}

func (r *subjectRepository) CreateSubject(subject *domain.Subject) error {
	return r.db.Create(subject).Error
}

func (r *subjectRepository) GetAllSubjects(schoolID string) ([]*domain.Subject, error) {
	var subjects []*domain.Subject
	// Menambahkan filter schoolID dan Preload School
	err := r.db.Preload("School").Where("sub_sch_id = ?", schoolID).Find(&subjects).Error
	return subjects, err
}

func (r *subjectRepository) GetSubjectByID(id string) (*domain.Subject, error) {
	var subject domain.Subject
	err := r.db.Preload("School").Where("sub_id = ?", id).First(&subject).Error
	return &subject, err
}

func (r *subjectRepository) GetSubjectByCode(code string, schoolID string) (*domain.Subject, error) {
	var subject domain.Subject
	err := r.db.Preload("School").Where("sub_sch_id = ? and sub_code = ? ", schoolID, code).First(&subject).Error
	return &subject, err
}

func (r *subjectRepository) UpdateSubject(subject *domain.Subject) error {
	return r.db.Save(subject).Error
}

func (r *subjectRepository) DeleteSubject(id string) error {
	return r.db.Delete(&domain.Subject{}, "sub_id = ?", id).Error
}