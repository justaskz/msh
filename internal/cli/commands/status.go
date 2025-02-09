package commands

import (
	"fmt"
	"os"

	"github.com/justaskz/mashina/internal/dev"
	r "github.com/justaskz/mashina/internal/repos"
	"github.com/justaskz/mashina/internal/table"
)

func Status() {
	configPath := dev.GetConfigPath()
	repos, err := dev.ReadConfig(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	optPath := dev.GetOptPath()
	installedRepos, err := dev.ReadInstalled(optPath)
	repos, err = dev.FindInstalledRepos(repos, installedRepos)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printTable(repos)
}

func printTable(repos r.Repos) {
	var tableData [][]string
	for _, repo := range repos {
		row := []string{repo.Name, repo.Version, repo.State.String()}
		tableData = append(tableData, row)
	}

	table.PrintAlignedTable(tableData)
}
