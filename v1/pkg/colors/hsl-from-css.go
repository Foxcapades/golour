package colors

import (
	"github.com/foxcapades/golour/v1/internal/util"
	"github.com/foxcapades/tally-go/v1/tally"
)

var (
	ErrHslTooShort = NewColorError("hsl(a) call string too short, refusing to parse")
)

const (
	errNoLeadParen = "no opening parenthesis found"
)

type parseHsxFunc = func([]byte, *tally.UTally8) (h uint16, s, l, a float32, err error)

func NewHSLFromCSS(fnString string) (hsl HSL, err error) {
	length := uint8(len(fnString))
	bytes := []byte(fnString)
	offset := tally.UTally8(0)

	if length < CSSFuncHSL.MinLength() {
		return hsl, ErrHslTooShort.WithValue(fnString, 0)
	}

	pPos, ok := util.FirstParen(bytes)
	if !ok {
		return hsl, util.ErrMalformedCSSFuncString.WithContext(errNoLeadParen, fnString, 0)
	}

	var fn parseHsxFunc
	if CSSFuncHSLA.Matches(bytes[0:pPos]) {
		if length < CSSFuncHSLA.MinLength() {
			return hsl, ErrHslTooShort.WithValue(fnString, 0)
		}

		offset = tally.UTally8(4)
		fn = util.ParseCSSFloat4
	} else if CSSFuncHSL.Matches(bytes[0:pPos]) {
		offset = tally.UTally8(3)
		fn = util.ParseCSSFloat3
	} else {
		return hsl, util.ErrBadCssFuncName.WithValue(fnString, 0)
	}

	hsl.hue, hsl.saturation, hsl.lightness, hsl.alpha, err = fn(bytes, &offset)

	hsl.saturation /= 100
	hsl.lightness /= 100

	return
}
