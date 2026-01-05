package repos

import "path/filepath"

type Url string

func (url Url) BasePath() string {
	return filepath.Base(url.String())
}

func (url Url) String() string {
	return string(url)
}
