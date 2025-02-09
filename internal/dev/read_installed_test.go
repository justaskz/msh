package dev

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadInstalled(t *testing.T) {
	Convey("", t, func() {
		path := "testdata/opt"
		repos, err := ReadInstalled(path)

		So(len(repos), ShouldEqual, 1)
		So(err, ShouldBeNil)
	})
}

func TestListFolders(t *testing.T) {
	Convey("", t, func() {
		path := "testdata/opt"
		folders, err := listFolders(path)
		expected := []string{"helix-editor__helix__24.07"}
		So(folders, ShouldEqual, expected)
		So(err, ShouldBeNil)
	})
}
