package exercises

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"
)

func TestDay7(t *testing.T) {
	Convey("Day 7 should calculate values", t, func() {
		fp, _ := os.Open("testdata/7")
		defer fp.Close()
		a, b := Day7(fp)
		So(a, ShouldEqual, 366376)
		So(b, ShouldEqual, 21596786)
	})
}
