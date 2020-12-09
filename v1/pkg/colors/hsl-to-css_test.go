package colors_test

import (
	"fmt"
	. "github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestHSL_CSSFuncHSL(t *testing.T) {
	Convey("RGB.CSSFuncRGB", t, func() {
		tests := [...]struct {
			i HSL
			o string
		}{
			{NewHSL(0, 0, 0), "hsl(0, 0%, 0%)"},
			{NewHSL(90, 0.2, 0.15), "hsl(90, 20%, 15%)"},
			{NewHSL(80, 0.4, 0.3), "hsl(80, 40%, 30%)"},
			{NewHSL(70, 0.6, 0.45), "hsl(70, 60%, 45%)"},
			{NewHSL(60, 0.8, 0.6), "hsl(60, 80%, 60%)"},
			{NewHSL(50, 1, 0.75), "hsl(50, 100%, 75%)"},
			{NewHSL(40, 0.75, 0.9), "hsl(40, 75%, 90%)"},
		}

		for i, test := range tests {
			Convey(fmt.Sprintf("%d. (RGB) -> %s", i, test.o), func() {
				So(test.i.CSSFuncHSL(), ShouldEqual, test.o)
			})
		}
	})
}

func TestHSL_CSSFuncHSLA(t *testing.T) {
	Convey("RGB.CSSFuncRGBA", t, func() {
		tests := [...]struct {
			i HSL
			o string
		}{
			{NewHSLA(0, 0, 0, 0), "hsla(0, 0%, 0%, 0)"},
			{NewHSLA(90, .2, .15, 0.1), "hsla(90, 20%, 15%, 0.1)"},
			{NewHSLA(80, .4, .3, 0.2), "hsla(80, 40%, 30%, 0.2)"},
			{NewHSLA(70, .6, .45, 0.4), "hsla(70, 60%, 45%, 0.4)"},
			{NewHSLA(60, .8, .6, 0.6), "hsla(60, 80%, 60%, 0.6)"},
			{NewHSLA(50, 1, .75, 0.8), "hsla(50, 100%, 75%, 0.8)"},
			{NewHSLA(40, .75, .9, 1), "hsla(40, 75%, 90%, 1)"},
		}

		for i, test := range tests {
			Convey(fmt.Sprintf("%d. (RGB) -> %s", i, test.o), func() {
				So(test.i.CSSFuncHSLA(), ShouldEqual, test.o)
			})
		}
	})
}
