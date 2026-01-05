package system

import "runtime"

type System struct {
	Arch     Arch
	Platform Platform
}

type Arch string
type Platform string

type archs struct {
	Arm64 Arch
	Amd64 Arch
}

type platforms struct {
	Linux  Platform
	Darwin Platform
}

var Archs archs = archs{
	Arm64: Arch("arm64"),
	Amd64: Arch("amd64"),
}

var Platforms platforms = platforms{
	Linux:  Platform("linux"),
	Darwin: Platform("darwin"),
}

func New() System {
	system := System{
		Arch:     Arch(runtime.GOARCH),
		Platform: Platform(runtime.GOOS),
	}

	return system
}
