package repos

import (
	"fmt"
	"os"
	"path/filepath"

	p "github.com/justaskz/mashina/internal/models/path"
	r "github.com/justaskz/mashina/internal/models/repo"
)

func ReadInstalled(path p.Path) (Repos, error) {
	uids, err := listFolders(path)
	if err != nil {
		return Repos{}, nil
	}

	repos := New()

	for _, uid := range uids {
		repo, err := r.Parse(uid)
		if err != nil {
			return Repos{}, err
		}

		repos = repos.Add(repo)
	}

	return repos, nil
}

func listFolders(path p.Path) ([]r.Uid, error) {
	pathGlob := fmt.Sprintf("%s/*", path)
	items, _ := filepath.Glob(pathGlob)

	var uids []r.Uid
	for _, item := range items {
		fileInfo, err := os.Stat(item)
		if err != nil {
			return uids, err
		}

		if fileInfo.IsDir() {
			basename := filepath.Base(item)
			uids = append(uids, r.Uid(basename))
		}
	}

	return uids, nil
}
