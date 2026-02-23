package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type FeedRepository interface {
	Create(feed *domain.Feed) error
	GetByClass(classID string, page int, limit int) ([]*domain.Feed, int64, error)
	GetByID(id string) (*domain.Feed, error)
	Update(feed *domain.Feed) error
	Delete(id string) error
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
	var feeds []*domain.Feed
	var total int64

	query := r.db.Model(&domain.Feed{}).Preload("Creator").Where("fds_cls_id = ?", classID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err := query.Limit(limit).Offset(offset).Order("created_at desc").Find(&feeds).Error
	return feeds, total, err
}

func (r *feedRepository) GetByID(id string) (*domain.Feed, error) {
	var feed domain.Feed
	err := r.db.Preload("Creator").Where("fds_id = ?", id).First(&feed).Error
	return &feed, err
}

func (r *feedRepository) Update(feed *domain.Feed) error {
	result := r.db.Model(&domain.Feed{}).Where("fds_id = ?", feed.ID).Updates(feed)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}

func (r *feedRepository) Delete(id string) error {
	return r.db.Delete(&domain.Feed{}, "fds_id = ?", id).Error
}
