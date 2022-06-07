package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
)

func NewGetRequest() GetRequest {
	return &getRequest{}
}

type GetRequest interface {
	Bind(echo.Context) error
	Validate() error
	IsPagination() bool
	GetID() int
	GetPagination() (int, int)
}

type getRequest struct {
	ID    int `param:"id"`
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

func (r *getRequest) Bind(e echo.Context) error {
	return e.Bind(r)
}

func (r getRequest) Validate() error {
	// if r.ID is not provided
	if r.ID == 0 {
		// if pagination not provided as well
		if r.Page == 0 || r.Limit == 0 {
			return errors.InvalidRequest
		}
		if r.Page < 0 {
			return errors.NewParamErr("page must be non negative")
		}
		if r.Limit < 0 {
			return errors.NewParamErr("limit must be non negative")
		}
	}

	if r.ID < 0 {
		return errors.InvalidIdError
	}

	return nil
}

func (r getRequest) IsPagination() bool {
	return r.ID == 0
}

func (r getRequest) GetID() int {
	return r.ID
}

func (r getRequest) GetPagination() (int, int) {
	return r.Page, r.Limit
}
