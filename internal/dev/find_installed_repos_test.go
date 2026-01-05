package dev

// import (
// 	"testing"

// 	r "github.com/justaskz/mashina/internal/repos"
// 	. "github.com/smartystreets/goconvey/convey"
// 	"golang.org/x/exp/maps"
// )

// func TestFindInstalledRepos(t *testing.T) {
// 	Convey("", t, func() {
// 		configRepos := r.BuildRepos("config")
// 		installedRepos := r.BuildRepos("installed")
// 		repos, err := FindInstalledRepos(configRepos, installedRepos)

// 		reposWithState := maps.Values(repos)
// 		So(len(repos), ShouldEqual, len(configRepos))
// 		So(reposWithState[0].State, ShouldEqual, r.Installed)
// 		So(reposWithState[1].State, ShouldEqual, r.Missing)
// 		So(err, ShouldBeNil)
// 	})
// }
