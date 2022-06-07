package repositories

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/lstrgiang/gitscan/internal/infra/inject"
	"github.com/lstrgiang/gitscan/internal/usecase/services"
)

func NewPutHandler() PutHandler {
	return &putHandler{
		NewRequest: NewPutRequest,
	}
}

type PutHandler interface {
	Handle(echo.Context) error
}

type putHandler struct {
	NewRequest func() PutRequest
}

func (h putHandler) Handle(e echo.Context) error {
	request := h.NewRequest()

	if err := request.Bind(e); err != nil {
		return errors.InvalidRequest
	}

	if err := request.Validate(); err != nil {
		return err
	}

	ctx, tx := inject.CtxTx(e)
	repositoriesService := services.GetRepositoryService(ctx, tx)
	data, err := repositoriesService.FindByIDForUpdate(request.GetID())
	if err != nil {
		return err
	}

	data.Name = request.GetName()
	data.Link = request.GetLink()

	if err := repositoriesService.Update(data); err != nil {
		tx.Rollback()
		return errors.InternalServerErr
	}

	tx.Commit()

	return e.JSON(http.StatusOK, NewResponse(data))
}
