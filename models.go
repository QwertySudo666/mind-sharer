package main

import (
	"github.com/google/uuid"
	"time"
)

type user struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	//Posts    []post    `json:"posts"`
}

type post struct {
	Id          uuid.UUID `json:"id,"`
	AuthorID    uuid.UUID `json:"author_id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
}
