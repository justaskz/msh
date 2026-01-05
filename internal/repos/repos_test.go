package repos

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateRepos(t *testing.T) {
	Convey("", t, func() {
		repos := CreateRepos()

		So(repos, ShouldEqual, Repos{})
	})
}

func TestReposAdd(t *testing.T) {
	Convey("", t, func() {
		repo1, err1 := CreateRepo("owner", "repo", "1.2.3")
		repo2, err2 := CreateRepo("owner", "repo", "2.2.3")
		repos := CreateRepos()
		repos = repos.Add(repo1)
		repos = repos.Add(repo2)

		So(err1, ShouldBeNil)
		So(err2, ShouldBeNil)
		So(repos, ShouldHaveLength, 2)
		So(repos, ShouldContainKey, repo1.Uid())
		So(repos, ShouldContainKey, repo2.Uid())
	})
}

func TestReposGet(t *testing.T) {
	Convey("", t, func() {
		repo := BuildRepo("dummy")
		repos := CreateRepos().Add(repo)
		result, err := repos.Get(repo.Uid())
		So(result, ShouldEqual, repo)
		So(err, ShouldBeNil)
	})
}

func TestRepoNameParse(t *testing.T) {
	Convey("with correctly formatted repo name", t, func() {
		repoName := RepoName("owner/repo")
		owner, name, err := repoName.Parse()

		So(owner, ShouldEqual, "owner")
		So(name, ShouldEqual, "repo")
		So(err, ShouldBeNil)
	})

	Convey("with incorrectly formatted repo name", t, func() {
		repoName := RepoName("owner/repo/version")
		owner, name, err := repoName.Parse()
		So(err.Error(), ShouldEqual, "Repository name format is incorrect: 'owner/repo/version'")
		So(owner, ShouldEqual, "")
		So(name, ShouldEqual, "")
	})
}
