package exercises

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestDay2(t *testing.T) {
	Convey("Day2a should calculate value", t, func() {
		fp, _ := os.Open("testdata/2")
		defer fp.Close()
		So(Day2a(fp), ShouldEqual, 5098658)
	})

	Convey("Day2b should caculate value", t, func() {
		fp, _ := os.Open("testdata/2")
		defer fp.Close()
		So(Day2b(fp), ShouldEqual, 5064)
	})
}
