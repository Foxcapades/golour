package colors_test

import (
	. "github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRGB_HexRGB(t *testing.T) {
	Convey("RGB.HexRGB", t, func() {
		tests := [...]struct{
			i RGB
			o string
		} {
			{NewRGB(0xFF, 0xEE, 0xDD), "#FFEEDD"},
			{NewRGB(0x00, 0x55, 0x44), "#005544"},
			{NewRGBA(0x12, 0x34, 0x56, 0x78), "#123456"},
			{NewRGBA(0x98, 0x76, 0x54, 0x32), "#987654"},
		}

		for _, test := range tests {
			So(test.i.HexRGB(), ShouldEqual, test.o)
		}
	})
}

func TestRGB_HexRGBA(t *testing.T) {
	Convey("RGB.HexRGBA", t, func() {
		tests := [...]struct{
			i RGB
			o string
		} {
			{NewRGB(0xFF, 0xEE, 0xDD), "#FFEEDDFF"},
			{NewRGB(0x00, 0x55, 0x44), "#005544FF"},
			{NewRGBA(0x12, 0x34, 0x56, 0x78), "#12345678"},
			{NewRGBA(0x98, 0x76, 0x54, 0x32), "#98765432"},
		}

		for _, test := range tests {
			So(test.i.HexRGBA(), ShouldEqual, test.o)
		}
	})
}
