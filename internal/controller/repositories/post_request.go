package repositories

import (
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/lstrgiang/gitscan/internal/infra/errors"
)

func NewPostRequest() PostRequest {
	return &postRequest{}
}

type PostRequest interface {
	Bind(echo.Context) error
	Validate() error
	GetName() string
	GetLink() string
}

type postRequest struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (r *postRequest) Bind(e echo.Context) error {
	return e.Bind(r)
}

func (r postRequest) Validate() error {
	if r.Name == "" {
		return errors.NewParamErr("name must be provided")
	}
	if r.Link == "" {
		return errors.NewParamErr("link must be provided")
	}
	if _, err := url.ParseRequestURI(r.Link); err != nil {
		return errors.InvalidUrlError
	}
	return nil
}

func (r postRequest) GetName() string {
	return r.Name
}

func (r postRequest) GetLink() string {
	return r.Link
}
