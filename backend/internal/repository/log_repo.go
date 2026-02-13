package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type LogRepository interface {
	Create(log *domain.Log) error
	GetBySchool(schoolID string, page int, limit int) ([]*domain.Log, int64, error)
	GetByUser(userID string, page int, limit int) ([]*domain.Log, int64, error)
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) Create(log *domain.Log) error {
	return r.db.Create(log).Error
}

func (r *logRepository) GetBySchool(schoolID string, page int, limit int) ([]*domain.Log, int64, error) {
	var logs []*domain.Log
	var total int64

	query := r.db.Model(&domain.Log{}).Preload("User").Where("log_sch_id = ?", schoolID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Order("created_at desc").Find(&logs).Error
	return logs, total, err
}

func (r *logRepository) GetByUser(userID string, page int, limit int) ([]*domain.Log, int64, error) {
	var logs []*domain.Log
	var total int64

	query := r.db.Model(&domain.Log{}).Where("log_usr_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Order("created_at desc").Find(&logs).Error
	return logs, total, err
}
