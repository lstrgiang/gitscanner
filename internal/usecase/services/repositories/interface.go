package repositories

import (
	"context"

	"github.com/lstrgiang/gitscan/internal/infra/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewRepositoryService(ctx context.Context, exec boil.ContextExecutor) RepositoryService {
	return &repositoryService{
		ctx:  ctx,
		exec: exec,
	}

}

type RepositoryService interface {
	CreateNewRepository(string, string) (*models.Repository, error)
	FindByID(int) (*models.Repository, error)
	FindByIDForUpdate(int) (*models.Repository, error)
	GetPagination(int, int) ([]*models.Repository, error)
	Update(*models.Repository) error
	DeleteByID(int) error
}
