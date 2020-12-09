package colors_test

import (
	"fmt"
	. "github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHSV_CSSFuncHSV(t *testing.T) {
	Convey("RGB.CSSFuncRGB", t, func() {
		tests := [...]struct {
			i HSV
			o string
		}{
			//{NewHSV(0, 0, 0), "hsv(0, 0%, 0%)"},
			{NewHSV(90, 0.2, 0.15), "hsv(90, 20%, 15%)"},
			{NewHSV(80, 0.4, 0.3), "hsv(80, 40%, 30%)"},
			{NewHSV(70, 0.6, 0.45), "hsv(70, 60%, 45%)"},
			{NewHSV(60, 0.8, 0.6), "hsv(60, 80%, 60%)"},
			{NewHSV(50, 1, 0.75), "hsv(50, 100%, 75%)"},
			{NewHSV(40, 0.75, 0.9), "hsv(40, 75%, 90%)"},
		}

		for i, test := range tests {
			Convey(fmt.Sprintf("%d. (RGB) -> %s", i, test.o), func() {
				So(test.i.CSSFuncHSV(), ShouldEqual, test.o)
			})
		}
	})
}

func TestHSV_CSSFuncHSVA(t *testing.T) {
	Convey("RGB.CSSFuncRGBA", t, func() {
		tests := [...]struct {
			i HSV
			o string
		}{
			{NewHSVA(0, 0, 0, 0), "hsva(0, 0%, 0%, 0)"},
			{NewHSVA(90, .2, .15, 0.1), "hsva(90, 20%, 15%, 0.1)"},
			{NewHSVA(80, .4, .3, 0.2), "hsva(80, 40%, 30%, 0.2)"},
			{NewHSVA(70, .6, .45, 0.4), "hsva(70, 60%, 45%, 0.4)"},
			{NewHSVA(60, .8, .6, 0.6), "hsva(60, 80%, 60%, 0.6)"},
			{NewHSVA(50, 1, .75, 0.8), "hsva(50, 100%, 75%, 0.8)"},
			{NewHSVA(40, .75, .9, 1), "hsva(40, 75%, 90%, 1)"},
		}

		for i, test := range tests {
			Convey(fmt.Sprintf("%d. (RGB) -> %s", i, test.o), func() {
				So(test.i.CSSFuncHSVA(), ShouldEqual, test.o)
			})
		}
	})
}
