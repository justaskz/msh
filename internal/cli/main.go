package cli

import (
	"github.com/alecthomas/kingpin"
	"github.com/justaskz/mashina/internal/cli/commands"
)

var (
	app = kingpin.New("msh", "Description msh")

	info         = app.Command("info", "Description for info")
	infoRepoName = info.Arg("repo_name", "Repository name").Required().String()

	// status = app.Command("status", "Description for status")
	// debug = app.Command("debug", "Description for debug")
)

func Init(args []string) {
	switch kingpin.MustParse(app.Parse(args)) {

	// case status.FullCommand():
	// 	Status()

	case info.FullCommand():
		commands.Info(*infoRepoName)

		// Info(infoPackage)
		// releases, _ := commands.Info(*infoPackage)

		// for _, release := range releases {
		// 	fmt.Println(*release.Name)
		// }

		// case debug.FullCommand():
		// 	fmt.Println("debug")
		// 	// result := wip.ReadInstalled("fixtures/msh/opt")
		// 	// result := wip.ReadInstalled("./*")
		// 	// fmt.Println(result)
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
