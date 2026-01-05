package repos

import (
	"testing"

	p "github.com/justaskz/mashina/internal/models/path"
	r "github.com/justaskz/mashina/internal/models/repo"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/exp/maps"
)

func TestReadConfigFile(t *testing.T) {
	Convey("", t, func() {
		repos, err := ReadConfig(p.Fixtures.ConfigPath)
		repo := maps.Values(repos)[0]

		So(err, ShouldBeNil)
		So(repos, ShouldHaveLength, 2)
		So(repo.Owner, ShouldEqual, r.Owner("helix-editor"))
		So(repo.Name, ShouldEqual, r.Name("helix"))
	})
}
