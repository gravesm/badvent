package exercises

import (
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	Convey("Distance should count total distance in graph", t, func() {
		data := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L"
		g := NewGraph(strings.NewReader(data))
		i := 0
		Distance(g["COM"], 0, &i)
		So(i, ShouldEqual, 42)
	})
	Convey("Between should count distance between nodes", t, func() {
		data := "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L\nK)YOU\nI)SAN"
		g := NewGraph(strings.NewReader(data))
		So(Between(g["YOU"], g["SAN"]), ShouldEqual, 4)
	})
	Convey("6 should output distances", t, func() {
		fp, _ := os.Open("testdata/6")
		defer fp.Close()
		a, b := Day6(fp)
		So(a, ShouldEqual, 301100)
		So(b, ShouldEqual, 547)
	})
}
