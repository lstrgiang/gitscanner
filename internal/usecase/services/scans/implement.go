package scans

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/lstrgiang/gitscan/internal/infra/models"
	"github.com/lstrgiang/gitscan/internal/usecase/tasks"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type scanService struct {
	ctx         context.Context
	exec        boil.ContextExecutor
	asyncClient *asynq.Client
}

func (s scanService) CreateNewScan(repository *models.Repository) (*models.Scan, error) {
	newScan := &models.Scan{
		RepositoryID: repository.ID,
	}
	if err := newScan.Insert(s.ctx, s.exec, boil.Infer()); err != nil {
		return nil, err
	}
	return newScan, nil
}

func (s scanService) TriggerScan(repository *models.Repository, scan *models.Scan, dbUrl string) error {
	// Define tasks.
	task := tasks.NewScanTask(repository, scan, dbUrl)

	// Process the task immediately in critical queue.
	if _, err := s.asyncClient.Enqueue(
		task,                    // task payload
		asynq.Queue("critical"), // set queue for task
	); err != nil {
		return err
	}
	return nil
}

func (s scanService) ListScansByStatus(status int, page int, limit int) ([]*models.Scan, error) {
	scanList, err := models.Scans(
		models.ScanWhere.Status.EQ(status),
		qm.Offset(page*limit-limit),
		qm.Limit(limit),
		qm.Load(models.ScanRels.Repository),
	).All(s.ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return scanList, nil
}

func (s scanService) ListScans(page int, limit int) ([]*models.Scan, error) {
	scanList, err := models.Scans(
		qm.Offset(page*limit-limit),
		qm.Limit(limit),
		qm.Load(models.ScanRels.Repository),
	).All(s.ctx, s.exec)
	if err != nil {
		return nil, err
	}
	return scanList, nil

}
func (s scanService) FindScanById(id int) (*models.Scan, error) {
	scan, err := models.Scans(
		models.ScanWhere.ID.EQ(id),
		qm.Load(models.ScanRels.Repository),
	).One(s.ctx, s.exec)

	if err != nil {
		return nil, err
	}

	return scan, nil
}
