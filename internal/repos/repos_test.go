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
		So(len(repos), ShouldEqual, 2)
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
