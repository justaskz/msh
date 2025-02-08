package commands

import (
	"fmt"
	"os"

	"github.com/justaskz/mashina/internal/repos"
)

func Info(repoName string) {
	owner, name, err := repos.RepoName(repoName).Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	releases, err := repos.GetReleases(owner, name)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Versions of", repoName, ":")
	for _, release := range releases {
		fmt.Println(*release.Name)
	}
}
