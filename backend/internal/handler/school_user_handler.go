package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SchoolUserHandler struct {
	service       service.SchoolUserService
	schoolService service.SchoolService
}

func NewSchoolUserHandler(service service.SchoolUserService, schoolService service.SchoolService) *SchoolUserHandler {
	return &SchoolUserHandler{
		service:       service,
		schoolService: schoolService,
	}
}

func (h *SchoolUserHandler) Enroll(c *gin.Context) {
	var input dto.AddSchoolUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	scu := domain.SchoolUser{
		UserID:   input.UserID,
		SchoolID: input.SchoolID,
	}

	if err := h.service.Enroll(&scu); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User enrolled to school successfully"})
}

func (h *SchoolUserHandler) GetMembersBySchool(c *gin.Context) {
	schoolID := c.Param("schoolId")

	// 1. Ambil data sekolah
	school, err := h.schoolService.GetSchoolByID(schoolID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

	// 2. Ambil daftar anggota
	members, err := h.service.GetMembersBySchool(schoolID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var membersDTO []dto.SchoolUserResponseDTO
	for _, m := range members {
		membersDTO = append(membersDTO, dto.SchoolUserResponseDTO{
			ID:        m.ID,
			UserID:    m.UserID,
			FullName:  m.User.FullName,
			Email:     m.User.Email,
			SchoolID:  m.SchoolID,
			CreatedAt: m.CreatedAt.Format("02-01-2006 15:04:05"),
		})
	}

	response := dto.SchoolWithMembersDTO{
		School:  h.mapSchoolToHeader(school),
		Members: membersDTO,
	}

	c.JSON(http.StatusOK, response)
}

func (h *SchoolUserHandler) mapSchoolToHeader(s *domain.School) dto.SchoolHeaderDTO {
	return dto.SchoolHeaderDTO{
		ID:     s.ID,
		Name:   s.Name,
		Code:   s.Code,
		LogoID: s.LogoID,
	}
}

func (h *SchoolUserHandler) GetSchoolsByUser(c *gin.Context) {
	userID := c.Param("userId")
	schools, err := h.service.GetSchoolsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.SchoolUserResponseDTO
	for _, s := range schools {
		response = append(response, dto.SchoolUserResponseDTO{
			ID:         s.ID,
			UserID:     s.UserID,
			SchoolID:   s.SchoolID,
			SchoolName: s.School.Name,
			SchoolCode: s.School.Code,
			CreatedAt:  s.CreatedAt.Format("02-01-2006 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *SchoolUserHandler) Unenroll(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Unenroll(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User unenrolled from school successfully"})
}
