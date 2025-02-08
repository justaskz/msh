package dev

import (
	"os"

	r "github.com/justaskz/mashina/internal/repos"
	"gopkg.in/yaml.v3"
)

type configFile struct {
	Repos repoConfig `yaml:"repos"`
}

type repoConfig map[string]string

func ReadConfig(path string) (r.Repos, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		return r.CreateRepos(), err
	}

	var config configFile
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return r.CreateRepos(), err
	}

	repos := r.CreateRepos()
	for repoName, version := range config.Repos {
		owner, name, err := r.RepoName(repoName).Parse()
		if err != nil {
			return r.CreateRepos(), err
		}

		repo, err := r.CreateRepo(string(owner), string(name), version)
		if err != nil {
			return r.Repos{}, err
		}

		repos = repos.Add(repo)
	}

	return repos, nil
}
