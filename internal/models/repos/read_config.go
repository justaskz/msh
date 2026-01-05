package repos

import (
	"os"

	p "github.com/justaskz/mashina/internal/models/path"
	r "github.com/justaskz/mashina/internal/models/repo"
	"gopkg.in/yaml.v3"
)

type configFile struct {
	Repos repos `yaml:"repos"`
}

type repos map[r.RepoName]r.Version

func ReadConfig(path p.Path) (Repos, error) {
	yamlFile, err := os.ReadFile(path.String())
	if err != nil {
		return Repos{}, err
	}

	var cfg configFile
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return Repos{}, err
	}

	repos := New()
	for repoName, version := range cfg.Repos {
		owner, name, err := repoName.Parse()
		if err != nil {
			return Repos{}, err
		}

		repo, err := r.New(owner, name, version)
		if err != nil {
			return Repos{}, err
		}

		repos = repos.Add(repo)
	}

	return repos, nil
}
