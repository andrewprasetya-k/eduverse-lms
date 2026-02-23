package service

import (
	"backend/internal/domain"
	"backend/internal/repository"
	"fmt"
	"time"
)

type AssignmentService interface {
	// Category
	CreateCategory(cat *domain.AssignmentCategory) error
	GetCategoriesBySchool(schoolID string) ([]*domain.AssignmentCategory, error)

	// Assignment
	CreateAssignment(asg *domain.Assignment, mediaIDs []string) error
	GetAssignmentsBySubjectClass(subjectClassID string) ([]*domain.Assignment, error)
	GetAssignmentByID(id string) (*domain.Assignment, error)
	GetAssignmentWithSubmissions(id string) (*domain.Assignment, error)

	// Submission
	Submit(sbm *domain.Submission, mediaIDs []string) error
	GetSubmissions(asgID string) ([]*domain.Submission, error)
	GetSubmissionByID(id string) (*domain.Submission, error)
	UpdateSubmission(id string, mediaIDs []string) error
	DeleteSubmission(id string) error

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

func (s *assignmentService) GetAssignmentWithSubmissions(id string) (*domain.Assignment, error) {
	asg, err := s.repo.GetAssignmentWithSubmissions(id)
	if err != nil {
		return nil, err
	}

	// Load attachments for each submission
	for i := range asg.Submissions {
		atts, _ := s.attService.GetBySource(string(domain.SourceSubmission), asg.Submissions[i].ID)
		for _, a := range atts {
			asg.Submissions[i].Attachments = append(asg.Submissions[i].Attachments, *a)
		}
	}

	return asg, nil
}

func (s *assignmentService) Submit(sbm *domain.Submission, mediaIDs []string) error {
	sbm.SubmittedAt = time.Now()
	
	// Check deadline before submitting
	assignment, err := s.repo.GetAssignmentByID(sbm.AssignmentID)
	if err != nil {
		return err
	}
	
	if !assignment.AllowLateSubmission && assignment.Deadline != nil && assignment.Deadline.Before(sbm.SubmittedAt) {
		return fmt.Errorf("submission past due")
	}

	err = s.repo.UpsertSubmission(sbm)
	if err != nil {
		return err
	}

	// Unlink existing attachments for this submission if updating
	s.attService.UnlinkBySource(string(domain.SourceSubmission), sbm.ID)

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

func (s *assignmentService) UpdateSubmission(id string, mediaIDs []string) error {
	sbm, err := s.repo.GetSubmissionByID(id)
	if err != nil {
		return err
	}

	sbm.SubmittedAt = time.Now()
	err = s.repo.UpdateSubmission(sbm)
	if err != nil {
		return err
	}

	s.attService.UnlinkBySource(string(domain.SourceSubmission), id)
	for _, mID := range mediaIDs {
		att := &domain.Attachment{
			SchoolID:   sbm.SchoolID,
			SourceID:   id,
			SourceType: domain.SourceSubmission,
			MediaID:    mID,
		}
		s.attService.Link(att)
	}
	return nil
}

func (s *assignmentService) DeleteSubmission(id string) error {
	s.attService.UnlinkBySource(string(domain.SourceSubmission), id)
	return s.repo.DeleteSubmission(id)
}
