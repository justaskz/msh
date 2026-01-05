package dev

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseVersion(t *testing.T) {
	Convey("", t, func() {
		version := ParseVersion("Nvim 0.9.5")
		So(version, ShouldEqual, "0.9.5")
	})

	Convey("", t, func() {
		version := ParseVersion("Nvim v0.7.2")
		So(version, ShouldEqual, "0.7.2")
	})
}
