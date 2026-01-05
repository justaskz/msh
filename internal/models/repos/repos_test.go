package repos

import (
	"testing"

	r "github.com/justaskz/mashina/internal/models/repo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("", t, func() {
		repos := New()

		So(repos, ShouldEqual, Repos{})
		So(repos.Uids(), ShouldHaveLength, 0)
	})
}

func TestReposAdd(t *testing.T) {
	Convey("", t, func() {
		repo1, err1 := r.New("owner", "repo", "1.2.3")
		repo2, err2 := r.New("owner", "repo", "2.2.3")
		repos := New()
		repos = repos.Add(repo1)
		repos = repos.Add(repo2)

		So(err1, ShouldBeNil)
		So(err2, ShouldBeNil)
		So(repos.Uids(), ShouldHaveLength, 2)
		So(repos, ShouldContainKey, repo1.Uid())
		So(repos, ShouldContainKey, repo2.Uid())
	})
}

func TestReposGet(t *testing.T) {
	Convey("", t, func() {
		repo := r.Fixtures.Helix
		repos := New().Add(repo)
		result, err := repos.Get(repo.Uid())
		So(err, ShouldBeNil)
		So(result, ShouldEqual, repo)
	})
}

func TestGetAll(t *testing.T) {
	Convey("", t, func() {
		repo1 := r.Fixtures.Helix
		repo2 := r.Fixtures.Starship
		repos := New()
		repos = repos.Add(repo1)
		repos = repos.Add(repo2)
		all := repos.GetAll()

		So(all, ShouldEqual, []r.Repo{repo1, repo2})
	})
}
