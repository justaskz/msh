package wip

import (
	"errors"
	"strings"
)

type Uid string

type Pack struct {
	Uid     Uid
	Owner   string
	Name    string
	Version string
}

func CreatePack(owner string, name string, version string) Pack {
	pack := Pack{
		Uid:     Uid("uid_"),
		Owner:   owner,
		Name:    name,
		Version: version,
	}

	return pack
}

func CreatePackFromFolderName(folderName string) (Pack, error) {
	parts := strings.Split(folderName, "__")

	if len(parts) != 3 {
		return Pack{}, errors.New("folder name format is incorrect")
	}

	pack := Pack{
		Owner:   parts[0],
		Name:    parts[1],
		Version: parts[2],
	}

	return pack, nil
}

type Packs map[Uid]Pack

func CreatePacks() Packs {
	return Packs{}
}

func (packs Packs) Add(pack Pack) Packs {
	packs[pack.Uid] = pack

	return packs
}
