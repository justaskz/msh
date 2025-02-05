package cli

import (
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/justaskz/mashina/internal/commands"
)

var (
	app  = kingpin.New("msh", "Description msh")
	list = app.Command("list", "Description for list")

	info        = app.Command("info", "Description for info")
	infoPackage = info.Arg("package", "Package name").Required().String()
)

func Init(args []string) {
	switch kingpin.MustParse(app.Parse(args)) {

	case list.FullCommand():
		fmt.Println("list")

	case info.FullCommand():
		releases, _ := commands.Info(*infoPackage)

		for _, release := range releases {
			fmt.Println(*release.Name)
		}
	}
}
