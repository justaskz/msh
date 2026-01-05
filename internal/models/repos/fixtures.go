package repos

import r "github.com/justaskz/mashina/internal/models/repo"

type fixtures struct {
	ConfigRepos    Repos
	InstalledRepos Repos
}

func newFixtures() fixtures {
	return fixtures{
		ConfigRepos:    createConfigRepos(),
		InstalledRepos: createInstalledRepos(),
	}
}

func createConfigRepos() Repos {
	repos := New()
	repos = repos.Add(r.Fixtures.Helix)
	repos = repos.Add(r.Fixtures.Starship)

	return repos
}

func createInstalledRepos() Repos {
	repos := New()
	repos = repos.Add(r.Fixtures.Helix)

	return repos
}

var Fixtures fixtures = newFixtures()
