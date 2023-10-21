package usecases

import (
	"github.com/google/uuid"
	"mind-sharer/interface_adapters"
	"mind-sharer/internal/domain/models"
)

type UserUseCase struct {
	userRepo interface_adapters.UsersRepository
}

func NewUsersUseCase(repo interface_adapters.UsersRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: repo,
	}
}

func (uc *UserUseCase) SearchUsers() ([]models.User, error) {
	return uc.userRepo.SearchUsers()
}

func (uc *UserUseCase) GetUserByID(id uuid.UUID) (models.User, error) {
	return uc.userRepo.FetchUserById(id)
}

func (uc *UserUseCase) RegisterUser(user models.User) (models.User, error) {
	// Add validation and business logic here
	// e.g., password hashing, validation checks, etc.
	return uc.userRepo.CreateUser(user)
}

func (uc *UserUseCase) UpdateUser(user models.User) (models.User, error) {
	return uc.userRepo.UpdateUser(user)
}

func (uc *UserUseCase) DeleteUser(id uuid.UUID) (models.User, error) {
	return uc.userRepo.DeleteUser(id)
}
