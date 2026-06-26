package service

import (
	"backend/internal/domain"
	"backend/internal/dto"
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"net/mail"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AdminSchoolMemberImportService interface {
	PreviewCSV(schoolID string, reader io.Reader) (*dto.AdminSchoolMemberImportPreviewResponseDTO, error)
	Commit(schoolID string, defaultPassword string, rows []dto.AdminSchoolMemberImportRowDTO) (*dto.AdminSchoolMemberImportCommitResponseDTO, error)
}

type adminSchoolMemberImportService struct {
	db *gorm.DB
}

func NewAdminSchoolMemberImportService(db *gorm.DB) AdminSchoolMemberImportService {
	return &adminSchoolMemberImportService{db: db}
}

type normalizedImportRow struct {
	RowNumber int
	FullName  string
	Email     string
	Role      string
	ClassCode string
	Errors    []string
}

var allowedSchoolMemberImportRoles = map[string]bool{
	"student": true,
	"teacher": true,
	"admin":   true,
}

func (s *adminSchoolMemberImportService) PreviewCSV(schoolID string, reader io.Reader) (*dto.AdminSchoolMemberImportPreviewResponseDTO, error) {
	content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	rows, err := s.parseCSV(content)
	if err != nil {
		return nil, err
	}

	return s.validateRows(schoolID, rows)
}

func (s *adminSchoolMemberImportService) Commit(schoolID string, defaultPassword string, rows []dto.AdminSchoolMemberImportRowDTO) (*dto.AdminSchoolMemberImportCommitResponseDTO, error) {
	if strings.TrimSpace(defaultPassword) == "" {
		return nil, errors.New("default password wajib diisi")
	}
	if len(rows) == 0 {
		return nil, errors.New("baris import wajib diisi")
	}

	normalizedRows := make([]normalizedImportRow, 0, len(rows))
	for _, row := range rows {
		normalizedRows = append(normalizedRows, normalizedImportRow{
			RowNumber: row.RowNumber,
			FullName:  strings.TrimSpace(row.FullName),
			Email:     strings.ToLower(strings.TrimSpace(row.Email)),
			Role:      strings.ToLower(strings.TrimSpace(row.Role)),
			ClassCode: strings.TrimSpace(row.ClassCode),
		})
	}

	preview, err := s.validateRows(schoolID, normalizedRows)
	if err != nil {
		return nil, err
	}
	if preview.InvalidCount > 0 {
		results := make([]dto.AdminSchoolMemberImportResultDTO, 0, len(preview.Rows))
		for _, row := range preview.Rows {
			reason := strings.Join(row.Errors, "; ")
			if reason == "" {
				reason = "Data belum valid."
			}
			results = append(results, dto.AdminSchoolMemberImportResultDTO{
				RowNumber: row.RowNumber,
				FullName:  row.FullName,
				Email:     row.Email,
				Role:      row.Role,
				ClassCode: row.ClassCode,
				Status:    "failed",
				Reason:    reason,
			})
		}
		return &dto.AdminSchoolMemberImportCommitResponseDTO{
			FailedCount: len(results),
			Results:     results,
		}, errors.New("data import masih memiliki baris yang tidak valid")
	}

	results := make([]dto.AdminSchoolMemberImportResultDTO, 0, len(preview.Rows))
	err = s.db.Transaction(func(tx *gorm.DB) error {
		for _, row := range preview.Rows {
			result := dto.AdminSchoolMemberImportResultDTO{
				RowNumber: row.RowNumber,
				FullName:  row.FullName,
				Email:     row.Email,
				Role:      row.Role,
				ClassCode: row.ClassCode,
				Status:    "imported",
			}

			user, createdUser, err := s.findOrCreateUser(tx, row.FullName, row.Email, defaultPassword)
			if err != nil {
				return fmt.Errorf("baris %d: %w", row.RowNumber, err)
			}
			schoolUser, createdMembership, err := s.findOrCreateSchoolUser(tx, schoolID, user.ID)
			if err != nil {
				return fmt.Errorf("baris %d: %w", row.RowNumber, err)
			}
			roleID, err := s.findRoleID(tx, row.Role)
			if err != nil {
				return fmt.Errorf("baris %d: %w", row.RowNumber, err)
			}
			roleAssigned, err := s.ensureRole(tx, schoolUser.ID, roleID)
			if err != nil {
				return fmt.Errorf("baris %d: %w", row.RowNumber, err)
			}
			classTouched := false
			if row.ClassCode != "" && row.Role == "student" {
				classID, err := s.findClassIDByCode(tx, schoolID, row.ClassCode)
				if err != nil {
					return fmt.Errorf("baris %d: %w", row.RowNumber, err)
				}
				classTouched, err = s.ensureActiveStudentEnrollment(tx, schoolID, schoolUser.ID, classID)
				if err != nil {
					return fmt.Errorf("baris %d: %w", row.RowNumber, err)
				}
			}

			if !createdUser && !createdMembership && !roleAssigned && !classTouched {
				result.Status = "skipped"
				result.Reason = "Akun sudah menjadi warga sekolah dengan data yang sama."
			}
			results = append(results, result)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	response := dto.AdminSchoolMemberImportCommitResponseDTO{Results: results}
	for _, result := range results {
		switch result.Status {
		case "imported":
			response.ImportedCount++
		case "skipped":
			response.SkippedCount++
		case "failed":
			response.FailedCount++
		}
	}
	return &response, nil
}

func (s *adminSchoolMemberImportService) parseCSV(content []byte) ([]normalizedImportRow, error) {
	reader := csv.NewReader(bytes.NewReader(content))
	reader.TrimLeadingSpace = true
	reader.FieldsPerRecord = -1

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("file CSV tidak bisa dibaca: %w", err)
	}
	if len(records) == 0 {
		return nil, errors.New("file CSV kosong")
	}

	header := make(map[string]int)
	for idx, value := range records[0] {
		header[strings.ToLower(strings.TrimSpace(value))] = idx
	}

	requiredHeaders := []string{"fullname", "email", "role"}
	for _, name := range requiredHeaders {
		if _, ok := header[name]; !ok {
			return nil, fmt.Errorf("kolom %s wajib ada", name)
		}
	}

	rows := make([]normalizedImportRow, 0, len(records)-1)
	classCodeIndex, hasClassCode := header["classcode"]
	if !hasClassCode {
		classCodeIndex = -1
	}

	for index, record := range records[1:] {
		rowNumber := index + 2
		if isEmptyCSVRecord(record) {
			continue
		}
		rows = append(rows, normalizedImportRow{
			RowNumber: rowNumber,
			FullName:  csvValue(record, header["fullname"]),
			Email:     strings.ToLower(csvValue(record, header["email"])),
			Role:      strings.ToLower(csvValue(record, header["role"])),
			ClassCode: csvValue(record, classCodeIndex),
		})
	}

	if len(rows) == 0 {
		return nil, errors.New("file CSV tidak memiliki baris data")
	}
	return rows, nil
}

func (s *adminSchoolMemberImportService) validateRows(schoolID string, rows []normalizedImportRow) (*dto.AdminSchoolMemberImportPreviewResponseDTO, error) {
	duplicateEmails := map[string]int{}
	classCodes := map[string]bool{}

	for _, row := range rows {
		if row.Email != "" {
			duplicateEmails[row.Email]++
		}
		if row.ClassCode != "" {
			classCodes[row.ClassCode] = true
		}
	}

	existingClassCodes, err := s.existingClassCodes(schoolID, classCodes)
	if err != nil {
		return nil, err
	}

	response := &dto.AdminSchoolMemberImportPreviewResponseDTO{
		Rows: make([]dto.AdminSchoolMemberImportRowDTO, 0, len(rows)),
	}

	for _, row := range rows {
		errorsForRow := append([]string{}, row.Errors...)
		if row.FullName == "" {
			errorsForRow = append(errorsForRow, "Nama lengkap wajib diisi.")
		}
		if row.Email == "" {
			errorsForRow = append(errorsForRow, "Email wajib diisi.")
		} else if _, err := mail.ParseAddress(row.Email); err != nil {
			errorsForRow = append(errorsForRow, "Format email tidak valid.")
		}
		if duplicateEmails[row.Email] > 1 {
			errorsForRow = append(errorsForRow, "Email duplikat di file import.")
		}
		if row.Role == "" {
			errorsForRow = append(errorsForRow, "Peran wajib diisi.")
		} else if !allowedSchoolMemberImportRoles[row.Role] {
			errorsForRow = append(errorsForRow, "Peran hanya boleh student, teacher, atau admin.")
		}
		if row.ClassCode != "" {
			if row.Role != "student" {
				errorsForRow = append(errorsForRow, "classCode hanya berlaku untuk peran student.")
			}
			if !existingClassCodes[row.ClassCode] {
				errorsForRow = append(errorsForRow, "Kode kelas tidak ditemukan di sekolah aktif.")
			}
		}

		status := "valid"
		if len(errorsForRow) > 0 {
			status = "invalid"
			response.InvalidCount++
		} else {
			response.ValidCount++
		}
		response.Rows = append(response.Rows, dto.AdminSchoolMemberImportRowDTO{
			RowNumber: row.RowNumber,
			FullName:  row.FullName,
			Email:     row.Email,
			Role:      row.Role,
			ClassCode: row.ClassCode,
			Status:    status,
			Errors:    errorsForRow,
		})
	}

	return response, nil
}

func (s *adminSchoolMemberImportService) existingClassCodes(schoolID string, classCodes map[string]bool) (map[string]bool, error) {
	result := map[string]bool{}
	if len(classCodes) == 0 {
		return result, nil
	}

	codes := make([]string, 0, len(classCodes))
	for code := range classCodes {
		codes = append(codes, code)
	}

	var found []string
	if err := s.db.Model(&domain.Class{}).
		Where("cls_sch_id = ? AND cls_code IN ? AND deleted_at IS NULL", schoolID, codes).
		Pluck("cls_code", &found).Error; err != nil {
		return nil, err
	}
	for _, code := range found {
		result[code] = true
	}
	return result, nil
}

func (s *adminSchoolMemberImportService) findOrCreateUser(tx *gorm.DB, fullName string, email string, defaultPassword string) (*domain.User, bool, error) {
	var user domain.User
	err := tx.Where("LOWER(usr_email) = ?", strings.ToLower(email)).First(&user).Error
	if err == nil {
		if !user.IsActive {
			return nil, false, errors.New("akun global tidak aktif")
		}
		return &user, false, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, false, err
	}
	user = domain.User{
		FullName: strings.TrimSpace(fullName),
		Email:    strings.ToLower(strings.TrimSpace(email)),
		Password: string(hashedPassword),
		IsActive: true,
	}
	if err := tx.Create(&user).Error; err != nil {
		return nil, false, err
	}
	return &user, true, nil
}

func (s *adminSchoolMemberImportService) findOrCreateSchoolUser(tx *gorm.DB, schoolID string, userID string) (*domain.SchoolUser, bool, error) {
	var schoolUser domain.SchoolUser
	err := tx.Where("scu_sch_id = ? AND scu_usr_id = ?", schoolID, userID).First(&schoolUser).Error
	if err == nil {
		return &schoolUser, false, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, false, err
	}
	schoolUser = domain.SchoolUser{
		SchoolID: schoolID,
		UserID:   userID,
	}
	if err := tx.Create(&schoolUser).Error; err != nil {
		return nil, false, err
	}
	return &schoolUser, true, nil
}

func (s *adminSchoolMemberImportService) findRoleID(tx *gorm.DB, roleName string) (string, error) {
	var role domain.Role
	if err := tx.Where("rol_name = ?", roleName).First(&role).Error; err != nil {
		return "", err
	}
	return role.ID, nil
}

func (s *adminSchoolMemberImportService) ensureRole(tx *gorm.DB, schoolUserID string, roleID string) (bool, error) {
	userRole := domain.UserRole{
		SchoolUserID: schoolUserID,
		RoleID:       roleID,
	}
	result := tx.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "urol_scu_id"}, {Name: "urol_rol_id"}},
		DoNothing: true,
	}).Create(&userRole)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}

