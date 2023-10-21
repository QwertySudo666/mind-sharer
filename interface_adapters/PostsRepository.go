package interface_adapters

import (
	"github.com/google/uuid"
	"mind-sharer/internal/domain/models"
)

type PostsRepository interface {
	FetchAllPosts() ([]models.Post, error)
	SearchPosts() ([]models.Post, error)
	FetchPostsByAuthorId(authorID uuid.UUID) ([]models.Post, error)
	FetchPostById(id uuid.UUID) (models.Post, error)
	CreatePost(user models.Post) (models.Post, error)
	UpdatePost(user models.Post) (models.Post, error)
	DeletePost(id uuid.UUID) (models.Post, error)
}
