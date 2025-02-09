package dev

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetConfigPath(t *testing.T) {
	path := GetConfigPath()
	Convey("", t, func() {
		So(path, ShouldContainSubstring, ".config/msh/config.yaml")
	})
}

func TestGetOptPath(t *testing.T) {
	Convey("", t, func() {
		path := GetOptPath()
		So(path, ShouldContainSubstring, ".msh/opt")
	})
}
