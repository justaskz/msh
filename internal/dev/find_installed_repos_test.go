package dev

import (
	"testing"

	"github.com/justaskz/mashina/internal/repos"
	r "github.com/justaskz/mashina/internal/repos"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/exp/maps"
)

func TestFoo(t *testing.T) {
	Convey("", t, func() {
		t.Skip()
		configRepos := repos.BuildRepos("config")
		installedRepos := repos.BuildRepos("installed")
		repos, err := FindInstalledRepos(configRepos, installedRepos)

		reposFoo := maps.Values(repos)
		So(len(repos), ShouldEqual, len(configRepos))
		So(reposFoo[0].State, ShouldEqual, r.Installed)
		So(reposFoo[1].State, ShouldEqual, r.Missing)
		So(err, ShouldBeNil)
	})
}
