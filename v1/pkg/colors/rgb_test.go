package colors_test

import (
	"testing"

	. "github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNewRGB(t *testing.T) {
	Convey("NewRGB", t, func() {
		test := NewRGB(0xFE, 0xDC, 0xBA)
		So(test.Red(), ShouldEqual, 0xFE)
		So(test.Green(), ShouldEqual, 0xDC)
		So(test.Blue(), ShouldEqual, 0xBA)
		So(test.Alpha(), ShouldEqual, 0xFF)
	})
}

func TestNewRGBA(t *testing.T) {
	Convey("NewRGBA", t, func() {
		test := NewRGBA(0xFE, 0xDC, 0xBA, 0x98)
		So(test.Red(), ShouldEqual, 0xFE)
		So(test.Green(), ShouldEqual, 0xDC)
		So(test.Blue(), ShouldEqual, 0xBA)
		So(test.Alpha(), ShouldEqual, 0x98)
	})
}

func TestRGBA_Red(t *testing.T) {
	Convey("RGB.Red()", t, func() {
		var test RGB
		So(test.Red(), ShouldEqual, 0)
		test.SetRed(255)
		So(test.Red(), ShouldEqual, 255)
	})
}

func TestRGBA_Green(t *testing.T) {
	Convey("RGB.Green()", t, func() {
		var test RGB
		So(test.Green(), ShouldEqual, 0)
		test.SetGreen(255)
		So(test.Green(), ShouldEqual, 255)
	})
}

func TestRGBA_Blue(t *testing.T) {
	Convey("RGB.Blue()", t, func() {
		var test RGB
		So(test.Blue(), ShouldEqual, 0)
		test.SetBlue(255)
		So(test.Blue(), ShouldEqual, 255)
	})
}

func TestRGBA_Alpha(t *testing.T) {
	Convey("RGB.Alpha()", t, func() {
		var test RGB
		So(test.Alpha(), ShouldEqual, 0)
		test.SetAlpha(1)
		So(test.Alpha(), ShouldEqual, 1)
	})
}

func TestRGB_RedF32(t *testing.T) {
	Convey("RGB.RedF32()", t, func() {
		var test RGB

		So(test.RedF32(), ShouldEqual, 0)

		test.SetRedF32(1)
		So(test.Red(), ShouldEqual, 255)
		So(test.RedF32(), ShouldEqual, 1)

		test.SetRedF32(2)
		So(test.Red(), ShouldEqual, 255)
		So(test.RedF32(), ShouldEqual, 1)

		test.SetRedF32(0.5)
		So(test.Red(), ShouldEqual, 127)
		So(test.RedF32(), ShouldEqual, 0.5)
	})
}

func TestRGB_GreenF32(t *testing.T) {
	Convey("RGB.GreenF32()", t, func() {
		var test RGB

		So(test.GreenF32(), ShouldEqual, 0)

		test.SetGreenF32(1)
		So(test.Green(), ShouldEqual, 255)
		So(test.GreenF32(), ShouldEqual, 1)

		test.SetGreenF32(2)
		So(test.Green(), ShouldEqual, 255)
		So(test.GreenF32(), ShouldEqual, 1)

		test.SetGreenF32(0.5)
		So(test.Green(), ShouldEqual, 127)
		So(test.GreenF32(), ShouldEqual, 0.5)
	})
}

func TestRGB_BlueF32(t *testing.T) {
	Convey("RGB.BlueF32()", t, func() {
		var test RGB

		So(test.BlueF32(), ShouldEqual, 0)

		test.SetBlueF32(1)
		So(test.Blue(), ShouldEqual, 255)
		So(test.BlueF32(), ShouldEqual, 1)

		test.SetBlueF32(2)
		So(test.Blue(), ShouldEqual, 255)
		So(test.BlueF32(), ShouldEqual, 1)

		test.SetBlueF32(0.5)
		So(test.Blue(), ShouldEqual, 127)
		So(test.BlueF32(), ShouldEqual, 0.5)
	})
}

func TestRGB_AlphaF32(t *testing.T) {
	Convey("RGB.AlphaF32()", t, func() {
		var test RGB

		So(test.AlphaF32(), ShouldEqual, 0)

		test.SetAlphaF32(1)
		So(test.Alpha(), ShouldEqual, 255)
		So(test.AlphaF32(), ShouldEqual, 1)

		test.SetAlphaF32(2)
		So(test.Alpha(), ShouldEqual, 255)
		So(test.AlphaF32(), ShouldEqual, 1)

		test.SetAlphaF32(0.5)
		So(test.Alpha(), ShouldEqual, 127)
		So(test.AlphaF32(), ShouldEqual, 0.5)
	})
}
