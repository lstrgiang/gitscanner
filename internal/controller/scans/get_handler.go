package scans

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/lstrgiang/gitscan/internal/infra/inject"
	"github.com/lstrgiang/gitscan/internal/infra/models"
	"github.com/lstrgiang/gitscan/internal/usecase/services"
)

func NewGetHandler() GetHandler {
	return &getHandler{
		NewRequest: NewGetRequest,
	}
}

type GetHandler interface {
	Handle(e echo.Context) error
}

type getHandler struct {
	NewRequest func() GetRequest
}

func (h getHandler) Handle(e echo.Context) error {
	request := h.NewRequest()
	if err := request.Bind(e); err != nil {
		return err
	}
	if err := request.Validate(); err != nil {
		return err
	}

	if request.IsGetOne() {
		return h.getOne(e, request)
	}

	return h.getPagination(e, request)
}

func (h getHandler) getOne(e echo.Context, request GetRequest) error {
	db, ctx := inject.DBCtx(e)
	scanService := services.GetScanService(ctx, db, nil)
	scan, err := scanService.FindScanById(request.GetID())
	if err != nil {
		return errors.InternalServerErr
	}
	return e.JSON(http.StatusOK, NewResponse(scan))
}

func (h getHandler) getPagination(e echo.Context, request GetRequest) error {
	db, ctx := inject.DBCtx(e)
	scanService := services.GetScanService(ctx, db, nil)
	var scans []*models.Scan
	var err error
	if request.IsStatusProvided() {
		scans, err = scanService.ListScansByStatus(
			request.GetStatus(),
			request.GetPage(),
			request.GetLimit(),
		)
	} else {
		scans, err = scanService.ListScans(
			request.GetPage(),
			request.GetLimit(),
		)
	}
	if err != nil {
		return errors.InternalServerErr
	}
	return e.JSON(http.StatusOK, NewResponses(scans))

}
