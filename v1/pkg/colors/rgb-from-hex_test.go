package colors_test

import (
	"fmt"
	"github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestNewRGBFromHex(t *testing.T) {
	Convey("NewRGBFromHex", t, func() {
		tests := [...]struct{
			i string
			o colors.RGB
			e error
		} {
			{i: "DEF", o: colors.NewRGBA(0xDD, 0xEE, 0xFF, 0xFF)},
			{i: "#ABC", o: colors.NewRGBA(0xAA, 0xBB, 0xCC, 0xFF)},
			{i: "ABCD", o: colors.NewRGBA(0xAA, 0xBB, 0xCC, 0xDD)},
			{i: "#ABCD", o: colors.NewRGBA(0xAA, 0xBB, 0xCC, 0xDD)},
			{i: "012345", o: colors.NewRGBA(0x01, 0x23, 0x45, 0xFF)},
			{i: "#543210", o: colors.NewRGBA(0x54, 0x32, 0x10, 0xFF)},
			{i: "14785236", o: colors.NewRGBA(0x14, 0x78, 0x52, 0x36)},
			{i: "#78945612", o: colors.NewRGBA(0x78, 0x94, 0x56, 0x12)},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%s) -> (RGB, error)", test.i), func() {
				val, err := colors.NewRGBFromHex(test.i)

				So(err, ShouldEqual, test.e)
				So(val.Red(), ShouldEqual, test.o.Red())
				So(val.Green(), ShouldEqual, test.o.Green())
				So(val.Blue(), ShouldEqual, test.o.Blue())
				So(val.Alpha(), ShouldEqual, test.o.Alpha())
			})
		}
	})
}
