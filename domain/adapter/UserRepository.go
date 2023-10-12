package adapter

import (
	"github.com/google/uuid"
	"mind-sharer/domain/models"
)

type UserRepository interface {
	SearchUsers() ([]models.User, error)
	FetchUserById(id uuid.UUID) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uuid.UUID) (models.User, error)
	//GetByID(id uuid.UUID) (models.User, error)
	// Add other methods as needed
}
