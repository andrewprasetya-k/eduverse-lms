package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type FeedRepository interface {
	Create(feed *domain.Feed) error
	GetByClass(classID string, page int, limit int) ([]*domain.Feed, int64, error)
	GetByClassInSchool(classID string, schoolID string, page int, limit int) ([]*domain.Feed, int64, error)
	GetByID(id string) (*domain.Feed, error)
	GetByIDInSchool(id string, schoolID string) (*domain.Feed, error)
	UpdateInSchool(feed *domain.Feed, schoolID string) error
	DeleteInSchool(id string, schoolID string) error
}

type feedRepository struct {
	db *gorm.DB
}

func NewFeedRepository(db *gorm.DB) FeedRepository {
	return &feedRepository{db: db}
}

func (r *feedRepository) Create(feed *domain.Feed) error {
	return r.db.Create(feed).Error
}

func (r *feedRepository) GetByClass(classID string, page int, limit int) ([]*domain.Feed, int64, error) {
	return r.GetByClassInSchool(classID, "", page, limit)
}

func (r *feedRepository) GetByClassInSchool(classID string, schoolID string, page int, limit int) ([]*domain.Feed, int64, error) {
	var feeds []*domain.Feed
	var total int64

	query := r.db.Model(&domain.Feed{}).Preload("Creator").Where("fds_cls_id = ?", classID)
	if schoolID != "" {
		query = query.Where("fds_sch_id = ?", schoolID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Order("created_at desc").Find(&feeds).Error
	return feeds, total, err
}

func (r *feedRepository) GetByID(id string) (*domain.Feed, error) {
	return r.GetByIDInSchool(id, "")
}

func (r *feedRepository) GetByIDInSchool(id string, schoolID string) (*domain.Feed, error) {
	var feed domain.Feed
	query := r.db.Preload("Creator").Where("fds_id = ?", id)
	if schoolID != "" {
		query = query.Where("fds_sch_id = ?", schoolID)
	}
	err := query.First(&feed).Error
	return &feed, err
}

func (r *feedRepository) UpdateInSchool(feed *domain.Feed, schoolID string) error {
	result := r.db.Model(&domain.Feed{}).
		Where("fds_id = ? AND fds_sch_id = ?", feed.ID, schoolID).
		Updates(feed)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *feedRepository) DeleteInSchool(id string, schoolID string) error {
	result := r.db.Where("fds_id = ? AND fds_sch_id = ?", id, schoolID).
		Delete(&domain.Feed{})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
