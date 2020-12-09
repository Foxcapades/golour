package colors_test

import (
	"fmt"
	. "github.com/foxcapades/golour/v1/pkg/colors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCSSFunc_Matches(t *testing.T) {
	Convey("CSSFunc.Matches", t, func() {
		tests := [...]struct{
			i string
			a CSSFunc
			o bool
		} {
			// VALID
			{"cmyk", CSSFuncCMYK, true},
			{"cmyka", CSSFuncCMYKA, true},
			{"hsb", CSSFuncHSB, true},
			{"hsba", CSSFuncHSBA, true},
			{"hsl", CSSFuncHSL, true},
			{"hsla", CSSFuncHSLA, true},
			{"hsv", CSSFuncHSV, true},
			{"hsva", CSSFuncHSVA, true},
			{"hwb", CSSFuncHWB, true},
			{"hwba", CSSFuncHWBA, true},
			{"rgb", CSSFuncRGB, true},
			{"rgba", CSSFuncRGBA, true},

			// Valid (all caps)
			{"CMYK", CSSFuncCMYK, true},
			{"CMYKA", CSSFuncCMYKA, true},
			{"HSB", CSSFuncHSB, true},
			{"HSBA", CSSFuncHSBA, true},
			{"HSL", CSSFuncHSL, true},
			{"HSLA", CSSFuncHSLA, true},
			{"HSV", CSSFuncHSV, true},
			{"HSVA", CSSFuncHSVA, true},
			{"HWB", CSSFuncHWB, true},
			{"HWBA", CSSFuncHWBA, true},
			{"RGB", CSSFuncRGB, true},
			{"RGBA", CSSFuncRGBA, true},

			// Valid (some caps)
			{"cMYK", CSSFuncCMYK, true},
			{"cMYKA", CSSFuncCMYKA, true},
			{"hSB", CSSFuncHSB, true},
			{"hSBA", CSSFuncHSBA, true},
			{"hSL", CSSFuncHSL, true},
			{"hSLA", CSSFuncHSLA, true},
			{"hSV", CSSFuncHSV, true},
			{"hSVA", CSSFuncHSVA, true},
			{"hWB", CSSFuncHWB, true},
			{"hWBA", CSSFuncHWBA, true},
			{"rGB", CSSFuncRGB, true},
			{"rGBA", CSSFuncRGBA, true},

			// INVALID (EXTRA TEXT)
			{"cmykblah", CSSFuncCMYK, false},
			{"cmykablah", CSSFuncCMYKA, false},
			{"hsbblah", CSSFuncHSB, false},
			{"hsbablah", CSSFuncHSBA, false},
			{"hslblah", CSSFuncHSL, false},
			{"hslablah", CSSFuncHSLA, false},
			{"hsvblah", CSSFuncHSV, false},
			{"hsvablah", CSSFuncHSVA, false},
			{"hwbblah", CSSFuncHWB, false},
			{"hwbablah", CSSFuncHWBA, false},
			{"rgbblah", CSSFuncRGB, false},
			{"rgbablah", CSSFuncRGBA, false},


			// INVALID (NOT CMYK)
			{"cmyka", CSSFuncCMYK, false},
			{"hsb", CSSFuncCMYK, false},
			{"hsba", CSSFuncCMYK, false},
			{"hsl", CSSFuncCMYK, false},
			{"hsla", CSSFuncCMYK, false},
			{"hsv", CSSFuncCMYK, false},
			{"hsva", CSSFuncCMYK, false},
			{"hwb", CSSFuncCMYK, false},
			{"hwba", CSSFuncCMYK, false},
			{"rgb", CSSFuncCMYK, false},
			{"rgba", CSSFuncCMYK, false},

			// INVALID (NOT CMYKA)
			{"cmyk", CSSFuncCMYKA, false},
			{"hsb", CSSFuncCMYKA, false},
			{"hsba", CSSFuncCMYKA, false},
			{"hsl", CSSFuncCMYKA, false},
			{"hsla", CSSFuncCMYKA, false},
			{"hsv", CSSFuncCMYKA, false},
			{"hsva", CSSFuncCMYKA, false},
			{"hwb", CSSFuncCMYKA, false},
			{"hwba", CSSFuncCMYKA, false},
			{"rgb", CSSFuncCMYKA, false},
			{"rgba", CSSFuncCMYKA, false},

			// INVALID (NOT HSB)
			{"cmyk", CSSFuncHSB, false},
			{"cmyka", CSSFuncHSB, false},
			{"hsba", CSSFuncHSB, false},
			{"hsl", CSSFuncHSB, false},
			{"hsla", CSSFuncHSB, false},
			//{"hsv", colors.CSSFuncHSB, false}, // Alias of hsb
			{"hsva", CSSFuncHSB, false},
			{"hwb", CSSFuncHSB, false},
			{"hwba", CSSFuncHSB, false},
			{"rgb", CSSFuncHSB, false},
			{"rgba", CSSFuncHSB, false},

			// INVALID (NOT HSBA)
			{"cmyk", CSSFuncHSBA, false},
			{"cmyka", CSSFuncHSBA, false},
			{"hsb", CSSFuncHSBA, false},
			{"hsl", CSSFuncHSBA, false},
			{"hsla", CSSFuncHSBA, false},
			{"hsv", CSSFuncHSBA, false},
			//{"hsva", colors.CSSFuncHSBA, false}, // Alias of hsba
			{"hwb", CSSFuncHSBA, false},
			{"hwba", CSSFuncHSBA, false},
			{"rgb", CSSFuncHSBA, false},
			{"rgba", CSSFuncHSBA, false},

			// INVALID (NOT HSL)
			{"cmyk", CSSFuncHSL, false},
			{"cmyka", CSSFuncHSL, false},
			{"hsb", CSSFuncHSL, false},
			{"hsba", CSSFuncHSL, false},
			{"hsla", CSSFuncHSL, false},
			{"hsv", CSSFuncHSL, false},
			{"hsva", CSSFuncHSL, false},
			{"hwb", CSSFuncHSL, false},
			{"hwba", CSSFuncHSL, false},
			{"rgb", CSSFuncHSL, false},
			{"rgba", CSSFuncHSL, false},

			// INVALID (NOT HSLA)
			{"cmyk", CSSFuncHSLA, false},
			{"cmyka", CSSFuncHSLA, false},
			{"hsb", CSSFuncHSLA, false},
			{"hsba", CSSFuncHSLA, false},
			{"hsl", CSSFuncHSLA, false},
			{"hsv", CSSFuncHSLA, false},
			{"hsva", CSSFuncHSLA, false},
			{"hwb", CSSFuncHSLA, false},
			{"hwba", CSSFuncHSLA, false},
			{"rgb", CSSFuncHSLA, false},
			{"rgba", CSSFuncHSLA, false},

			// INVALID (NOT HSV)
			{"cmyk", CSSFuncHSV, false},
			{"cmyka", CSSFuncHSV, false},
			//{"hsb", colors.CSSFuncHSV, false}, // Alias of hsv
			{"hsba", CSSFuncHSV, false},
			{"hsl", CSSFuncHSV, false},
			{"hsla", CSSFuncHSV, false},
			{"hsva", CSSFuncHSV, false},
			{"hwb", CSSFuncHSV, false},
			{"hwba", CSSFuncHSV, false},
			{"rgb", CSSFuncHSV, false},
			{"rgba", CSSFuncHSV, false},

			// INVALID (NOT HSVA)
			{"cmyk", CSSFuncHSVA, false},
			{"cmyka", CSSFuncHSVA, false},
			{"hsb", CSSFuncHSVA, false},
			//{"hsba", colors.CSSFuncHSVA, false}, // Alias of hsva
			{"hsl", CSSFuncHSVA, false},
			{"hsla", CSSFuncHSVA, false},
			{"hsv", CSSFuncHSVA, false},
			{"hwb", CSSFuncHSVA, false},
			{"hwba", CSSFuncHSVA, false},
			{"rgb", CSSFuncHSVA, false},
			{"rgba", CSSFuncHSVA, false},

			// INVALID (NOT HWB)
			{"cmyk", CSSFuncHWB, false},
			{"cmyka", CSSFuncHWB, false},
			{"hsb", CSSFuncHWB, false},
			{"hsba", CSSFuncHWB, false},
			{"hsl", CSSFuncHWB, false},
			{"hsla", CSSFuncHWB, false},
			{"hsv", CSSFuncHWB, false},
			{"hsva", CSSFuncHWB, false},
			{"hwba", CSSFuncHWB, false},
			{"rgb", CSSFuncHWB, false},
			{"rgba", CSSFuncHWB, false},

			// INVALID (NOT HWBA)
			{"cmyk", CSSFuncHWBA, false},
			{"cmyka", CSSFuncHWBA, false},
			{"hsb", CSSFuncHWBA, false},
			{"hsba", CSSFuncHWBA, false},
			{"hsl", CSSFuncHWBA, false},
			{"hsla", CSSFuncHWBA, false},
			{"hsv", CSSFuncHWBA, false},
			{"hsva", CSSFuncHWBA, false},
			{"hwb", CSSFuncHWBA, false},
			{"rgb", CSSFuncHWBA, false},
			{"rgba", CSSFuncHWBA, false},

			// INVALID (NOT RGB)
			{"cmyk", CSSFuncRGB, false},
			{"cmyka", CSSFuncRGB, false},
			{"hsb", CSSFuncRGB, false},
			{"hsba", CSSFuncRGB, false},
			{"hsl", CSSFuncRGB, false},
			{"hsla", CSSFuncRGB, false},
			{"hsv", CSSFuncRGB, false},
			{"hsva", CSSFuncRGB, false},
			{"hwb", CSSFuncRGB, false},
			{"hwba", CSSFuncRGB, false},
			{"rgba", CSSFuncRGB, false},

			// INVALID (NOT RGBA)
			{"cmyk", CSSFuncRGBA, false},
			{"cmyka", CSSFuncRGBA, false},
			{"hsb", CSSFuncRGBA, false},
			{"hsba", CSSFuncRGBA, false},
			{"hsl", CSSFuncRGBA, false},
			{"hsla", CSSFuncRGBA, false},
			{"hsv", CSSFuncRGBA, false},
			{"hsva", CSSFuncRGBA, false},
			{"hwb", CSSFuncRGBA, false},
			{"hwba", CSSFuncRGBA, false},
			{"rgb", CSSFuncRGBA, false},
		}

		for i, test := range tests {
			Convey(fmt.Sprintf("%d. %s.(\"%[2]s\") -> %t", i, test.i, test.o), func() {
				So(test.a.Matches([]byte(test.i)), ShouldEqual, test.o)
			})
		}
	})
}

