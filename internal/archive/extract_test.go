package archive

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

var testFolder = "/tmp/test"

func TestExtract(t *testing.T) {
	Convey("when archive does not have root folder", t, func() {
		path := "testdata/without_folder.tar.gz"
		destination := testFolder + "/msh/test_repo"
		bin_path := destination + "/bin"
		os.RemoveAll(testFolder)

		err := Extract(path, destination)

		So(err, ShouldBeNil)
		So(fileExists(bin_path), ShouldBeTrue)
	})

	Convey("when archive has root folder", t, func() {
		path := "testdata/with_folder.tar.gz"
		destination := testFolder + "/msh/test_repo"
		bin_path := destination + "/bin"
		os.RemoveAll(testFolder)

		err := Extract(path, destination)

		So(err, ShouldBeNil)
		So(fileExists(bin_path), ShouldBeTrue)
		os.RemoveAll(testFolder)
	})
}
