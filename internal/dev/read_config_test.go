package dev

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadConfigFile(t *testing.T) {
	Convey("", t, func() {
		path := "testdata/config.yaml"
		repos, err := ReadConfig(path)

		So(err, ShouldBeNil)
		So(len(repos), ShouldEqual, 2)
	})
}
