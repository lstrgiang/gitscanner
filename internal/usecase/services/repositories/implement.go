package repositories

import (
	"context"
	"database/sql"

	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/lstrgiang/gitscan/internal/infra/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	ctx  context.Context
	exec boil.ContextExecutor
}

var changeRepositoryColumes = boil.Whitelist(
	models.RepositoryColumns.Name,
	models.RepositoryColumns.Link,
)

func (s repositoryService) CreateNewRepository(name string, link string) (*models.Repository, error) {
	model := &models.Repository{
		Name: name,
		Link: link,
	}
	if err := model.Insert(s.ctx, s.exec, changeRepositoryColumes); err != nil {
		return nil, err
	}
	return model, nil
}

func (s repositoryService) FindByID(id int) (*models.Repository, error) {
	model, err := models.Repositories(models.RepositoryWhere.ID.EQ(id)).One(s.ctx, s.exec)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.DataNotFoundError
		}
		return nil, err
	}
	if model == nil {
		return nil, errors.DataNotFoundError
	}
	return model, nil
}

func (s repositoryService) FindByIDForUpdate(id int) (*models.Repository, error) {
	model, err := models.Repositories(
		models.RepositoryWhere.ID.EQ(id),
		qm.For("update"),
	).One(s.ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (s repositoryService) Update(repo *models.Repository) error {
	_, err := repo.Update(s.ctx, s.exec, changeRepositoryColumes)
	return err
}

func (s repositoryService) DeleteByID(id int) error {
	_, err := models.Repositories(
		models.RepositoryWhere.ID.EQ(id),
	).DeleteAll(s.ctx, s.exec, false)

	return err
}

func (s repositoryService) GetPagination(page int, limit int) ([]*models.Repository, error) {

	modelList, err := models.Repositories(
		qm.Offset(page*limit-limit),
		qm.Limit(limit),
	).All(s.ctx, s.exec)

	if err != nil {
		return nil, err
	}
	return modelList, nil
}
