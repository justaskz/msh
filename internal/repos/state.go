package repos

type State int

const (
	Missing State = iota
	Installed
)

func (state State) String() string {
	switch state {
	case 0:
		return "missing"
	case 1:
		return "installed"
	default:
		return "missing"
	}
}
