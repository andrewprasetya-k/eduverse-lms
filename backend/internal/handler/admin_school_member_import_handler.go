package handler

import (
	"backend/internal/dto"
	"backend/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminSchoolMemberImportHandler struct {
	service service.AdminSchoolMemberImportService
}

func NewAdminSchoolMemberImportHandler(service service.AdminSchoolMemberImportService) *AdminSchoolMemberImportHandler {
	return &AdminSchoolMemberImportHandler{service: service}
}

func (h *AdminSchoolMemberImportHandler) Preview(c *gin.Context) {
	schoolID, ok := getImportActiveSchoolID(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Konteks sekolah aktif wajib tersedia."})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File import wajib diunggah."})
		return
	}

	openedFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File import belum bisa dibuka."})
		return
	}
	defer openedFile.Close()

	response, err := h.service.PreviewCSV(schoolID, openedFile)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *AdminSchoolMemberImportHandler) Commit(c *gin.Context) {
	schoolID, ok := getImportActiveSchoolID(c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Konteks sekolah aktif wajib tersedia."})
		return
	}

	var input dto.AdminSchoolMemberImportCommitRequestDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		HandleBindingError(c, err)
		return
	}

	response, err := h.service.Commit(schoolID, input.DefaultPassword, input.Rows)
	if err != nil {
		if response != nil {
			c.JSON(http.StatusBadRequest, response)
			return
		}
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func getImportActiveSchoolID(c *gin.Context) (string, bool) {
	if value, exists := c.Get("school_id"); exists {
		if schoolID, ok := value.(string); ok && schoolID != "" {
			return schoolID, true
		}
	}
	schoolID := c.GetHeader("SchoolId")
	return schoolID, schoolID != ""
}
