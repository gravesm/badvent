package exercises

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestDay2(t *testing.T) {
	Convey("Day 2 should calculate values", t, func() {
		fp, _ := os.Open("testdata/2")
		defer fp.Close()
		a, b := Day2(fp)
		So(a, ShouldEqual, 5098658)
		So(b, ShouldEqual, 5064)
	})
}
