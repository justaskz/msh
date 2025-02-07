package wip

import (
	"fmt"
	"os"
	"path/filepath"
)

func listFolders(path string) ([]string, error) {
	pathGlob := fmt.Sprintf("%s/*", path)
	items, _ := filepath.Glob(pathGlob)

	var folders []string
	for _, item := range items {
		f, _ := os.Stat(item)
		if f.IsDir() {
			basename := filepath.Base(item)
			folders = append(folders, basename)
		}
	}

	return folders, nil
}
