package dev

import (
	r "github.com/justaskz/mashina/internal/repos"
)

func GetReposWithState() (r.Repos, error) {
	configPath := GetConfigPath()
	repos, err := ReadConfig(configPath)
	if err != nil {
		return r.Repos{}, err
	}

	libPath := GetLibPath()
	installedRepos, err := ReadInstalled(libPath)
	if err != err {
		return r.Repos{}, err
	}

	updatedRepos, err := FindInstalledRepos(repos, installedRepos)
	if err != nil {
		return r.Repos{}, err
	}

	return updatedRepos, nil
}
