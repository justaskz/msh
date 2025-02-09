package dev

import (
	"fmt"
	"os"
	"path/filepath"

	r "github.com/justaskz/mashina/internal/repos"
)

func ReadInstalled(path string) (r.Repos, error) {
	folderNames, _ := listFolders(path)
	repos := r.CreateRepos()

	for _, folderName := range folderNames {
		owner, name, version, err := ParseFolderName(folderName)
		if err != nil {
			return r.Repos{}, err
		}

		repo, err := r.CreateRepo(*owner, *name, *version)
		if err != nil {
			return r.Repos{}, err
		}

		repos = repos.Add(repo)
	}

	return repos, nil
}

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
