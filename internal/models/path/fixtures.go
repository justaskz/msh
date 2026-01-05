package path

type fixtures struct {
	ConfigPath Path
	LibPath    Path
}

func newFixtures() fixtures {
	return fixtures{
		ConfigPath: Path("testdata/.config/msh/config.yaml"),
		LibPath:    Path("testdata/.local/lib/msh"),
	}
}

var Fixtures fixtures = newFixtures()
