package models

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID          uuid.UUID
	Content     string
	AuthorID    uuid.UUID
	PublishedAt time.Time
}
