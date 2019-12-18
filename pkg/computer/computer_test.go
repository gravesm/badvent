package computer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestComputer(t *testing.T) {
	Convey("Computer", t, func() {

		Convey("should find value", func() {
			p := NewTask()
			p.Load([]int{1002, 4, 3, 4, 33})
			p.Run()
			So(p.tape[4], ShouldEqual, 99)
		})

		Convey("should support negative codes", func() {
			p := NewTask()
			p.Load([]int{1101, 100, -1, 4, 0})
			p.Run()
			So(p.tape[4], ShouldEqual, 99)
		})

		Convey("should read input", func() {
			t := NewTask()
			t.Load([]int{3, 3, 99, 0})
			t.In <- 8
			t.Run()
			So(t.tape[3], ShouldEqual, 8)
		})

		Convey("should write output", func() {
			t := NewTask()
			t.Load([]int{3, 5, 4, 5, 99, 0})
			t.In <- 8
			t.Run()
			So(<-t.Out, ShouldEqual, 8)
		})

		Convey("should support jumps", func() {
			t := NewTask()
			t.Load([]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9})
			t.In <- 0
			t.Run()
			So(<-t.Out, ShouldEqual, 0)
		})
	})
}
