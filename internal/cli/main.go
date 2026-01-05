package cli

import (
	"github.com/alecthomas/kingpin"
	"github.com/justaskz/mashina/internal/cli/commands"
)

var (
	app = kingpin.New("msh", "Description msh")

	info         = app.Command("info", "Description for info")
	infoRepoName = info.Arg("repo_name", "Repository name").Required().String()
	install      = app.Command("install", "Description for install")
	status       = app.Command("status", "Description for status")
	rehash       = app.Command("rehash", "Description for rehash")
)

func Init(args []string) {
	switch kingpin.MustParse(app.Parse(args)) {
	case info.FullCommand():
		commands.Info(*infoRepoName)
	case install.FullCommand():
		commands.Install()
	case status.FullCommand():
		commands.Status()
	case rehash.FullCommand():
		commands.Rehash()
	}
}
