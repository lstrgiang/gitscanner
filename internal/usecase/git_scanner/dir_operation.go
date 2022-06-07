package git_scanner

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

func cloneRepository(repoUrl string, path string) error {
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL: repoUrl,
	})
	return err
}

func scanFile(path string, rootDir string) []ScanFinding {
	lines := make([]ScanFinding, 0)
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	currentLine := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Contains(line, "private_key") ||
			strings.Contains(line, "public_key") {
			lines = append(lines, NewSecretKeyFinding(
				currentLine,
				strings.ReplaceAll(path, rootDir, ""),
			))
		}
		currentLine++
	}

	readFile.Close()
	return lines
}

func listDir(path string) []string {
	fileList := make([]string, 0)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		//ignore .git dir
		if strings.Contains(path, ".git") {
			return nil
		}
		if !info.IsDir() {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList
}

func removeRepository(path string) error {
	return os.RemoveAll(path)
}
