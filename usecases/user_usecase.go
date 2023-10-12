package usecases

import (
	"github.com/google/uuid"
	"mind-sharer/domain"
)

type UserRepository interface {
	Create(user domain.User) (uuid.UUID, error)
	GetByID(id uuid.UUID) (domain.User, error)
	// Add other methods as needed
}

type UserUseCase struct {
	userRepo UserRepository
}

func NewUserUseCase(repo UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

func (uc *UserUseCase) RegisterUser(user domain.User) (uuid.UUID, error) {
	// Add validation and business logic here
	// e.g., password hashing, validation checks, etc.
	return uc.userRepo.Create(user)
}

func (uc *UserUseCase) GetUserByID(id uuid.UUID) (domain.User, error) {
	return uc.userRepo.GetByID(id)
}
