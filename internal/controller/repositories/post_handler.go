package repositories

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
	repository, err := repositoriesService.CreateNewRepository(
		request.GetName(),
		request.GetLink(),
	)
	if err != nil {
		tx.Rollback()
		return errors.InternalServerErr
	}

	tx.Commit()

	return e.JSON(http.StatusCreated, repository)
}
