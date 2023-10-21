package entities

import (
	"github.com/google/uuid"
	"time"
)

type PostEntity struct {
	Id          uuid.UUID
	Author      UserEntity `gorm:"foreignKey:ID"`
	Title       string
	Content     string
	PublishedAt time.Time
}
