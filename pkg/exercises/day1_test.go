package exercises

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"strings"
	"testing"
)

func TestDay1(t *testing.T) {
	Convey("Day1a should calculate fuel", t, func() {
		data, _ := ioutil.ReadFile("testdata/1")
		cases := []struct {
			in       string
			expected int
		}{
			{"1969", 654},
			{"100756", 33583},
			{string(data), 3488702},
		}
		for i, test := range cases {
			Convey(fmt.Sprintf("for case %d", i), func() {
				So(Day1a(strings.NewReader(test.in)), ShouldEqual, test.expected)
			})
		}
	})

	Convey("Day1b should calculate fuel", t, func() {
		data, _ := ioutil.ReadFile("testdata/1")
		cases := []struct {
			in       string
			expected int
		}{
			{"1969", 966},
			{"100756", 50346},
			{string(data), 5230169},
		}
		for i, test := range cases {
			Convey(fmt.Sprintf("for case %d", i), func() {
				So(Day1b(strings.NewReader(test.in)), ShouldEqual, test.expected)
			})
		}
	})
}
