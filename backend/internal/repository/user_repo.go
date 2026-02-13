package repository

import (
	"backend/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id string) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
	Update(user *domain.User) error
	Delete(id string) error
	CheckEmailExists(email string, excludeID string) (bool, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("usr_id = ?", id).First(&user).Error
	return &user, err
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("usr_email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) Update(user *domain.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *userRepository) Delete(id string) error {
	result := r.db.Delete(&domain.User{}, "usr_id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *userRepository) CheckEmailExists(email string, excludeID string) (bool, error) {
	var count int64
	query := r.db.Model(&domain.User{}).Where("usr_email = ?", email)
	if excludeID != "" {
		query = query.Where("usr_id != ?", excludeID)
	}
	err := query.Count(&count).Error
	return count > 0, err
}
