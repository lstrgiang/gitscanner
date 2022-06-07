package repositories

import (
	"path"

	"github.com/labstack/echo/v4"
)

func RegisterHandler(e *echo.Echo, apiPrefix string) {
	e.POST(apiPrefix, NewPostHandler().Handle)
	e.GET(apiPrefix, NewGetHandler().Handle)
	e.PUT(apiPrefix, NewPutHandler().Handle)

	// id handlers
	e.GET(path.Join(apiPrefix, ":id"), NewGetHandler().Handle)
	e.DELETE(path.Join(apiPrefix, ":id"), NewDeleteHandler().Handle)
}
