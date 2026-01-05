package repo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStates(t *testing.T) {
	Convey("", t, func() {
		So(States.Installed.String(), ShouldEqual, "installed")
		So(States.Missing.String(), ShouldEqual, "missing")
	})
}

func TestIsInstalled(t *testing.T) {
	Convey("", t, func() {
		So(States.Installed.IsInstalled(), ShouldEqual, true)
		So(States.Missing.IsInstalled(), ShouldEqual, false)
	})
}
