package exercises

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestDay5(t *testing.T) {
	Convey("Day 5 should calculate values", t, func() {
		fp, _ := os.Open("testdata/5")
		defer fp.Close()
		a, b := Day5(fp)
		So(a, ShouldEqual, 13933662)
		So(b, ShouldEqual, 2369720)
	})
}
