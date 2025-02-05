package commands

import (
	"context"
	"strings"

	"github.com/google/go-github/v35/github"
)

func Info(reponame string) ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	splitted := strings.Split(reponame, "/")
	owner := splitted[0]
	repo := splitted[1]
	releases, _, err := client.Repositories.ListReleases(context.Background(), owner, repo, nil)

	if err != nil {
		return nil, err
	}

	return releases, nil
}
