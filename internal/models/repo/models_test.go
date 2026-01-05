package repo

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRepoNameParse(t *testing.T) {
	Convey("with correctly formatted repo name", t, func() {
		repoName := RepoName("owner/repo")
		owner, name, err := repoName.Parse()

		So(err, ShouldBeNil)
		So(owner, ShouldEqual, Owner("owner"))
		So(name, ShouldEqual, Name("repo"))
	})

	Convey("with incorrectly formatted repo name", t, func() {
		repoName := RepoName("owner/repo/version")
		owner, name, err := repoName.Parse()

		So(err.Error(), ShouldEqual, "Repository name format is incorrect: 'owner/repo/version'")
		So(owner, ShouldEqual, Owner(""))
		So(name, ShouldEqual, Name(""))
	})
}
