package dev

import (
	"os"
	"path/filepath"
)

func GetMshPath() string {
	home := os.Getenv("HOME")
	return filepath.Join(home, ".local", "lib", "msh")
}

func GetBinPath() string {
	return filepath.Join(GetMshPath(), "bin")
}

func GetLibPath() string {
	return filepath.Join(GetMshPath(), "lib")
}

func GetConfigPath() string {
	xdgConfigHome := getXDGConfigHomePath()
	configPath := filepath.Join(xdgConfigHome, "msh", "config.yaml")
	return configPath
}

func GetOptPath() string {
	home := os.Getenv("HOME")
	return filepath.Join(home, ".msh", "opt")
}

func getXDGConfigHomePath() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome != "" {
		return xdgConfigHome
	}

	home := os.Getenv("HOME")
	return filepath.Join(home, ".config")
}
