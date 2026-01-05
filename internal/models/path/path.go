package path

type Path string

func (path Path) String() string {
	return string(path)
}
