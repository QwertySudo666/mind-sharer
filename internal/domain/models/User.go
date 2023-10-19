package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	//Posts    []Post    `json:"posts"`
	//Comments    []Comment    `json:"comments"`
	//Groups    []Group    `json:"groups"`
}
