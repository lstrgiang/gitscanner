package consts

import "strings"

type ScanStatus int

const (
	ScanStatusQueued = iota
	ScanStatusInProgress
	ScanStatusSuccess
	ScanStatusFailure
)

var ScanStatusName = map[ScanStatus]string{
	ScanStatusQueued:     "Queued",
	ScanStatusInProgress: "In progress",
	ScanStatusSuccess:    "Success",
	ScanStatusFailure:    "Failure",
}

var ScanStatusMap = map[string]ScanStatus{
	"queued":      ScanStatusQueued,
	"in progress": ScanStatusInProgress,
	"success":     ScanStatusSuccess,
	"failure":     ScanStatusFailure,
}

func IsScanStatusNameValid(name string) bool {
	if _, ok := ScanStatusMap[strings.ToLower(name)]; !ok {
		return false
	}
	return true
}

func GetStatusVal(name string) ScanStatus {
	return ScanStatusMap[strings.ToLower(name)]
}
