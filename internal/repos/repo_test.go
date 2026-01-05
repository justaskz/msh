package repos

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateRepo(t *testing.T) {
	Convey("when values are present", t, func() {
		repo, err := CreateRepo("owner", "repo", "1.2.3")

		So(err, ShouldBeNil)
		So(repo.Owner, ShouldEqual, "owner")
		So(repo.Name, ShouldEqual, "repo")
		So(repo.Version, ShouldEqual, "1.2.3")
		So(repo.State, ShouldEqual, Missing)
		So(repo.Url, ShouldEqual, Url(""))
		So(repo.State.String(), ShouldEqual, "missing")

		So(repo.Uid(), ShouldEqual, Uid("owner__repo__1.2.3"))
	})

	Convey("when values are not present", t, func() {
		repo, err := CreateRepo("", "", "")

		So(repo, ShouldEqual, Repo{})
		So(err.Error(), ShouldEqual, "value is undefined")
	})
}

func TestRepoUpdateState(t *testing.T) {
	Convey("", t, func() {
		repo := BuildRepo("dummy")
		So(repo.State, ShouldEqual, Missing)
		repo = repo.UpdateState(Installed)
		So(repo.State, ShouldEqual, Installed)
	})
}
