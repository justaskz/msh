package path

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPaths(t *testing.T) {
	Convey("", t, func() {
		So(Paths.MshPath.String(), ShouldContainSubstring, ".local/lib/msh")
		So(Paths.ConfigPath.String(), ShouldContainSubstring, ".config/msh/config.yaml")
		So(Paths.LibPath.String(), ShouldContainSubstring, ".local/lib/msh/lib")
	})
}
