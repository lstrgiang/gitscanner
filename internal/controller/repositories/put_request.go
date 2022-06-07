package repositories

import (
	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
)

func NewPutRequest() PutRequest {
	return &putRequest{}
}

type PutRequest interface {
	Bind(echo.Context) error
	Validate() error
	GetID() int
	GetName() string
	GetLink() string
}

type putRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

func (r *putRequest) Bind(e echo.Context) error {
	return e.Bind(r)
}

func (r putRequest) Validate() error {
	if r.ID <= 0 {
		return errors.InvalidIdError
	}
	if r.Name == "" {
		return errors.NewParamErr("name is missing")
	}
	if r.Link == "" {
		return errors.NewParamErr("link is missing")
	}

	return nil
}

func (r putRequest) GetName() string {
	return r.Name
}

func (r putRequest) GetLink() string {
	return r.Link
}

func (r putRequest) GetID() int {
	return r.ID
}
