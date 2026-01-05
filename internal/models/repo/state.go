package repo

type State string

const (
	missing   State = "missing"
	installed State = "installed"
)

type states struct {
	Missing   State
	Installed State
}

var States states = states{
	Missing:   missing,
	Installed: installed,
}

func (state State) String() string {
	return string(state)
}

func (state State) IsInstalled() bool {
	if state == installed {
		return true
	}

	return false
}
