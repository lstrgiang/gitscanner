package tasks

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/hibiken/asynq"
	"github.com/lstrgiang/gitscan/internal/infra/db"
	"github.com/lstrgiang/gitscan/internal/infra/models"
	"github.com/lstrgiang/gitscan/internal/usecase/consts"
	gitScanner "github.com/lstrgiang/gitscan/internal/usecase/git_scanner"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// HandleWelcomeEmailTask handler for welcome email task.
func HandleScan(c context.Context, t *asynq.Task) error {
	payload := &PayloadData{}

	if err := json.Unmarshal(t.Payload(), payload); err != nil {
		return err
	}
	db, err := db.NewDB(payload.DbUrl)
	if err != nil {
		return err
	}

	scan := payload.Scan

	scan.ScannedAt = null.TimeFrom(time.Now())
	if err := updateScanWithStatus(c, db, scan, consts.ScanStatusInProgress); err != nil {
		return err
	}

	result, err := scanRepository(payload.Repository)
	if err != nil {
		scan.FinishedAt = null.TimeFrom(time.Now())
		if err := updateScanWithStatus(c, db, scan, consts.ScanStatusFailure); err != nil {
			return err
		}
		return err
	}
	resultByte, _ := json.Marshal(result)
	scan.Findings = null.JSONFrom(resultByte)
	scan.FinishedAt = null.TimeFrom(time.Now())

	if err := updateScanWithStatus(c, db, scan, consts.ScanStatusSuccess); err != nil {
		return err
	}

	return nil
}

func scanRepository(repository *models.Repository) (interface{}, error) {
	//clone repository
	result, err := gitScanner.ScanRepository(repository.Link)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func updateScanWithStatus(ctx context.Context, db *sql.DB, scan *models.Scan, status int) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}
	scan.Status = status
	_, err = scan.Update(ctx, tx, boil.Whitelist(
		models.ScanColumns.Status,
		models.ScanColumns.ScannedAt,
		models.ScanColumns.FinishedAt,
		models.ScanColumns.Findings,
	))

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
