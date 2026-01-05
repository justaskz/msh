package path

import (
	"os"
	"path/filepath"
)

type paths struct {
	MshPath    Path
	ConfigPath Path
	LibPath    Path
}

func new() paths {
	return paths{
		MshPath:    getMshPath(),
		ConfigPath: getConfigPath(),
		LibPath:    getLibPath(),
	}
}

func getMshPath() Path {
	home := os.Getenv("HOME")
	path := filepath.Join(home, ".local", "lib", "msh")
	return Path(path)
}

func getConfigPath() Path {
	xdgConfigHome := getXDGConfigHomePath()
	configPath := filepath.Join(xdgConfigHome.String(), "msh", "config.yaml")
	return Path(configPath)
}

func getLibPath() Path {
	path := filepath.Join(getMshPath().String(), "lib")
	return Path(path)
}

func getXDGConfigHomePath() Path {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome != "" {
		return Path(xdgConfigHome)
	}

	home := os.Getenv("HOME")
	path := filepath.Join(home, ".config")
	return Path(path)
}

var Paths paths = new()
