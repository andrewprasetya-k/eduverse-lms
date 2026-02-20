package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type AssignmentRepository interface {
	// Category
	CreateCategory(cat *domain.AssignmentCategory) error
	GetCategoriesBySchool(schoolID string) ([]*domain.AssignmentCategory, error)

	// Assignment
	CreateAssignment(asg *domain.Assignment) error
	GetAssignmentsBySubjectClass(subjectClassID string) ([]*domain.Assignment, error)
	GetAssignmentByID(id string) (*domain.Assignment, error)

	// Submission
	CreateSubmission(sbm *domain.Submission) error
	GetSubmissionsByAssignment(asgID string) ([]*domain.Submission, error)
	GetSubmissionByID(id string) (*domain.Submission, error)

	// Assessment
	UpsertAssessment(asm *domain.Assessment) error
	GetAssessmentBySubmission(sbmID string) (*domain.Assessment, error)

	// Weights
	SetWeight(weight *domain.AssessmentWeight) error
	GetWeightsBySubject(subID string) ([]*domain.AssessmentWeight, error)
}

type assignmentRepository struct {
	db *gorm.DB
}

func NewAssignmentRepository(db *gorm.DB) AssignmentRepository {
	return &assignmentRepository{db: db}
}

func (r *assignmentRepository) CreateCategory(cat *domain.AssignmentCategory) error {
	return r.db.Create(cat).Error
}

func (r *assignmentRepository) GetCategoriesBySchool(schoolID string) ([]*domain.AssignmentCategory, error) {
	var cats []*domain.AssignmentCategory
	err := r.db.Where("asc_sch_id = ?", schoolID).Find(&cats).Error
	return cats, err
}

func (r *assignmentRepository) CreateAssignment(asg *domain.Assignment) error {
	return r.db.Create(asg).Error
}

func (r *assignmentRepository) GetAssignmentsBySubjectClass(subjectClassID string) ([]*domain.Assignment, error) {
	var results []*domain.Assignment
	err := r.db.Preload("Category").Preload("SubjectClass.Subject").
		Where("asg_scl_id = ?", subjectClassID).
		Order("created_at desc").Find(&results).Error
	return results, err
}

func (r *assignmentRepository) GetAssignmentByID(id string) (*domain.Assignment, error) {
	var asg domain.Assignment
	err := r.db.Preload("Category").Preload("SubjectClass.Subject").
		Where("asg_id = ?", id).First(&asg).Error
	return &asg, err
}

func (r *assignmentRepository) CreateSubmission(sbm *domain.Submission) error {
	return r.db.Create(sbm).Error
}

func (r *assignmentRepository) GetSubmissionsByAssignment(asgID string) ([]*domain.Submission, error) {
	var results []*domain.Submission
	err := r.db.Preload("User").Where("sbm_asg_id = ?", asgID).Order("submitted_at asc").Find(&results).Error
	return results, err
}

func (r *assignmentRepository) GetSubmissionByID(id string) (*domain.Submission, error) {
	var sbm domain.Submission
	err := r.db.Preload("User").Where("sbm_id = ?", id).First(&sbm).Error
	return &sbm, err
}

func (r *assignmentRepository) UpsertAssessment(asm *domain.Assessment) error {
	return r.db.Save(asm).Error
}

func (r *assignmentRepository) GetAssessmentBySubmission(sbmID string) (*domain.Assessment, error) {
	var asm domain.Assessment
	err := r.db.Preload("Assessor").Where("asm_sbm_id = ?", sbmID).First(&asm).Error
	return &asm, err
}

func (r *assignmentRepository) SetWeight(weight *domain.AssessmentWeight) error {
	return r.db.Save(weight).Error
}

func (r *assignmentRepository) GetWeightsBySubject(subID string) ([]*domain.AssessmentWeight, error) {
	var weights []*domain.AssessmentWeight
	err := r.db.Preload("Category").Where("asw_sub_id = ?", subID).Find(&weights).Error
	return weights, err
}
