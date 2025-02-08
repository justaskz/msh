package repos

import (
	"context"

	"github.com/google/go-github/v69/github"
)

func GetReleases(owner string, name string) ([]*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	bg := context.Background()
	releases, _, err := client.Repositories.ListReleases(bg, owner, name, nil)

	if err != nil {
		return nil, err
	}

	return releases, nil
}
