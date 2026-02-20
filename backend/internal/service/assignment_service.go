package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
)

type AssignmentService interface {
	// Category
	CreateCategory(cat *domain.AssignmentCategory) error
	GetCategoriesBySchool(schoolID string) ([]*domain.AssignmentCategory, error)

	// Assignment
	CreateAssignment(asg *domain.Assignment, mediaIDs []string) error
	GetAssignmentsBySubjectClass(subjectClassID string) ([]*domain.Assignment, error)
	GetAssignmentByID(id string) (*domain.Assignment, error)

	// Submission
	Submit(sbm *domain.Submission, mediaIDs []string) error
	GetSubmissions(asgID string) ([]*domain.Submission, error)
	GetSubmissionByID(id string) (*domain.Submission, error)

	// Assessment
	Assess(asm *domain.Assessment) error
}

type assignmentService struct {
	repo       repository.AssignmentRepository
	attService AttachmentService
}

func NewAssignmentService(repo repository.AssignmentRepository, attService AttachmentService) AssignmentService {
	return &assignmentService{
		repo:       repo,
		attService: attService,
	}
}

func (s *assignmentService) CreateCategory(cat *domain.AssignmentCategory) error {
	return s.repo.CreateCategory(cat)
}

func (s *assignmentService) GetCategoriesBySchool(schoolID string) ([]*domain.AssignmentCategory, error) {
	return s.repo.GetCategoriesBySchool(schoolID)
}

func (s *assignmentService) CreateAssignment(asg *domain.Assignment, mediaIDs []string) error {
	err := s.repo.CreateAssignment(asg)
	if err != nil {
		return err
	}

	for _, mID := range mediaIDs {
		att := &domain.Attachment{
			SchoolID:   asg.SchoolID,
			SourceID:   asg.ID,
			SourceType: domain.SourceAssignment,
			MediaID:    mID,
		}
		s.attService.Link(att)
	}
	return nil
}

func (s *assignmentService) GetAssignmentsBySubjectClass(subjectClassID string) ([]*domain.Assignment, error) {
	results, err := s.repo.GetAssignmentsBySubjectClass(subjectClassID)
	if err != nil {
		return nil, err
	}

	for _, asg := range results {
		atts, _ := s.attService.GetBySource(string(domain.SourceAssignment), asg.ID)
		for _, a := range atts {
			asg.Attachments = append(asg.Attachments, *a)
		}
	}
	return results, nil
}

func (s *assignmentService) GetAssignmentByID(id string) (*domain.Assignment, error) {
	asg, err := s.repo.GetAssignmentByID(id)
	if err != nil {
		return nil, err
	}

	atts, _ := s.attService.GetBySource(string(domain.SourceAssignment), id)
	for _, a := range atts {
		asg.Attachments = append(asg.Attachments, *a)
	}
	return asg, nil
}

func (s *assignmentService) Submit(sbm *domain.Submission, mediaIDs []string) error {
	err := s.repo.CreateSubmission(sbm)
	if err != nil {
		return err
	}

	for _, mID := range mediaIDs {
		att := &domain.Attachment{
			SchoolID:   sbm.SchoolID,
			SourceID:   sbm.ID,
			SourceType: domain.SourceSubmission,
			MediaID:    mID,
		}
		s.attService.Link(att)
	}
	return nil
}

func (s *assignmentService) GetSubmissions(asgID string) ([]*domain.Submission, error) {
	results, err := s.repo.GetSubmissionsByAssignment(asgID)
	if err != nil {
		return nil, err
	}

	for _, sbm := range results {
		atts, _ := s.attService.GetBySource(string(domain.SourceSubmission), sbm.ID)
		for _, a := range atts {
			sbm.Attachments = append(sbm.Attachments, *a)
		}
	}
	return results, nil
}

func (s *assignmentService) GetSubmissionByID(id string) (*domain.Submission, error) {
	sbm, err := s.repo.GetSubmissionByID(id)
	if err != nil {
		return nil, err
	}

	atts, _ := s.attService.GetBySource(string(domain.SourceSubmission), id)
	for _, a := range atts {
		sbm.Attachments = append(sbm.Attachments, *a)
	}
	return sbm, nil
}

func (s *assignmentService) Assess(asm *domain.Assessment) error {
	return s.repo.UpsertAssessment(asm)
}
