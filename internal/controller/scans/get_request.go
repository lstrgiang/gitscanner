package scans

import (
	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
	"github.com/lstrgiang/gitscan/internal/usecase/consts"
)

func NewGetRequest() GetRequest {
	return &getRequest{}
}

type GetRequest interface {
	Bind(echo.Context) error
	Validate() error
	GetID() int
	GetStatus() int
	GetPage() int
	GetLimit() int
	IsGetOne() bool
	IsStatusProvided() bool
}

type getRequest struct {
	ID     int    `param:"id"`
	Status string `query:"status"`
	Page   int    `query:"page"`
	Limit  int    `query:"limit"`
}

func (r *getRequest) Bind(e echo.Context) error {
	return e.Bind(r)
}

func (r getRequest) Validate() error {
	if r.Status != "" && !consts.IsScanStatusNameValid(r.Status) {
		return errors.NewParamErr("invalid status")
	}
	return nil
}

func (r getRequest) GetID() int {
	return r.ID
}

func (r getRequest) GetStatus() int {
	return int(consts.GetStatusVal(r.Status))
}

func (r getRequest) GetPage() int {
	return r.Page
}

func (r getRequest) GetLimit() int {
	return r.Limit
}

func (r getRequest) IsGetOne() bool {
	return r.ID > 0
}

func (r getRequest) IsStatusProvided() bool {
	return r.Status != ""
}
