package colors_test

import (
	. "github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRGB_ToHSL(t *testing.T) {
	Convey("RGB.ToHSL", t, func() {
		tests := [...]struct{
			i RGB
			o HSL
		} {
			{
				i: NewRGB(123, 123, 123),
				o: NewHSL(0, 0, 0.482),
			},
			{
				i: NewRGB(0, 128, 255),
				o: NewHSL(210, 1, 0.5),
			},
			{
				i: NewRGB(255, 128, 0),
				o: NewHSL(30, 1, 0.5),
			},
			{
				i: NewRGB(255, 0, 128),
				o: NewHSL(330, 1, 0.5),
			},
			{
				i: NewRGB(128, 32, 255),
				o: NewHSL(266, 1, 0.563),
			},
		}

		for _, test := range tests {
			tgt := test.i.ToHSL()

			So(tgt.Hue(), ShouldEqual, test.o.Hue())
			So(tgt.Saturation(), ShouldEqual, test.o.Saturation())
			So(tgt.Lightness(), ShouldAlmostEqual, test.o.Lightness(), 0.002)
			So(tgt.Alpha(), ShouldEqual, test.o.Alpha())
		}
	})
}