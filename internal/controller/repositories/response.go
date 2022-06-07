package repositories

import (
	"time"

	"github.com/lstrgiang/gitscan/internal/infra/models"
	"github.com/volatiletech/null/v8"
)

type Repository struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Link      string    `json:"link"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt null.Time `json:"deleted_at"`
}

func NewResponse(repository *models.Repository) *Repository {
	return &Repository{
		ID:        repository.ID,
		Name:      repository.Name,
		Link:      repository.Link,
		CreatedAt: repository.CreatedAt,
		UpdatedAt: repository.UpdatedAt,
		DeletedAt: repository.DeletedAt,
	}
}

func NewResponses(repositories []*models.Repository) []*Repository {
	responses := make([]*Repository, 0)
	for _, repository := range repositories {
		responses = append(responses, NewResponse(repository))
	}
	return responses
}
