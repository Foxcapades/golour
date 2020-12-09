package colors_test

import (
	"fmt"
	"github.com/foxcapades/golour/v1/pkg/colors"

	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHsvFromCssFunc(t *testing.T) {
	Convey("NewHSVFromCSS", t, func() {
		tests := [...]struct {
			i string
			h uint16
			s float32
			l float32
			a float32
			e error
		}{
			{"hsv(0, 0%, 0%)", 0, 0, 0, 1, nil},
			{"hsva(0, 0%, 0%, 0)", 0, 0, 0, 0, nil},
			{"hsv(0, 0, 0)", 0, 0, 0, 1, nil},
			{"hsva(0, 0, 0, 0)", 0, 0, 0, 0, nil},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%s) -> {%d, %.2f, %.2f, %.2f", test.i, test.h, test.s, test.l, test.a), func() {
				val, err := colors.NewHSVFromCSS(test.i)

				So(err, ShouldEqual, test.e)
				So(val.Hue(), ShouldEqual, test.h)
				So(val.Saturation(), ShouldEqual, test.s)
				So(val.Brightness(), ShouldEqual, test.l)
				So(val.Alpha(), ShouldEqual, test.a)
			})
		}
	})
}
