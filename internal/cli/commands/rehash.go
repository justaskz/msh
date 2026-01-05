package commands

import (
	"fmt"
	"os"

	"github.com/justaskz/mashina/internal/dev"
)

func Rehash() {
	repos, err := dev.GetReposWithState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, repo := range repos {
		err = dev.UpdateSymlink(repo)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
