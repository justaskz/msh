package repos

import (
	"fmt"
	"strings"
)

type Repo struct {
	Owner string
	Name  string
}

func ParseRepoName(repoName string) (Repo, error) {
	splitted := strings.Split(repoName, "/")

	if len(splitted) != 2 {
		return Repo{}, fmt.Errorf("Repository name format is incorrect: '%s'", repoName)
	}

	repo := Repo{
		Owner: splitted[0],
		Name:  splitted[1],
	}

	return repo, nil
}
