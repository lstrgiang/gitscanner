package repositories

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/lstrgiang/gitscan/internal/infra/inject"
	"github.com/lstrgiang/gitscan/internal/usecase/services"
)

func NewGetHandler() GetHandler {
	return &getHandler{
		NewRequest: NewGetRequest,
	}
}

type GetHandler interface {
	Handle(echo.Context) error
}

type getHandler struct {
	NewRequest func() GetRequest
}

func (h getHandler) Handle(e echo.Context) error {
	request := h.NewRequest()

	if err := request.Bind(e); err != nil {
		return errors.InvalidRequest
	}

	if err := request.Validate(); err != nil {
		return err
	}

	if request.IsPagination() {
		return h.handlePagination(request, e)
	}

	return h.handleGetOne(request, e)
}

// handle getting pagination
func (h getHandler) handlePagination(request GetRequest, e echo.Context) error {
	ctx, tx := inject.CtxTx(e)
	repositoriesService := services.GetRepositoryService(ctx, tx)

	models, err := repositoriesService.GetPagination(
		request.GetPagination(),
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return e.JSON(http.StatusOK, NewResponses(models))
}

// handle get one by id
func (h getHandler) handleGetOne(request GetRequest, e echo.Context) error {
	ctx, tx := inject.CtxTx(e)
	repositoriesService := services.GetRepositoryService(ctx, tx)

	model, err := repositoriesService.FindByID(request.GetID())
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return e.JSON(http.StatusOK, NewResponse(model))
}
