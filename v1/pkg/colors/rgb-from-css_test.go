package colors_test

import (
	"fmt"
	"github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewRGBFromCSS(t *testing.T) {
	Convey("NewRGBFromCSS", t, func() {
		tests := [...]struct{
			i string
			o colors.RGB
			e error
		} {
			{"rgb(128, 128, 128)", colors.NewRGBA(128, 128, 128, 255), nil},
			{"rgba(128, 128, 128, 128)", colors.NewRGBA(128, 128, 128, 128), nil},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%s) -> (RGB, error)", test.i), func() {
				rgb, err := colors.NewRGBFromCSS(test.i)

				So(err, ShouldEqual, test.e)
				So(rgb.Red(), ShouldEqual, test.o.Red())
				So(rgb.Green(), ShouldEqual, test.o.Green())
				So(rgb.Blue(), ShouldEqual, test.o.Blue())
				So(rgb.Alpha(), ShouldEqual, test.o.Alpha())
			})
		}
	})
}
