package util

import (
	"github.com/foxcapades/tally-go/v1/tally"
)

// ParseCSSFloat4 parses a float/percent based css function with an arity of 3.
//
// This function assumes that the caller has already validated the css function
// name and is passing in an offset on the opening parenthesis character.
//
// Arguments:
//   - input  = string to parse
//   - offset = offset in the input string from which to start reading.  This
//              value should be pointing to the opening parenthesis character of
//              the CSS function call.  Value will almost always be 3.
//
//              Note: the caller need not verify the parenthesis character as
//              this function will perform that validation.
//
// Returns:
//   - h   = Hue value between [0, 359]
//   - s   = Saturation value between [0, 1]
//   - x   = Brightness/lightness value between [0, 1]
//   - a   = Alpha value defaulted to 1.0
//   - err = Any error that occurred while attempting to parse this CSS function
//           string.
func ParseCSSFloat3(input []byte, offset *tally.UTally8) (h uint16, s, x, a float32, err error) {
	h, s, x, err = cssFloatParserCommon(input, offset, TailParen)
	a = 1

	return
}

// ParseCSSFloat4 parses a float/percent based css function with an arity of 4.
//
// This function assumes that the caller has already validated the css function
// name and is passing in an offset on the opening parenthesis character.
//
// Arguments:
//   - input  = string to parse
//   - offset = offset in the input string from which to start reading.  This
//              value should be pointing to the opening parenthesis character of
//              the CSS function call.  Value will almost always be 4.
//
//              Note: the caller need not verify the parenthesis character as
//              this function will perform that validation.
//
// Returns:
//   - h   = Hue value between [0, 359]
//   - s   = Saturation value between [0, 1]
//   - x   = Brightness/lightness value between [0, 1]
//   - a   = Alpha value between [0, 1]
//   - err = Any error that occurred while attempting to parse this CSS function
//           string.
func ParseCSSFloat4(input []byte, offset *tally.UTally8) (h uint16, s, x, a float32, err error) {
	if h, s, x, err = cssFloatParserCommon(input, offset, TailComma); err != nil {
		return
	}

	// Parse alpha value
	a, err = parseFloatArg(input, offset, TailParen, false)

	return
}

func cssFloatParserCommon(
	input []byte,
	index *tally.UTally8,
	tail byte,
) (h uint16, s, x float32, err error) {
	if input[index.Inc()] != '(' {
		err = ErrMalformedCSSFuncString.WithValue(string(input), int(index.Cur()-1))
		return
	}

	if h, err = parseU16Arg(input, index, TailComma); err != nil {
		return
	}
	if s, err = parseFloatArg(input, index, TailComma, true); err != nil {
		return
	}
	if x, err = parseFloatArg(input, index, tail, true); err != nil {
		return
	}

	return
}

func parseU16Arg(
	input []byte,
	index *tally.UTally8,
	tail  byte,
) (val uint16, err error) {
	// Skip any space before the value
	if !SkipSpace(input, index) {
		return 0, ErrMalformedCSSFuncString.WithValue(string(input), int(index.Cur()-1))
	}

	buffer := [8]byte{}

	read := 0
	for ; IsIntegralDigit(input[index.Cur()]); index.Inc() {
		buffer[read] = input[index.Cur()]
		read++
	}

	// If we read nothing, it wasn't a valid number.
	if read == 0 {
		return 0, ErrMalformedCSSFuncString.
			WithContext("invalid or missing hue value", string(input), int(index.Cur()-1))
	}

	val, err = ParseU16(buffer[:read])
	if err != nil {
		return
	}

	// Skip any spaces between the value and the next comma.
	if !SkipSpace(input, index) {
		err = ErrMalformedCSSFuncString.WithValue(string(input), int(index.Cur()-1))
		return
	}

	if input[index.Inc()] != tail {
		err = ErrMalformedCSSFuncString.WithValue(string(input), int(index.Cur()-1))
	}

	return
}

func parseFloatArg(
	input []byte,
	index *tally.UTally8,
	tail  byte,
	allowPercent bool,
) (val float32, err error) {
	// Skip any leading space before the value
	if !SkipSpace(input, index) {
		return val, ErrMalformedCSSFuncString.
			WithContext("unexpected end of string", string(input), int(index.Cur()-1))
	}

	buffer := [8]byte{}

	// Read all [0-9.] characters into the buffer
	read := 0
	for ; IsDecimalDigit(input[*index]); index.Inc() {
		buffer[read] = input[*index]
		read++
	}

	// If we read nothing, it wasn't a valid number.
	if read == 0 {
		return 0, ErrMalformedCSSFuncString.WithValue(string(input), int(index.Cur()-1))
	}

	if val, err = ParseFloat32(buffer[:read]); err != nil {
		return
	}

	// Skip any trailing space after the value
	if !SkipSpace(input, index) {
		return val, ErrMalformedCSSFuncString.
			WithContext("unexpected end of string", string(input), int(index.Cur()-1))
	}

	// If this value is allowed to have a percent sign and the current byte is a
	// percent sign, skip it and continue.
	if allowPercent && input[index.Cur()] == TailPercent {
		index.Inc()

		// Skip any trailing space between the percent sign and the next comma.
		if !SkipSpace(input, index) {
			return val, ErrMalformedCSSFuncString.
				WithContext("unexpected end of string", string(input), int(index.Cur()-1))
		}
	}

	if input[index.Inc()] != tail {
		return 0, ErrMalformedCSSFuncString.
			WithContext(
				"unexpected character, wanted '" + string(tail) + "' got '" + string(input[*index-1]) + "'",
				string(input), int(index.Cur()-1))
	}

	return
}
