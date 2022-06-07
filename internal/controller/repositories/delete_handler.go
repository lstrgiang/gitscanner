package repositories

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/lstrgiang/gitscan/internal/infra/inject"
	"github.com/lstrgiang/gitscan/internal/usecase/services"
)

func NewDeleteHandler() DeleteHandler {
	return &deleteHandler{
		NewRequest: NewDeleteRequest,
	}
}

type DeleteHandler interface {
	Handle(echo.Context) error
}

type deleteHandler struct {
	NewRequest func() DeleteRequest
}

func (h deleteHandler) Handle(e echo.Context) error {
	request := h.NewRequest()

	if err := request.Bind(e); err != nil {
		return errors.InvalidRequest
	}
	if err := request.Validate(); err != nil {
		return err
	}

	ctx, tx := inject.CtxTx(e)
	repositoriesService := services.GetRepositoryService(ctx, tx)
	if err := repositoriesService.DeleteByID(request.GetID()); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return e.NoContent(http.StatusOK)
}
