package repositories

import "github.com/labstack/echo/v4"

func NewDeleteRequest() DeleteRequest {
	return &deleteRequest{}
}

type DeleteRequest interface {
	Bind(echo.Context) error
	Validate() error
	GetID() int
}

type deleteRequest struct {
	ID int `param:"id"`
}

func (r *deleteRequest) Bind(e echo.Context) error {
	return e.Bind(r)
}

func (r deleteRequest) Validate() error {
	if r.ID <= 0 {
		return nil
	}
	return nil
}

func (r deleteRequest) GetID() int {
	return r.ID
}
