package interface_adapters

import (
	"github.com/google/uuid"
	"mind-sharer/internal/domain/models"
)

type UsersRepository interface {
	SearchUsers() ([]models.User, error)
	FetchUserById(id uuid.UUID) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id uuid.UUID) (models.User, error)
}
