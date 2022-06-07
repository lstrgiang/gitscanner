package scans

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/lstrgiang/gitscan/internal/infra/inject"
	"github.com/lstrgiang/gitscan/internal/usecase/services"
)

func NewPostHandler() PostHandler {
	return &postHandler{
		NewRequest: NewPostRequest,
	}

}

type PostHandler interface {
	Handle(echo.Context) error
}

type postHandler struct {
	NewRequest func() PostRequest
}

func (h postHandler) Handle(e echo.Context) error {
	request := h.NewRequest()
	if err := request.Bind(e); err != nil {
		return errors.InvalidRequest
	}
	if err := request.Validate(); err != nil {
		return err
	}
	ctx, tx := inject.CtxTx(e)
	repositoriesService := services.GetRepositoryService(ctx, tx)
	asynqClient := inject.AsynqClient(e)
	scanService := services.GetScanService(ctx, tx, asynqClient)

	repo, err := repositoriesService.FindByID(request.GetID())
	if err != nil {
		tx.Rollback()
		return err
	}

	scan, err := scanService.CreateNewScan(repo)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return errors.InternalServerErr
	}

	if err := scanService.TriggerScan(repo, scan, inject.DBUrl(e)); err != nil {
		return err
	}
	return e.JSON(http.StatusCreated, NewResponseWithRepo(scan, repo))
}
