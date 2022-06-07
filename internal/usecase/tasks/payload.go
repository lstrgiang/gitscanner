package tasks

import (
	"encoding/json"

	"github.com/hibiken/asynq"
	"github.com/lstrgiang/gitscan/internal/infra/models"
)

const (
	TypeNewScanTask = "task:scan"
)

type PayloadData struct {
	Repository *models.Repository
	Scan       *models.Scan
	DbUrl      string
}

func NewScanTask(repository *models.Repository, scan *models.Scan, dbUrl string) *asynq.Task {
	payloadData, _ := json.Marshal(PayloadData{
		Repository: repository,
		Scan:       scan,
		DbUrl:      dbUrl,
	})

	// Return a new task with given type and payload.
	return asynq.NewTask(TypeNewScanTask, payloadData)
}
