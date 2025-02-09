package repos

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBuildRepo(t *testing.T) {
	Convey("", t, func() {
		object := BuildRepo("dummy")
		So(object.Owner, ShouldEqual, "owner")
	})
}
