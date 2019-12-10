package exercises

import (
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestDay5(t *testing.T) {
	Convey("5a should output value", t, func() {
		So(Day5a(strings.NewReader("1002,6,3,6,4,6,33")), ShouldEqual, 1002)
	})
}
