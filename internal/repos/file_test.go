package repos

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestName(t *testing.T) {
	Convey("", t, func() {
		file := File("fzf-0.56.0-darwin_arm64.tar.gz")
		So(file.Extention(), ShouldEqual, "tar.gz")
	})
}
