package repos

import (
	"fmt"
	"strings"
)

type Repo struct {
	Owner   string
	Name    string
	Version string
	State   State
	Url     Url
	File    File
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
