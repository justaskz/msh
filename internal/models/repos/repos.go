package repos

import (
	"fmt"

	r "github.com/justaskz/mashina/internal/models/repo"
	"golang.org/x/exp/maps"
)

type Repos map[r.Uid]r.Repo

func New() Repos {
	return Repos{}
}

func (repos Repos) Add(repo r.Repo) Repos {
	repos[repo.Uid()] = repo

	return repos
}

func (repos Repos) Uids() []r.Uid {
	return maps.Keys(repos)
}

func (repos Repos) Get(uid r.Uid) (r.Repo, error) {
	repo, ok := repos[uid]
	if !ok {
		return r.Repo{}, fmt.Errorf("repo with uid %s not found", uid)
	}

	return repo, nil
}

func (repos Repos) GetAll() []r.Repo {
	return maps.Values(repos)
}
