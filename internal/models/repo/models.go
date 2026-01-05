package repo

import (
	"fmt"
	"strings"
)

type Owner string
type Name string
type Version string
type Uid string
type RepoName string

func (uid Uid) String() string {
	return string(uid)
}

func (owner Owner) String() string {
	return string(owner)
}

func (name Name) String() string {
	return string(name)
}

func (version Version) String() string {
	return string(version)
}

func (repoName RepoName) String() string {
	return string(repoName)
}

func (repoName RepoName) Parse() (Owner, Name, error) {
	splitted := strings.Split(repoName.String(), "/")

	if len(splitted) != 2 {
		return Owner(""), Name(""), fmt.Errorf("Repository name format is incorrect: '%s'", repoName)
	}

	return Owner(splitted[0]), Name(splitted[1]), nil
}
