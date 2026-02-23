package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AssignmentHandler struct {
	service             service.AssignmentService
	schoolService       service.SchoolService
	subjectClassService service.SubjectClassService
}

func NewAssignmentHandler(service service.AssignmentService, schoolService service.SchoolService, subjectClassService service.SubjectClassService) *AssignmentHandler {
	return &AssignmentHandler{
		service:             service,
		schoolService:       schoolService,
		subjectClassService: subjectClassService,
	}
}

func (h *AssignmentHandler) CreateCategory(c *gin.Context) {
	var input dto.CreateAssignmentCategoryDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	cat := domain.AssignmentCategory{
		SchoolID: input.SchoolID,
		Name:     input.Name,
	}

	if err := h.service.CreateCategory(&cat); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created"})
}

func (h *AssignmentHandler) GetCategoriesBySchool(c *gin.Context) {
	schoolCode := c.Param("schoolCode")

	// 1. Get School Header
	school, err := h.schoolService.GetSchoolByCode(schoolCode)
	if err != nil {
		HandleError(c, err)
		return
	}

	// 2. Get Categories
	cats, err := h.service.GetCategoriesBySchool(school.ID)
	if err != nil {
		HandleError(c, err)
		return
	}

	var response []dto.AssignmentCategoryResponseDTO
	for _, cat := range cats {
		response = append(response, dto.AssignmentCategoryResponseDTO{
			ID:        cat.ID,
			SchoolID:  cat.SchoolID,
			Name:      cat.Name,
			CreatedAt: cat.CreatedAt.Format("02-01-2006 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, dto.SchoolWithAssignmentCategoriesDTO{
		School: dto.SchoolHeaderDTO{
			ID:     school.ID,
			Name:   school.Name,
			Code:   school.Code,
			LogoID: school.LogoID,
		},
		Categories: response,
	})
}

func (h *AssignmentHandler) CreateAssignment(c *gin.Context) {
	var input dto.CreateAssignmentDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	asg := domain.Assignment{
		SchoolID:       input.SchoolID,
		SubjectClassID: input.SubjectClassID,
		CategoryID:     input.CategoryID,
		Title:          input.Title,
		Description:    input.Description,
		Deadline:       input.Deadline,
		CreatedBy:      input.CreatedBy,
	}

	if err := h.service.CreateAssignment(&asg, input.MediaIDs); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Assignment created"})
}

func (h *AssignmentHandler) GetBySubjectClass(c *gin.Context) {
	subjectClassID := c.Param("subjectClassId")

	// 1. Get SubjectClass Header
	subjectClassHeader, err := h.subjectClassService.GetByID(subjectClassID)
	if err != nil {
		HandleError(c, err)
		return
	}

	// 2. Get Assignments
	results, err := h.service.GetAssignmentsBySubjectClass(subjectClassID)
	if err != nil {
		HandleError(c, err)
		return
	}

	var assignments []dto.AssignmentResponseDTO
	for _, r := range results {
		assignments = append(assignments, h.mapAsgToResponse(r))
	}

	response := dto.AssignmentPerSubjectClassResponseDTO{
		SubjectClass: dto.SubjectClassHeaderDTO{
			ID:          subjectClassHeader.ID,
			SubjectCode: subjectClassHeader.Subject.Code,
			SubjectName: subjectClassHeader.Subject.Name,
			TeacherID:   subjectClassHeader.Teacher.ID,
			TeacherName: subjectClassHeader.Teacher.User.FullName,
		},
		Assignments: assignments,
	}

	c.JSON(http.StatusOK, response)
}

func (h *AssignmentHandler) GetSubmissionsByAssignment(c *gin.Context) {
	submissionId := c.Param("submissionId")
	asg, err := h.service.GetAssignmentWithSubmissions(submissionId)
	if err != nil {
		HandleError(c, err)
		return
	}

	var submissionsDTO []dto.SubmissionResponseDTO
	for _, s := range asg.Submissions {
		var assessmentDTO *dto.AssessmentResponseDTO
		if s.Assessment != nil {
			assessmentDTO = &dto.AssessmentResponseDTO{
				Score:      s.Assessment.Score,
				Feedback:   s.Assessment.Feedback,
				Assessor:   s.Assessment.Assessor.FullName,
				AssessedAt: s.Assessment.AssessedAt.Format("02-01-2006 15:04:05"),
			}
		}

		var atts []dto.MediaResponseDTO
		for _, a := range s.Attachments {
			atts = append(atts, dto.MediaResponseDTO{
				ID:       a.Media.ID,
				Name:     a.Media.Name,
				FileURL:  a.Media.FileURL,
				MimeType: a.Media.MimeType,
			})
		}

		submissionsDTO = append(submissionsDTO, dto.SubmissionResponseDTO{
			ID:          s.ID,
			UserName:    s.User.FullName,
			SubmittedAt: s.SubmittedAt.Format("02-01-2006 15:04:05"),
			Attachments: atts,
			Assessment:  assessmentDTO,
		})
	}

	response := dto.AssignmentWithSubmissionsDTO{
		Assignment: dto.AssignmentHeaderDTO{
			ID:           asg.ID,
			Title:        asg.Title,
			SubjectName:  asg.SubjectClass.Subject.Name,
			CategoryName: asg.Category.Name,
			Deadline:     asg.Deadline,
		},
		Submissions: submissionsDTO,
	}

	c.JSON(http.StatusOK, response)
}

func (h *AssignmentHandler) Submit(c *gin.Context) {
	var input dto.CreateSubmissionDTO
	var assignmentId = c.Param("assignmentId")
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	sbm := domain.Submission{
		SchoolID:     input.SchoolID,
		AssignmentID: assignmentId,
		UserID:       input.UserID,
	}

	if err := h.service.Submit(&sbm, input.MediaIDs); err != nil {
		if err.Error() == "Submission past due"{
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Submission received"})
}

func (h *AssignmentHandler) UpdateSubmission(c *gin.Context) {
	var input dto.CreateSubmissionDTO
	submissionId := c.Param("submissionId")
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	if err := h.service.UpdateSubmission(submissionId, input.MediaIDs); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission updated"})
}

func (h *AssignmentHandler) DeleteSubmission(c *gin.Context) {
	submissionId := c.Param("submissionId")

	if err := h.service.DeleteSubmission(submissionId); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Submission deleted"})
}

func (h *AssignmentHandler) Assess(c *gin.Context) {
	var input dto.CreateAssessmentDTO
	var submissionId = c.Param("submissionId")
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	asm := domain.Assessment{
		SubmissionID: submissionId,
		Score:        input.Score,
		Feedback:     input.Feedback,
		AssessedBy:   input.AssessedBy,
	}

	if err := h.service.Assess(&asm); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Assessment recorded"})
}

func (h *AssignmentHandler) mapAsgToResponse(a *domain.Assignment) dto.AssignmentResponseDTO {
	var atts []dto.MediaResponseDTO
	for _, att := range a.Attachments {
		atts = append(atts, dto.MediaResponseDTO{
			ID:       att.Media.ID,
			Name:     att.Media.Name,
			FileURL:  att.Media.FileURL,
			MimeType: att.Media.MimeType,
		})
	}

	return dto.AssignmentResponseDTO{
		ID:                  a.ID,
		Title:               a.Title,
		Description:         a.Description,
		CategoryName:        a.Category.Name,
		Deadline:            a.Deadline,
		AllowLateSubmission: a.AllowLateSubmission,
		CreatedAt:           a.CreatedAt.Format("02-01-2006 15:04:05"),
		Attachments:         atts,
	}
}
