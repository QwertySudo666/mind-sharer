package datastorage

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"mind-sharer/internal/domain/models"
)

type UsersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) *UsersRepositoryImpl {
	return &UsersRepositoryImpl{
		db: db,
	}
}

func (r *UsersRepositoryImpl) SearchUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return users, err
}

func (r *UsersRepositoryImpl) FetchUserById(id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	return user, err
}

func (r *UsersRepositoryImpl) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *UsersRepositoryImpl) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Updates(&user).Error
	return user, err
}

func (r *UsersRepositoryImpl) DeleteUser(id uuid.UUID) (models.User, error) {
	var user models.User
	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	r.db.Delete(&user)
	return user, err
}
