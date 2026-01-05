package repos

import (
	"path/filepath"
	"strings"
)

type File string

func (file File) Extention() string {
	switch {
	case strings.Contains(file.String(), "tar.gz"):
		return "tar.gz"
	case strings.Contains(file.String(), "tar.xz"):
		return "tar.xz"
	default:
		return filepath.Ext(string(file))
	}
}

func (file File) String() string {
	return string(file)
}
