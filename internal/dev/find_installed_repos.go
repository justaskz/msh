package dev

import (
	"github.com/justaskz/mashina/internal/repos"
	r "github.com/justaskz/mashina/internal/repos"
	"golang.org/x/exp/maps"
)

func FindInstalledRepos(configRepos r.Repos, installedRepos r.Repos) (r.Repos, error) {
	configUids := maps.Keys(configRepos)
	installedUids := maps.Keys(installedRepos)
	uids := intersection(configUids, installedUids)

	for _, uid := range uids {
		repo, err := configRepos.Get(uid)
		if err != nil {
			return r.Repos{}, err
		}

		repo.UpdateState(repos.Installed)
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
