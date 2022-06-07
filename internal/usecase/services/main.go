package services

import (
	"github.com/lstrgiang/gitscan/internal/usecase/services/repositories"
	"github.com/lstrgiang/gitscan/internal/usecase/services/scans"
)

var GetRepositoryService = repositories.NewRepositoryService
var GetScanService = scans.NewScanService
