package colors

import (
	"github.com/foxcapades/golour/v1/internal/util"
	"github.com/foxcapades/tally-go/v1/tally"
)

var (
	ErrRGBFuncTooShort  = NewColorError("input string is too short to be a valid rgb/rgba function")
	ErrRGBAFuncTooShort = NewColorError("input string is too short to be a valid rgba function")
)

// NewRGBFromCSS parses the given CSS function call string into a new RGB
// instance.
//
// Arguments:
//   - cssFnCall = CSS function call string, e.g. "rgb(255, 255, 255)".  This
//                 value should be trimmed of any leading or trailing whitespace
//                 characters.
//
// Returns:
//   - rgb = A newly constructed RGB instance populated with the values parsed
//           from the input cssFnCall string.
//   - err = Any error that occurred while attempting to parse the input
//           cssFnCall string.
func NewRGBFromCSS(cssFunc string) (rgb RGB, err error) {
	length := uint8(len(cssFunc))
	bytes  := util.StringToReadOnlyBytes(&cssFunc)

	if length < CSSFuncRGB.MinLength() {
		return rgb, ErrRGBFuncTooShort
	}

	pPos, ok := util.FirstParen(bytes)
	if !ok {
		return rgb, util.ErrMalformedCSSFuncString
	}
	offset := tally.UTally8(pPos)

	if CSSFuncRGBA.Matches(bytes[:pPos]) {
		if length < CSSFuncRGBA.MinLength() {
			return rgb, ErrRGBAFuncTooShort
		}

		rgb.red, rgb.green,
			rgb.blue, rgb.alpha,
			err = util.CSSByte4Parser(util.StringToReadOnlyBytes(&cssFunc), &offset)

	} else if CSSFuncRGB.Matches(bytes[:pPos]) {

		rgb.red, rgb.green,
			rgb.blue, rgb.alpha,
			err = util.CSSByte3Parser(util.StringToReadOnlyBytes(&cssFunc), &offset)

	} else {
		return rgb, util.ErrBadCssFuncName.WithValue(cssFunc, 0)
	}


	return
}
