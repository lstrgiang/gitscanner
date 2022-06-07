package controller

import (
	"github.com/labstack/echo/v4"
	repositoriesHandler "github.com/lstrgiang/gitscan/internal/controller/repositories"
)

func RegisterHandlers(e *echo.Echo, apiPrefix string) {
	repositoriesHandler.RegisterHandler(e, "/api/repositories")
}
