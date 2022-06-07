package scans

import (
	"path"

	"github.com/labstack/echo/v4"
)

func RegisterHandler(e *echo.Echo, apiPrefix string) {
	// handle with id
	e.GET(apiPrefix, NewGetHandler().Handle)
	e.GET(path.Join(apiPrefix, ":id"), NewGetHandler().Handle)
	e.POST(path.Join(apiPrefix, ":id"), NewPostHandler().Handle)
}
