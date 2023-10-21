package models

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	Id          uuid.UUID `json:"id,"`
	AuthorID    uuid.UUID `json:"author_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
}
