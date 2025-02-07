package repos

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestParsePackageName(t *testing.T) {
	Convey("with correctly formatted repo name", t, func() {
		repoName := "owner/repo"
		repo, err := ParseRepoName(repoName)
		So(repo, ShouldEqual, Repo{Owner: "owner", Name: "repo"})
		So(err, ShouldBeNil)
	})

	Convey("with incorrectly formatted repo name", t, func() {
		repoName := "owner/repo/version"
		repo, err := ParseRepoName(repoName)
		So(repo, ShouldEqual, Repo{})
		So(err.Error(), ShouldEqual, "Repository name format is incorrect: 'owner/repo/version'")
	})
}
