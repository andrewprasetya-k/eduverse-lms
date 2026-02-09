package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	service service.SubjectService
}

func NewSubjectHandler(service service.SubjectService) *SubjectHandler {
	return &SubjectHandler{service: service}
}

//Post
func (h *SubjectHandler) CreateSubject(c *gin.Context) {
	var input dto.CreateSubjectDTO
	//parse json dari request body ke struct school
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//mapping dto ke domain(model)
	subject := domain.Subject{
		SchoolID: input.SchoolID,
		Name:    input.Name,
		Code:    input.Code,
	}
	//panggil service untuk create subject
	if err := h.service.CreateSubject(&subject); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//return response
	c.JSON(http.StatusCreated, subject)
}

//Get All
func (h *SubjectHandler) GetAllSubjects(c *gin.Context) {
	subjects, err := h.service.GetAllSubjects()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}

//Get By ID
func (h *SubjectHandler) GetSubjectByID(c *gin.Context) {
	id := c.Param("id")
	subject, err := h.service.GetSubjectByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subject)
}

//Get By code
func (h *SubjectHandler) GetSubjectByCode(c *gin.Context) {
	code := c.Param("code")
	subject, err := h.service.GetSubjectByCode(code)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subject)
}

//Put
func (h *SubjectHandler) UpdateSubject(c *gin.Context) {
	var input dto.UpdateSubjectDTO
	//parse json dari request body ke struct school
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	code := c.Param("code")
	subject, err := h.service.GetSubjectByCode(code)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//update field yang diubah
	if input.Name != nil {
		subject.Name = *input.Name
	}
	if input.Code != nil {
		subject.Code = *input.Code
	}
	//panggil service untuk update subject
	if err := h.service.UpdateSubject(subject); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//return response
	c.JSON(http.StatusOK, subject)
}

//Delete
func (h *SubjectHandler) DeleteSubject(c *gin.Context) {
	code := c.Param("code")
	if err := h.service.DeleteSubject(code); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subject deleted successfully"})
}