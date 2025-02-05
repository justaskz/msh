package wip

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/justaskz/mashina/config"
)

func Init() string {
	initScript := loadInitScript()
	aliases := loadAliases()
	return strings.Join([]string{initScript, aliases}, "\n")
}

func loadInitScript() string {
	initScript, err := fs.ReadFile(config.EmbedFiles, "init.sh")
	if err != nil {
		panic(err)
	}

	return string(initScript)
}

func loadAliases() string {
	var output []string
	xdgConfigHome := getXDGConfigPath()
	config := ReadUserConfig(xdgConfigHome + "/mashina/config.yaml")
	for _, aliasFilePath := range config.Aliases {
		script, err := os.ReadFile(aliasFilePath)
		if err != nil {
			panic(err)
		}

		output = append(output, string(script))
	}

	return strings.Join(output, "\n")
}

func getXDGConfigPath() string {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		home := os.Getenv("HOME")
		xdgConfigHome = filepath.Join(home, ".config")
	}

	return xdgConfigHome
}
