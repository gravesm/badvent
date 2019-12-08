package exercises

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"strings"
	"testing"
)

func TestDay3(t *testing.T) {
	Convey("Line", t, func() {
		cases := []struct {
			description string
			s1          Line
			s2          Line
			intersects  bool
			point       Point
		}{
			{
				"both horizontal",
				Line{Point{0, 0}, Point{5, 0}},
				Line{Point{1, 1}, Point{5, 1}},
				false,
				Point{0, 0},
			},
			{
				"both vertical",
				Line{Point{0, 0}, Point{0, 5}},
				Line{Point{1, 1}, Point{1, 5}},
				false,
				Point{0, 0},
			},
			{
				"non-intersecting",
				Line{Point{3, 5}, Point{6, 5}},
				Line{Point{7, 2}, Point{7, 10}},
				false,
				Point{0, 0},
			},
			{
				"intersect horizontal, vertical",
				Line{Point{1, 1}, Point{5, 1}},
				Line{Point{4, 0}, Point{4, 3}},
				true,
				Point{4, 1},
			},
			{
				"intersect vertical, horizontal",
				Line{Point{4, 0}, Point{4, 3}},
				Line{Point{1, 1}, Point{5, 1}},
				true,
				Point{4, 1},
			},
		}
		Convey("Intersect", func() {
			for _, test := range cases {
				Convey("should check "+test.description, func() {
					So(test.s1.Intersects(test.s2), ShouldEqual, test.intersects)
				})
			}
		})

		Convey("Where", func() {
			for _, test := range cases[3:] {
				Convey("should return point for "+test.description, func() {
					So(test.s1.Where(test.s2), ShouldResemble, test.point)
				})
			}
		})

		Convey("Length should calculate length", func() {
			cases := []struct {
				line     Line
				expected int
			}{
				{
					Line{Point{0, 0}, Point{5, 0}},
					5,
				},
				{
					Line{Point{5, 0}, Point{5, 10}},
					10,
				},
			}
			for i, test := range cases {
				Convey(fmt.Sprintf("for case %d", i), func() {
					So(test.line.Length(), ShouldEqual, test.expected)
				})
			}
		})

		Convey("Contains should determine if point is on line", func() {
			cases := []struct {
				line     Line
				point    Point
				expected bool
			}{
				{
					Line{Point{0, 0}, Point{5, 0}},
					Point{3, 0},
					true,
				},
				{
					Line{Point{5, 0}, Point{5, 10}},
					Point{5, 9},
					true,
				},
				{
					Line{Point{0, 0}, Point{0, 10}},
					Point{1, 1},
					false,
				},
			}
			for i, test := range cases {
				Convey(fmt.Sprintf("for case %d", i), func() {
					So(test.line.Contains(test.point), ShouldEqual, test.expected)
				})
			}
		})
	})

	Convey("Path", t, func() {
		p := Path{
			Segments: []Line{
				Line{Point{0, 0}, Point{0, 4}},
				Line{Point{0, 4}, Point{5, 4}},
				Line{Point{5, 4}, Point{5, 2}},
				Line{Point{5, 2}, Point{4, 2}},
				Line{Point{4, 2}, Point{4, 10}},
			},
		}
		Convey("Boston should find shortest distance to point", func() {
			So(p.Boston(Point{4, 4}), ShouldEqual, 8)
		})

		Convey("NewPath should create new path", func() {
			So(NewPath("U4,R5,D2,L1,U8"), ShouldResemble, p)
		})
	})

	Convey("Segments should convert path to lines", t, func() {
		path := []Point{
			Point{0, 0}, Point{5, 0}, Point{5, -4}, Point{4, -4}, Point{4, 6},
		}
		expected := []Line{
			Line{Point{0, 0}, Point{5, 0}},
			Line{Point{5, 0}, Point{5, -4}},
			Line{Point{5, -4}, Point{4, -4}},
			Line{Point{4, -4}, Point{4, 6}},
		}
		So(segments(path), ShouldResemble, expected)
	})

	Convey("Manhattan should return distance", t, func() {
		p1 := Point{0, 0}
		p2 := Point{5, 7}
		So(p1.Manhattan(p2), ShouldEqual, 12)
		So(p1.Manhattan(p2), ShouldEqual, p2.Manhattan(p1))
	})

	Convey("3a should find shortest distance", t, func() {
		data, _ := ioutil.ReadFile("testdata/3")
		cases := []struct {
			in       string
			expected int
		}{
			{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
				159,
			},
			{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
				135,
			},
			{
				string(data),
				3247,
			},
		}
		for i, test := range cases {
			Convey(fmt.Sprintf("for case %d", i), func() {
				So(Day3a(strings.NewReader(test.in)), ShouldEqual, test.expected)
			})
		}
	})

	Convey("3b should find shortest path", t, func() {
		data, _ := ioutil.ReadFile("testdata/3")
		cases := []struct {
			in       string
			expected int
		}{
			{
				"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
				610,
			},
			{
				"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
				410,
			},
			{
				string(data),
				48054,
			},
		}
		for i, test := range cases {
			Convey(fmt.Sprintf("for case %d", i), func() {
				So(Day3b(strings.NewReader(test.in)), ShouldEqual, test.expected)
			})
		}
	})
}
