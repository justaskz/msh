package repos

import (
	"testing"

	r "github.com/justaskz/mashina/internal/models/repo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestFindInstalledRepos(t *testing.T) {
	Convey("", t, func() {
		repos, err := UpdateState(Fixtures.ConfigRepos, Fixtures.InstalledRepos)

		So(err, ShouldBeNil)
		So(repos, ShouldHaveLength, 2)

		all := repos.GetAll()
		So(all[0].State, ShouldEqual, r.States.Installed)
		So(all[1].State, ShouldEqual, r.States.Missing)
	})
}
