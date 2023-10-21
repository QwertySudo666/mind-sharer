package entities

import "github.com/google/uuid"

type UserEntity struct {
	ID       uuid.UUID
	Email    string
	Username string
	Password string
}
