package repos

import (
	"fmt"
	"strings"
)

type Uid string
type State int

type Repo struct {
	Owner   string
	Name    string
	Version string
	State   State
}

const (
	Missing State = iota
	Installed
)

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
		State:   Missing,
	}

	return repo, nil
}

func (repo Repo) Uid() Uid {
	elements := []string{repo.Owner, repo.Name, repo.Version}
	uid := strings.Join(elements, "__")
	return Uid(uid)
}

func (repo Repo) UpdateState(state State) Repo {
	repo.State = state
	return repo
}

func (state State) String() string {
	switch state {
	case 0:
		return "missing"
	case 1:
		return "installed"
	default:
		return "missing"
	}
}

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
