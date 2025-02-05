package wip

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Pack struct {
	Name    string
	Version string
}

type Packs []Pack

type Config struct {
	Aliases []string `yaml:"aliases"`
	Envs    []string `yaml:"envs"`
	Packs   []Pack   `yaml:"packs"`
}

func ReadUserConfig(path string) Config {
	var config Config
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}