func (s *adminSchoolMemberImportService) findClassIDByCode(tx *gorm.DB, schoolID string, classCode string) (string, error) {
	var class domain.Class
	if err := tx.Where("cls_sch_id = ? AND cls_code = ? AND deleted_at IS NULL", schoolID, classCode).First(&class).Error; err != nil {
		return "", err
	}
	return class.ID, nil
}

func (s *adminSchoolMemberImportService) ensureActiveStudentEnrollment(tx *gorm.DB, schoolID string, schoolUserID string, classID string) (bool, error) {
	var enrollment domain.Enrollment
	err := tx.Where("enr_scu_id = ? AND enr_cls_id = ?", schoolUserID, classID).First(&enrollment).Error
	if err == nil {
		if enrollment.LeftAt == nil && enrollment.Role == "student" {
			return false, nil
		}
		if err := tx.Model(&domain.Enrollment{}).
			Where("enr_id = ?", enrollment.ID).
			Updates(map[string]any{"enr_role": "student", "left_at": nil}).Error; err != nil {
			return false, err
		}
		return true, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, err
	}
	enrollment = domain.Enrollment{
		SchoolID:     schoolID,
		SchoolUserID: schoolUserID,
		ClassID:      classID,
		Role:         "student",
	}
	if err := tx.Create(&enrollment).Error; err != nil {
		return false, err
	}
	return true, nil
}

func csvValue(record []string, index int) string {
	if index < 0 || index >= len(record) {
		return ""
	}
	return strings.TrimSpace(record[index])
}

func isEmptyCSVRecord(record []string) bool {
	for _, value := range record {
		if strings.TrimSpace(value) != "" {
			return false
		}
	}
	return true
}
