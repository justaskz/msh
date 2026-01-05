package dev

import (
	"fmt"
	"os"
	"path/filepath"

	r "github.com/justaskz/mashina/internal/repos"
)

func UpdateSymlink(repo r.Repo) error {
	destination := filepath.Join(GetLibPath(), repo.Uid().String())
	binPath, err := findBin(destination, repo.Name)
	if err != nil {
		return err
	}

	fmt.Println("binpath:", binPath)
	symlink := filepath.Join(GetBinPath(), repo.Name)
	err = os.Symlink(binPath, symlink)
	if err != nil {
		return err
	}

	return nil
}

func findBin(folderPath string, filename string) (string, error) {
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == filename {
			return fmt.Errorf("found")
		}
		return nil
	})

	if err != nil && err.Error() == "found" {
		return filepath.Join(folderPath, filename), nil
	}
	return "", fmt.Errorf("file not found: %s", filename)
}
