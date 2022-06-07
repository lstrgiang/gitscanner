package git_scanner

import (
	"fmt"

	"github.com/google/uuid"
)

func ScanRepository(repository string) (interface{}, error) {
	path := uuid.New().String()
	err := cloneRepository(repository, path)
	if err != nil {
		return nil, err
	}

	findings := make([]ScanFinding, 0)
	defer removeRepository(path)
	files := listDir(path)

	fmt.Println(fmt.Sprintf("Scanning %d files", len(files)))
	for _, file := range files {
		findings = append(findings, scanFile(file, path)...)
	}

	return findings, nil
}
