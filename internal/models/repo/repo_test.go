package repo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("when values are present", t, func() {
		repo, err := New(
			Owner("owner"),
			Name("repo"),
			Version("1.2.3"),
		)

		So(err, ShouldBeNil)
		So(repo.Owner, ShouldEqual, Owner("owner"))
		So(repo.Name, ShouldEqual, Name("repo"))
		So(repo.Version, ShouldEqual, Version("1.2.3"))
		So(repo.Uid(), ShouldEqual, Uid("owner__repo__1.2.3"))
		So(repo.State, ShouldEqual, States.Missing)

		// So(repo.Url, ShouldEqual, Url(""))
	})

	Convey("when values are not present", t, func() {
		repo, err := New(
			Owner(""),
			Name("repo"),
			Version("1.2.3"),
		)

		So(repo, ShouldEqual, Repo{})
		So(err.Error(), ShouldEqual, "Owner is undefined")
	})
}

func TestParseUid(t *testing.T) {
	Convey("with correctly formatted string", t, func() {
		uid := Uid("owner__name__1.2.3")
		repo, err := Parse(uid)

		So(repo.Owner, ShouldEqual, Owner("owner"))
		So(repo.Name, ShouldEqual, Name("name"))
		So(repo.Version, ShouldEqual, Version("1.2.3"))
		So(err, ShouldBeNil)
	})

	Convey("with incorrectly formatted string", t, func() {
		uid := Uid("owner_name__1.2.3")
		repo, err := Parse(uid)

		So(repo.Owner, ShouldBeEmpty)
		So(repo.Name, ShouldBeEmpty)
		So(repo.Version, ShouldBeEmpty)
		So(err.Error(), ShouldEqual, "folder 'owner_name__1.2.3' format is incorrect")
	})
}

func TestSetState(t *testing.T) {
	Convey("", t, func() {
		repo := Fixtures.Helix
		So(repo.State, ShouldEqual, States.Missing)
		repo = repo.SetState(States.Installed)
		So(repo.State, ShouldEqual, States.Installed)
	})
}
