package repos

import (
	"fmt"
)

func BuildRepo(name string) Repo {
	dummy, _ := CreateRepo("owner", "repo", "1.2.3")
	helix, _ := CreateRepo("helix-editor", "helix", "24.07")

	collection := map[string]Repo{
		"dummy": dummy,
		"helix": helix,
	}

	object, ok := collection[name]
	if !ok {
		panic(fmt.Sprintf("fixture %s is undefined", name))
	}

	return object
}

func BuildRepos(name string) Repos {
	installed := CreateRepos()
	installed = installed.Add(BuildRepo("helix"))

	config := CreateRepos()
	config = config.Add(BuildRepo("dummy"))
	config = config.Add(BuildRepo("helix"))

	collection := map[string]Repos{
		"installed": installed,
		"config":    config,
	}

	object, ok := collection[name]
	if !ok {
		panic(fmt.Sprintf("fixture %s is undefined", name))
	}

	return object
}
