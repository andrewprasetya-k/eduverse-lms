package handler

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SchoolHandler struct {
	service service.SchoolService
}

func NewSchoolHandler(service service.SchoolService) *SchoolHandler {
	return &SchoolHandler{service: service}
}

// Create
func (h *SchoolHandler) CreateSchool(c *gin.Context) {
    var input dto.CreateSchoolDTO
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

    if err := h.service.CreateSchool(&school); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, h.mapToResponse(&school))
}

// Get Schools (with filter)
func (h *SchoolHandler) GetSchools(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page","1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit","10"))
	status := c.Query("status")
	search := c.Query("search")

	schools, err := h.service.GetSchools(search, status, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    var response []dto.SchoolResponseDTO
    for _, s := range schools {
        response = append(response, h.mapToResponse(s))
    }

	// TODO: Get real total count from database
	var totalItems int64 = int64(len(schools))
	totalPages := (totalItems + int64(limit) - 1) / int64(limit)
	
	paginatedResponse := dto.PaginatedResponse{
		Data: response,
		TotalItems: totalItems,
		Page: page,
		Limit: limit,
		TotalPages: int(totalPages),
	}
	c.JSON(http.StatusOK, paginatedResponse)
}

// Get By Code
func (h *SchoolHandler) GetSchoolByCode(c *gin.Context) {
	schoolCode := c.Param("schoolCode")
	school, err := h.service.GetSchoolByCode(schoolCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}
	c.JSON(http.StatusOK, h.mapToResponse(school))
}

// Update
func (h *SchoolHandler) UpdateSchool(c *gin.Context) {
	var input dto.UpdateSchoolDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	schoolCode := c.Param("schoolCode")
	school, err := h.service.GetSchoolByCode(schoolCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}

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

	if err := h.service.UpdateSchool(school); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.mapToResponse(school))
}


// Helper to map domain to DTO
func (h *SchoolHandler) mapToResponse(s *domain.School) dto.SchoolResponseDTO {
	return dto.SchoolResponseDTO{
		ID:        s.ID,
		Name:      s.Name,
		Code:      s.Code,
		LogoID:    s.LogoID,
		Address:   s.Address,
		Email:     s.Email,
		Phone:     s.Phone,
		Website:   s.Website,
		IsDeleted: s.DeletedAt.Valid,
		CreatedAt: s.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: s.UpdatedAt.Format("02-01-2006 15:04:05"),
	}
}

//restore deleted school
func (h *SchoolHandler) RestoreDeletedSchool(c *gin.Context) {
	schoolCode := c.Param("schoolCode")
	if err := h.service.RestoreDeletedSchool(schoolCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "School restored successfully"})
}

// Delete
func (h *SchoolHandler) DeleteSchool(c *gin.Context) {
	schoolCode := c.Param("schoolCode")
	if err := h.service.DeleteSchool(schoolCode); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "School deleted successfully"})
}
