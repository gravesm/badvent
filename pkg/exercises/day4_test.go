package exercises

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestDay4(t *testing.T) {
	Convey("valid should check for valid password", t, func() {
		cases := []struct {
			in       string
			expected bool
		}{
			{"111111", true},
			{"223450", false},
			{"123789", false},
		}
		for i, test := range cases {
			Convey(fmt.Sprintf("for case %d", i), func() {
				So(valid(test.in, false), ShouldEqual, test.expected)
			})
		}
		So(valid("111111", true), ShouldBeFalse)
		So(valid("112223", true), ShouldBeTrue)
	})

	Convey("4a should count valid passwords", t, func() {
		So(Day4a(strings.NewReader("156218-652527")), ShouldEqual, 1694)
	})

	Convey("4b should count valid passwords", t, func() {
		So(Day4b(strings.NewReader("156218-652527")), ShouldEqual, 1148)
	})
}
