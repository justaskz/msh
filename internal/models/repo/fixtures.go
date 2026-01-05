package repo

type fixtures struct {
	Helix    Repo
	Starship Repo
}

func newFixtures() fixtures {
	helix, _ := New("helix-editor", "helix", "24.07")
	starship, _ := New("starship", "starship", "1.17.1")

	return fixtures{
		Helix:    helix,
		Starship: starship,
	}
}

var Fixtures fixtures = newFixtures()
