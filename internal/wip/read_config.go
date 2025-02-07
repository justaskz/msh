package wip

import (
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type configFile struct {
	Packs PacksConfig `yaml:"packs"`
}

type PacksConfig map[string]string

func ReadConfig(path string) (Packs, error) {
	yamlFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var config configFile
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	packs := CreatePacks()
	for repo, version := range config.Packs {
		splitted := strings.Split(repo, "/")
		owner := splitted[0]
		name := splitted[1]
		pack := CreatePack(owner, name, version)
		packs = packs.Add(pack)
	}

	return packs, nil
}
