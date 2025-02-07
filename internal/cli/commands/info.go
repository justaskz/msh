package commands

import (
	"fmt"
	"os"

	"github.com/justaskz/mashina/internal/repos"
)

func Info(repoName string) {
	repo, err := repos.ParseRepoName(repoName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	releases, _ := repos.GetReleases(repo)

	fmt.Println("Versions of", repoName, ":")
	for _, release := range releases {
		fmt.Println(*release.Name)
	}
}
