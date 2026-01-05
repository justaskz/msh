package repos

import (
	r "github.com/justaskz/mashina/internal/models/repo"
)

func UpdateState(configRepos Repos, installedRepos Repos) (Repos, error) {
	uids := intersection(configRepos.Uids(), installedRepos.Uids())

	for _, uid := range uids {
		repo, err := configRepos.Get(uid)
		if err != nil {
			return Repos{}, err
		}

		repo = repo.SetState(r.States.Installed)
		configRepos.Add(repo)
	}

	return configRepos, nil
}

func intersection(first, second []r.Uid) []r.Uid {
	out := []r.Uid{}
	bucket := map[r.Uid]bool{}

	for _, i := range first {
		for _, j := range second {
			if i == j && !bucket[i] {
				out = append(out, i)
				bucket[i] = true
			}
		}
	}

	return out
}
