package system

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("", t, func() {
		system := New()

		So(system.Arch, ShouldEqual, Archs.Arm64)
		So(system.Platform, ShouldEqual, Platforms.Darwin)
	})
}

func TestEquality(t *testing.T) {
	Convey("", t, func() {
		system1 := System{Arch: Archs.Arm64, Platform: Platforms.Darwin}
		system2 := System{Arch: Archs.Arm64, Platform: Platforms.Darwin}
		So(system1, ShouldEqual, system2)
	})
}
