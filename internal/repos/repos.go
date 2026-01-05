package repos

import (
	"fmt"
	"strings"
)

type Repos map[Uid]Repo

func CreateRepos() Repos {
	return Repos{}
}

func (repos Repos) Add(repo Repo) Repos {
	repos[repo.Uid()] = repo

	return repos
}

func (repos Repos) Get(uid Uid) (Repo, error) {
	repo, ok := repos[uid]
	if !ok {
		return Repo{}, fmt.Errorf("repo with uid %s not found", uid)
	}

	return repo, nil
}

type RepoName string

func (repoName RepoName) Parse() (string, string, error) {
	splitted := strings.Split(string(repoName), "/")

	if len(splitted) != 2 {
		return string(""), string(""), fmt.Errorf("Repository name format is incorrect: '%s'", repoName)
	}

	return splitted[0], splitted[1], nil
}
