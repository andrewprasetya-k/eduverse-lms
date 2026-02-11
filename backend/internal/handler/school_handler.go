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

// Create
func (h *SchoolHandler) CreateSchool(c *gin.Context) {
    var input dto.CreateSchoolDTO
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // 1. Ubah DTO menjadi Domain Model
    school := domain.School{
        Name:    input.Name,
        Code:    input.Code,
        LogoID:  input.LogoID,
        Address: input.Address,
        Email:   input.Email,
        Phone:   input.Phone,
        Website: input.Website,
    }

    // 2. Kirim Domain Model ke Service
    if err := h.service.CreateSchool(&school); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, school)
}

// Get All
func (h *SchoolHandler) GetAllSchools(c *gin.Context) {
	schools, err := h.service.GetAllSchools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Buat slice baru untuk menampung kodenya saja
    var response []dto.SchoolResponseDTO
    for _, s := range schools {
        response = append(response, dto.SchoolResponseDTO{
			ID: s.ID,
			Name: s.Name,
            Code: s.Code,
			IsDeleted: s.DeletedAt.Valid,
			CreatedAt: s.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: s.UpdatedAt.Format("02-01-2006 15:04:05"),
        })
    }

	c.JSON(http.StatusOK, response)
}

// Get Active
func (h *SchoolHandler) GetActiveSchools(c *gin.Context) {
	schools, err := h.service.GetActiveSchools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Buat slice baru untuk menampung kodenya saja
    var response []dto.SchoolResponseDTO
    for _, s := range schools {
        response = append(response, dto.SchoolResponseDTO{
			ID: s.ID,
			Name: s.Name,
            Code: s.Code,
			IsDeleted: s.DeletedAt.Valid,
			CreatedAt: s.CreatedAt.Format("02-01-2006 15:04:05"),
			UpdatedAt: s.UpdatedAt.Format("02-01-2006 15:04:05"),
        })
    }
	c.JSON(http.StatusOK, response)
}

// Get Deleted
func (h *SchoolHandler) GetDeletedSchools(c *gin.Context) {
	schools, err := h.service.GetDeletedSchools()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
		// Buat slice baru untuk menampung kodenya saja
    var response []dto.SchoolResponseDTO
    for _, s := range schools {
        response = append(response, dto.SchoolResponseDTO{
			ID: s.ID,
			Name: s.Name,
            Code: s.Code,
			IsDeleted: s.DeletedAt.Valid,
			CreatedAt: s.CreatedAt.String(),
			UpdatedAt: s.UpdatedAt.Local().String(),
        })
    }
	c.JSON(http.StatusOK, response)
}

// Get By Code
func (h *SchoolHandler) GetSchoolByCode(c *gin.Context) {
	schoolCode := c.Param("schoolCode")
	school, err := h.service.GetSchoolByCode(schoolCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "School not found"})
		return
	}
	c.JSON(http.StatusOK, school)
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

	c.JSON(http.StatusOK, school)
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
