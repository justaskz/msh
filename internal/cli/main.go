package cli

import (
	"github.com/alecthomas/kingpin"
	"github.com/justaskz/mashina/internal/cli/commands"
)

var (
	app = kingpin.New("msh", "Description msh")

	status       = app.Command("status", "Description for status")
	info         = app.Command("info", "Description for info")
	infoRepoName = info.Arg("repo_name", "Repository name").Required().String()
)

func Init(args []string) {
	switch kingpin.MustParse(app.Parse(args)) {

	case status.FullCommand():
		commands.Status()

	case info.FullCommand():
		commands.Info(*infoRepoName)
	}
}

// func Status() {
// 	// home := os.Getenv("HOME")
// 	// path := filepath.Join(home, "bin/msh/opt")
// 	path := wip.GetConfigPath()

// 	// packs, err := wip.ReadInstalled(path)
// 	packs, err := wip.ReadConfig(path)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	// if len(packs) == 0 {
// 	// 	fmt.Println("No packs installed")
// 	// }

// 	for _, pack := range packs {
// 		fmt.Println(pack.Name)
// 	}
// }

// type Repos []Repo

// func readConfig(path string) Repos {
// 	// var config Config
// 	yamlFile, err := os.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = yaml.Unmarshal(yamlFile, &config)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return repos
// }
