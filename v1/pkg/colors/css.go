package colors

import "github.com/foxcapades/golour/v1/internal/util"

type CSSFunc string

// CSS color function types.
//
// Note: Not all of the types listed below are part of the CSS3 specification.
// Several of them are suggested for CSS4, but should not be used in standard
// HTML/CSS.
const (
	// CSS HSL color function name.
	CSSFuncHSL  CSSFunc = "hsl"

	// CSS HSL with alpha color function name.
	CSSFuncHSLA CSSFunc = "hsla"

	// CSS RGB color function name.
	CSSFuncRGB  CSSFunc = "rgb"

	// CSS RGB color with alpha function name.
	CSSFuncRGBA CSSFunc = "rgba"

	// CSS HSV color function name.
	// WARNING: This is not a part of the CSS3 spec, use with caution.
	CSSFuncHSV CSSFunc = "hsv"

	// CSS HSV color with alpha function name.
	// WARNING: This is not a part of the CSS3 spec, use with caution.
	CSSFuncHSVA CSSFunc = "hsva"

	// CSS HSB color function name.
	// WARNING: This is not a part of the CSS3 spec, use with caution.
	CSSFuncHSB CSSFunc = "hsb"

	// CSS HSB color with alpha function name.
	// WARNING: This is not a part of the CSS3 spec, use with caution.
	CSSFuncHSBA CSSFunc = "hsba"

	// CSS CMYK color function name.
	// WARNING: This is not a part of the CSS3 spec, use with caution.
	CSSFuncCMYK CSSFunc = "cmyk"

	// CSS CMYK color with alpha function name.
	// WARNING: This is not a part of the CSS3 spec, use with caution.
	CSSFuncCMYKA CSSFunc = "cmyka"

	// CSS HWB color function name.
	// WARNING: This is not a part of the CSS3 spec, use with caution.
	CSSFuncHWB CSSFunc = "hwb"

	// CSS HWB color with alpha function name.
	// WARNING: This is not a part of the CSS3 spec, use with caution.
	CSSFuncHWBA CSSFunc = "hwba"
)

// Matches returns whether or not the given target string matches the current
// CSSFunc value.
//
// This method's matching is not case sensitive.
//
// Arguments:
//   - tgt = Target string to match against.  The given string must start and
//           with the name of this CSSFunc value.  For example, "hsva" will not
//           match CSSFuncHSV.
//
// Returns:
//   - matches = Whether or not the entire tgt argument string matches this
//               CSS func name.
func (c CSSFunc) Matches(tgt []byte) (matches bool) {
	if len(tgt) != len(c) {
		return false
	}

	switch c {
	case CSSFuncHSB, CSSFuncHSV:
		return c.hsvb(tgt)
	case CSSFuncHSBA, CSSFuncHSVA:
		return c.hsvba(tgt)
	}

	for i, b := range util.StringToReadOnlyBytes((*string)(&c)) {
		if b != tgt[i] && b != util.ToggleCase(tgt[i]) {
			return false
		}
	}

	return true
}

// MinLength returns the minimum length of a valid call to the CSS function
// matching this CSSFunc value.
//
// Returns:
//   - min = minimum valid length for this CSSFunc's call string.
func (c CSSFunc) MinLength() (min uint8) {
	switch len(c) {
	case 3:
		// hsb(0,0,0)
		// hsl(0,0,0)
		// hsv(0,0,0)
		// hwb(0,0,0)
		// rgb(0,0,0)
		return 10
	case 4:
		// cmyk(0,0,0,0)
		// hsba(0,0,0,0)
		// hsla(0,0,0,0)
		// hsva(0,0,0,0)
		// hwba(0,0,0,0)
		// rgba(0,0,0,0)
		return 13
	case 5:
		// cmyka(0,0,0,0,0)
		return 16
	}

	panic("cannot size unknown CSS func")
}

func (c CSSFunc) hsvb(val []byte) bool {
	if val[0] != 'h' && val[0] != 'H' {
		return false
	}

	if val[1] != 's' && val[1] != 'S' {
		return false
	}

	if val[2] != 'v' && val[2] != 'b' && val[2] != 'V' && val[2] != 'B' {
		return false
	}

	return true
}

func (c CSSFunc) hsvba(val []byte) bool {
	if !c.hsvb(val) {
		return false
	}

	if val[3] != 'a' && val[3] != 'A' {
		return false
	}

	return true
}
