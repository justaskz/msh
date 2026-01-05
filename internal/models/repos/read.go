package repos

import p "github.com/justaskz/mashina/internal/models/path"

func Read(configPath p.Path, libPath p.Path) (Repos, error) {
	configRepos, err := ReadConfig(configPath)
	if err != nil {
		return Repos{}, err
	}

	installedRepos, err := ReadInstalled(libPath)
	if err != nil {
		return Repos{}, err
	}

	repos, err := UpdateState(configRepos, installedRepos)
	if err != nil {
		return Repos{}, err
	}

	return repos, nil
}
