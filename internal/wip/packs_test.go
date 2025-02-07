package wip

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreatePackFromFolderName(t *testing.T) {
	Convey("with correctly formatted string", t, func() {
		folderName := "helix-editor__helix__24.07"
		result, err := CreatePackFromFolderName(folderName)
		pack := Pack{Owner: "helix-editor", Name: "helix", Version: "24.07"}

		So(result, ShouldEqual, pack)
		So(err, ShouldBeNil)
	})

	Convey("with incorrectly formatted string", t, func() {
		folderName := "helix-editor_helix__24.07"
		pack, err := CreatePackFromFolderName(folderName)

		So(pack, ShouldEqual, Pack{})
		So(err.Error(), ShouldContainSubstring, "folder name format is incorrect")
	})

	Convey("with empty string", t, func() {
		folderName := ""
		pack, err := CreatePackFromFolderName(folderName)

		So(pack, ShouldEqual, Pack{})
		So(err.Error(), ShouldContainSubstring, "folder name format is incorrect")
	})
}

func TestCreatePack(t *testing.T) {
	Convey("", t, func() {
		pack := CreatePack("foo", "bar", "1.2.3")
		So(pack.Uid, ShouldEqual, Uid("uid_"))
		So(pack.Owner, ShouldEqual, "foo")
		So(pack.Name, ShouldEqual, "bar")
		So(pack.Version, ShouldEqual, "1.2.3")
	})
}

func TestCreatePacks(t *testing.T) {
	Convey("", t, func() {
		packs := CreatePacks()

		So(packs, ShouldEqual, Packs{})
	})
}

func TestPacksAdd(t *testing.T) {
	Convey("", t, func() {
		pack := Pack{}
		packs := CreatePacks()
		packs = packs.Add(pack)

		So(packs, ShouldContainKey, pack.Uid)
	})
}
