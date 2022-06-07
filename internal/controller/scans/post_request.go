package scans

import (
	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
)

func NewPostRequest() PostRequest {
	return &postRequest{}
}

type PostRequest interface {
	Bind(echo.Context) error
	Validate() error
	GetID() int
}
type postRequest struct {
	ID int `param:"id"`
}

func (r *postRequest) Bind(e echo.Context) error {
	return e.Bind(r)
}

func (r postRequest) Validate() error {
	if r.ID <= 0 {
		return errors.InvalidIdError
	}
	return nil
}

func (r postRequest) GetID() int {
	return r.ID
}
