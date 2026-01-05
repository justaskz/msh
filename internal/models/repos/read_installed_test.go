package repos

import (
	"testing"

	p "github.com/justaskz/mashina/internal/models/path"
	r "github.com/justaskz/mashina/internal/models/repo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestReadInstalled(t *testing.T) {
	Convey("", t, func() {
		repos, err := ReadInstalled(p.Fixtures.LibPath)

		So(repos.Uids(), ShouldHaveLength, 1)
		So(err, ShouldBeNil)
	})
}

func TestListFolders(t *testing.T) {
	Convey("", t, func() {
		path := p.Path("testdata/.local/lib/msh")
		folders, err := listFolders(path)
		expected := []r.Uid{"helix-editor__helix__24.07"}
		So(folders, ShouldEqual, expected)
		So(err, ShouldBeNil)
	})
}
