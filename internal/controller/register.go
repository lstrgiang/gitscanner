package controller

import (
	"path"

	"github.com/labstack/echo/v4"
	repositoriesHandler "github.com/lstrgiang/gitscan/internal/controller/repositories"
	"github.com/lstrgiang/gitscan/internal/controller/scans"
)

const (
	RepositoryPath = "repositories"
	ScanPath       = "scans"
)

func RegisterHandlers(e *echo.Echo, apiPrefix string) {
	repositoriesHandler.RegisterHandler(e, path.Join(apiPrefix, RepositoryPath))
	scans.RegisterHandler(e, path.Join(apiPrefix, ScanPath))
}
