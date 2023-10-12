package persistence

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mind-sharer/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user domain.User) (uuid.UUID, error) {
	err := r.db.Create(&user).Error
	return user.ID, err
}

func (r *UserRepository) GetByID(id uuid.UUID) (domain.User, error) {
	var user domain.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}
