package persistence

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mind-sharer/domain/models"
)

type PersistentUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *PersistentUserRepository {
	return &PersistentUserRepository{
		db: db,
	}
}

func (r *PersistentUserRepository) SearchUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return users, err
}

func (r *PersistentUserRepository) FetchUserById(id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *PersistentUserRepository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *PersistentUserRepository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Updates(&user).Error
	return user, err
}

func (r *PersistentUserRepository) DeleteUser(id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	r.db.Delete(&user)
	return user, err
}
