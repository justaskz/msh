package repos

import (
	"context"

	"github.com/google/go-github/v69/github"
)

func GetReleases(repo Repo) ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	releases, _, err := client.Repositories.ListReleases(context.Background(), repo.Owner, repo.Name, nil)

	if err != nil {
		return nil, err
	}

	return releases, nil
}
