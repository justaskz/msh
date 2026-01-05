package releasez

import (
	"context"
	"fmt"

	"github.com/google/go-github/v70/github"
	"github.com/justaskz/mashina/internal/models/repo"
)

type Release struct {
	Name string
}

func Get(repo repo.Repo) (Release, error) {
	releases, err := getGithub(repo)

	if err != nil {
		return Release{}, err
	}

	return releases[0], nil
}

func getGithub(repo repo.Repo) ([]Release, error) {
	client := github.NewClient(nil)
	bg := context.Background()
	owner := repo.Owner.String()
	name := repo.Name.String()
	githubReleases, _, err := client.Repositories.ListReleases(bg, owner, name, nil)

	if err != nil {
		return nil, err
	}

	var releases = []Release{}
	for _, githubRelease := range githubReleases {
		fmt.Println(*githubRelease.Name)
		// release := release.Release{}

		// 		releases = append(releases, release)
	}

	return releases, nil
}
