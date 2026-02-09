package repository

import (
	"backend/internal/domain"

	"gorm.io/gorm"
)

type SubjectRepository interface{
	CreateSubject(subject *domain.Subject) error
	GetAllSubjects() ([]*domain.Subject, error)
	GetSubjectByID(id string) (*domain.Subject, error)
	GetSubjectByCode(code string) (*domain.Subject, error)
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

func (r *subjectRepository) GetAllSubjects() ([]*domain.Subject, error) {
	var subject []*domain.Subject
	err := r.db.Find(&subject).Error
	return subject, err
}

func (r *subjectRepository) GetSubjectByID(id string) (*domain.Subject, error) {
	var subject domain.Subject
	err := r.db.First(&subject, "sub_id = ?", id).Error
	return &subject, err
}

func (r *subjectRepository) GetSubjectByCode(code string) (*domain.Subject, error) {
	var subject domain.Subject
	err := r.db.First(&subject, "sub_code = ?", code).Error
	return &subject, err
}

func (r *subjectRepository) UpdateSubject(subject *domain.Subject) error {
	return r.db.Save(subject).Error
}

func (r *subjectRepository) DeleteSubject(code string) error {
	return r.db.Delete(&domain.Subject{}, "sub_code = ?", code).Error
}