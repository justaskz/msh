package commands

import (
	"fmt"
	"os"

	"github.com/justaskz/mashina/internal/dev"
	"github.com/justaskz/mashina/internal/table"
	"github.com/justaskz/mashina/internal/wip"
)

func Status() {
	path := wip.GetConfigPath()
	repos, err := dev.ReadConfig(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var tableData [][]string
	for _, repo := range repos {
		row := []string{repo.Name, repo.Version}
		tableData = append(tableData, row)
	}

	table.PrintAlignedTable(tableData)
}
