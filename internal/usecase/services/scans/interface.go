package scans

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/lstrgiang/gitscan/internal/infra/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func NewScanService(ctx context.Context, exec boil.ContextExecutor, asyncClient *asynq.Client) ScanService {
	return &scanService{
		ctx:         ctx,
		exec:        exec,
		asyncClient: asyncClient,
	}
}

type ScanService interface {
	CreateNewScan(repository *models.Repository) (*models.Scan, error)
	TriggerScan(repository *models.Repository, scan *models.Scan, dbUrl string) error
	ListScansByStatus(status int, page int, limit int) ([]*models.Scan, error)
	ListScans(page int, limit int) ([]*models.Scan, error)
	FindScanById(int) (*models.Scan, error)
}
