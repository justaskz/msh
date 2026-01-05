package releasez

import (
	"testing"

	"github.com/justaskz/mashina/internal/models/repo"
	. "github.com/smartystreets/goconvey/convey"
)

func TestGet(t *testing.T) {
	Convey("", t, func() {
		repo := repo.Fixtures.Helix
		// system := system.New()
		release, err := Get(repo)

		So(err, ShouldBeNil)
		So(release, ShouldEqual, Release{})
	})
}
