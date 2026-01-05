package repos

import (
	"testing"

	p "github.com/justaskz/mashina/internal/models/path"
	r "github.com/justaskz/mashina/internal/models/repo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRead(t *testing.T) {
	Convey("", t, func() {
		repos, err := Read(p.Fixtures.ConfigPath, p.Fixtures.LibPath)

		So(err, ShouldBeNil)
		So(repos.Uids(), ShouldHaveLength, 2)

		helix := r.Fixtures.Helix
		starship := r.Fixtures.Starship
		_helix, _ := repos.Get(helix.Uid())
		_starship, _ := repos.Get(starship.Uid())

		So(_helix.State, ShouldEqual, r.States.Installed)
		So(_starship.State, ShouldEqual, r.States.Missing)
	})
}
