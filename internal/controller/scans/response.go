package scans

import (
	"time"

	"github.com/lstrgiang/gitscan/internal/infra/models"
	"github.com/lstrgiang/gitscan/internal/usecase/consts"
	"github.com/volatiletech/null/v8"
)

type ScanResponse struct {
	ID             int       `json:"id"`
	Status         string    `json:"status"`
	RepositoryName string    `json:"repository_name"`
	RepositoryUrl  string    `json:"repository_url"`
	Findings       null.JSON `json:"findings"`
	QueuedAt       time.Time `json:"queued_at"`
	ScanningAt     null.Time `json:"scanning_at"`
	FinishedAt     null.Time `json:"finished_at"`
}

func NewResponse(scan *models.Scan) *ScanResponse {
	response := &ScanResponse{
		ID:         scan.ID,
		Status:     consts.ScanStatusName[consts.ScanStatus(scan.Status)],
		Findings:   scan.Findings,
		QueuedAt:   scan.CreatedAt,
		ScanningAt: scan.ScannedAt,
		FinishedAt: scan.FinishedAt,
	}
	if scan.R != nil {
		if scan.R.Repository != nil {
			response.RepositoryName = scan.R.Repository.Name
			response.RepositoryUrl = scan.R.Repository.Link
		}
	}
	return response

}

func NewResponses(scans []*models.Scan) []*ScanResponse {
	responses := make([]*ScanResponse, 0)
	for _, scan := range scans {
		responses = append(responses, NewResponse(scan))
	}
	return responses
}

func NewResponseWithRepo(scan *models.Scan, repo *models.Repository) *ScanResponse {
	response := NewResponse(scan)
	response.RepositoryName = repo.Name
	response.RepositoryUrl = repo.Link
	return response
}
