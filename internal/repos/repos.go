package repos

import (
	"fmt"
	"strings"
)

type Uid string

type Repo struct {
	Owner   string
	Name    string
	Version string
}

func CreateRepo(owner string, name string, version string) (Repo, error) {
	values := []string{owner, name, version}
	for _, value := range values {
		if value == "" {
			return Repo{}, fmt.Errorf("value is undefined")
		}
	}

	repo := Repo{
		Owner:   owner,
		Name:    name,
		Version: version,
	}

	return repo, nil
}

func (repo Repo) Uid() Uid {
	elements := []string{repo.Owner, repo.Name, repo.Version}
	uid := strings.Join(elements, "__")
	return Uid(uid)
}

type Repos map[Uid]Repo

func CreateRepos() Repos {
	return Repos{}
}

func (repos Repos) Add(repo Repo) Repos {
	repos[repo.Uid()] = repo

	return repos
}

type RepoName string

func (repoName RepoName) Parse() (string, string, error) {
	splitted := strings.Split(string(repoName), "/")

	if len(splitted) != 2 {
		return string(""), string(""), fmt.Errorf("Repository name format is incorrect: '%s'", repoName)
	}

	return splitted[0], splitted[1], nil
}
