package wip

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReadConfigFile(t *testing.T) {
	Convey("", t, func() {
		path := "../../fixtures/config/mashina/config.yaml"
		packs, _ := ReadConfig(path)

		So(packs, ShouldEqual, Packs{})
	})
}
