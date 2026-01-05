package dev

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEquality(t *testing.T) {
	Convey("", t, func() {
		system1 := SystemInfo{Arch: Arm64, Platform: Darwin}
		system2 := SystemInfo{Arch: Arm64, Platform: Darwin}
		So(system1, ShouldEqual, system2)
	})
}

func TestGetSystemInfoFromUrl(t *testing.T) {
	Convey("", t, func() {
		url := "https://github.com/junegunn/fzf/releases/download/0.41.0/fzf-0.41.0-darwin_arm64.zip"
		systemInfo, err := GetSystemInfoFromUrl(url)
		So(err, ShouldBeNil)
		So(systemInfo.Arch, ShouldEqual, Arm64)
		So(systemInfo.Platform, ShouldEqual, Darwin)
	})

	Convey("", t, func() {
		url := "https://github.com/starship/starship/releases/download/v1.17.1/starship-x86_64-unknown-LINUX-gnu.tar.gz"
		systemInfo, err := GetSystemInfoFromUrl(url)
		So(err, ShouldBeNil)
		So(systemInfo.Arch, ShouldEqual, Amd64)
		So(systemInfo.Platform, ShouldEqual, Linux)
	})

	Convey("", t, func() {
		url := "https://github.com/starship/starship/releases/download/v1.17.1/starship-arch-unknown-platform-gnu.tar.gz"
		_, err := GetSystemInfoFromUrl(url)
		So(err.Error(), ShouldEqual, fmt.Sprintf("failed to parse architecture from url %s", url))
	})

	// Convey("", t, func() {
	// 	url := "https://github.com/neovim/neovim/releases/download/v0.9.5/nvim-macos.tar.gz"
	// 	systemInfo, err := GetSystemInfoFromUrl(url)
	// 	So(err, ShouldBeNil)
	// 	So(systemInfo.Arch, ShouldEqual, Amd64)
	// 	So(systemInfo.Platform, ShouldEqual, Darwin)
	// })
}
