package colors

import (
	"github.com/foxcapades/golour/v1/internal/util"
	"github.com/foxcapades/tally-go/v1/tally"
)

var (
	ErrHSVFuncTooShort  = NewColorError("input string is too short to be a valid hsv/hsva function")
	ErrHSVAFuncTooShort = NewColorError("input string is too short to be a valid hsva function")
)

func NewHSVFromCSS(v string) (out HSV, err error) {
	length := uint8(len(v))
	bytes  := util.StringToReadOnlyBytes(&v)

	if length < CSSFuncHSV.MinLength() {
		return out, ErrHSVFuncTooShort.WithValue(v, 0)
	}

	pPos, ok := util.FirstParen(bytes)
	if !ok {
		return out, util.ErrMalformedCSSFuncString
	}
	offset := tally.UTally8(pPos)

	if CSSFuncHSVA.Matches(bytes[:pPos]) {
		if length < CSSFuncHSVA.MinLength() {
			return out, ErrHSVAFuncTooShort.WithValue(v, 0)
		}

		out.hue, out.saturation,
			out.brightness, out.alpha,
			err = util.ParseCSSFloat4(bytes, &offset)

		if err != nil {
			return
		}
	} else if CSSFuncHSV.Matches(bytes[:pPos]) {
		out.hue, out.saturation,
			out.brightness, out.alpha,
			err = util.ParseCSSFloat3(bytes, &offset)

		if err != nil {
			return
		}
	}

	out.saturation /= 100
	out.brightness /= 100

	return
}
