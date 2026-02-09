package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SchoolHandler struct {
	service service.SchoolService
}

func NewSchoolHandler(service service.SchoolService) *SchoolHandler {
	return &SchoolHandler{service: service}
}

//Post
func (h *SchoolHandler) CreateSchool(c *gin.Context) {
	var input dto.CreateSchoolDTO
	//parse json dari request body ke struct school
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	school := domain.School{
		Name:    input.Name,
		Code:    input.Code,
		LogoID:  input.LogoID,
		Address: input.Address,
		Email:   input.Email,
		Phone:   input.Phone,
		Website: input.Website,
	}
	//panggil service untuk create school
	if err := h.service.CreateSchool(&school); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//return response
	c.JSON(http.StatusCreated, school)
}

//Get All
func (h *SchoolHandler) GetAllSchools(c *gin.Context) {
	schools, err := h.service.GetAllSchools()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schools)
}

//Get By ID
func (h *SchoolHandler) GetSchoolByID(c *gin.Context) {
	id := c.Param("id")
	school, err := h.service.GetSchoolByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, school)
}

//Put
func (h *SchoolHandler) UpdateSchool(c *gin.Context) {
	var input dto.UpdateSchoolDTO
	//parse json dari request body ke struct school
	if err:= c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id := c.Param("id")
	school, err := h.service.GetSchoolByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//update field yang diubah
	if input.Name != nil {
		school.Name = *input.Name
	}
	if input.Code != nil {
		school.Code = *input.Code
	}
	if input.LogoID != nil {
		school.LogoID = input.LogoID
	}
	if input.Address != nil {
		school.Address = *input.Address
	}
	if input.Email != nil {
		school.Email = *input.Email
	}
	if input.Phone != nil {
		school.Phone = *input.Phone
	}
	if input.Website != nil {
		school.Website = input.Website
	}
	//panggil service untuk update school
	if err := h.service.UpdateSchool(school); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	//return response
	c.JSON(http.StatusOK, school)
}

//Delete
func (h *SchoolHandler) DeleteSchool(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.DeleteSchool(id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "School deleted successfully"})
}