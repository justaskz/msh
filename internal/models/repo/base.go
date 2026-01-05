package repo

type baseString string

func (base baseString) String() string {
	return string(base)
}
