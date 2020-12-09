package colors_test

import (
	"fmt"
	. "github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRGB_CSSFuncRGB(t *testing.T) {
	Convey("RGB.CSSFuncRGB", t, func() {
		tests := [...]struct{
			i RGB
			o string
		} {
			{NewRGB(0, 0, 0), "rgb(0, 0, 0)"},
			{NewRGB(255, 0, 0), "rgb(255, 0, 0)"},
			{NewRGB(255, 255, 0), "rgb(255, 255, 0)"},
			{NewRGB(255, 255, 255), "rgb(255, 255, 255)"},
			{NewRGB(128, 255, 255), "rgb(128, 255, 255)"},
			{NewRGB(128, 128, 255), "rgb(128, 128, 255)"},
			{NewRGB(128, 128, 128), "rgb(128, 128, 128)"},
		}

		for i, test := range tests {
			Convey(fmt.Sprintf("%d. (RGB) -> %s", i, test.o), func() {
				So(test.i.CSSFuncRGB(), ShouldEqual, test.o)
			})
		}
	})
}

func TestRGB_CSSFuncRGBA(t *testing.T) {
	Convey("RGB.CSSFuncRGBA", t, func() {
		tests := [...]struct{
			i RGB
			o string
		} {
			{NewRGBA(0,   0,   0,   0),   "rgba(0, 0, 0, 0)"},
			{NewRGBA(255, 0,   0,   0),   "rgba(255, 0, 0, 0)"},
			{NewRGBA(255, 255, 0,   0),   "rgba(255, 255, 0, 0)"},
			{NewRGBA(255, 255, 255, 0),   "rgba(255, 255, 255, 0)"},
			{NewRGBA(255, 255, 255, 255), "rgba(255, 255, 255, 255)"},
			{NewRGBA(128, 255, 255, 255), "rgba(128, 255, 255, 255)"},
			{NewRGBA(128, 128, 255, 255), "rgba(128, 128, 255, 255)"},
			{NewRGBA(128, 128, 128, 255), "rgba(128, 128, 128, 255)"},
			{NewRGBA(128, 128, 128, 128), "rgba(128, 128, 128, 128)"},
		}

		for i, test := range tests {
			Convey(fmt.Sprintf("%d. (RGB) -> %s", i, test.o), func() {
				So(test.i.CSSFuncRGBA(), ShouldEqual, test.o)
			})
		}
	})
}
