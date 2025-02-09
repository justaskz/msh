package dev

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParseFolderName(t *testing.T) {
	Convey("with correctly formatted string", t, func() {
		folderName := "owner__name__1.2.3"
		owner, name, version, err := ParseFolderName(folderName)

		So(*owner, ShouldEqual, "owner")
		So(*name, ShouldEqual, "name")
		So(*version, ShouldEqual, "1.2.3")
		So(err, ShouldBeNil)
	})

	Convey("with incorrectly formatted string", t, func() {
		folderName := "owner_name__1.2.3"
		owner, name, version, err := ParseFolderName(folderName)

		So(owner, ShouldBeNil)
		So(name, ShouldBeNil)
		So(version, ShouldBeNil)
		So(err.Error(), ShouldEqual, "folder 'owner_name__1.2.3' format is incorrect")
	})
}