func TestCSSFunc_MinLength(t *testing.T) {
	Convey("CSSFunc.MinLength", t, func() {
		tests := [...]struct{
			i CSSFunc
			o int
		} {
			{CSSFuncCMYK, len("cmyk(0,0,0,0)")},
			{CSSFuncCMYKA, len("cmyka(0,0,0,0,0)")},
			{CSSFuncHSB, len("hsb(0,0,0)")},
			{CSSFuncHSBA, len("hsba(0,0,0,0)")},
			{CSSFuncHSL, len("hsl(0,0,0)")},
			{CSSFuncHSLA, len("hsla(0,0,0,0)")},
			{CSSFuncHSV, len("hsv(0,0,0)")},
			{CSSFuncHSVA, len("hsva(0,0,0,0)")},
			{CSSFuncHWB, len("hwb(0,0,0)")},
			{CSSFuncHWBA, len("hwba(0,0,0,0)")},
			{CSSFuncRGB, len("rgb(0,0,0)")},
			{CSSFuncRGBA, len("rgba(0,0,0,0)")},
		}

		for _, test := range tests {
			Convey(fmt.Sprintf("(%s) -> %d", test.i, test.o), func() {
				So(test.i.MinLength(), ShouldEqual, test.o)
			})
		}

		Convey("invalid function type", func() {
			So(func() {CSSFunc("goodbye").MinLength()}, ShouldPanic)
			So(CSSFuncHSBA.Matches([]byte("hsbb")), ShouldBeFalse)
			So(CSSFuncHSVA.Matches([]byte("hsbb")), ShouldBeFalse)
		})
	})
}
