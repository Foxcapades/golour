package util

import "github.com/foxcapades/tally-go/v1/tally"

// CSSByte3Parser parses a uint8 based css function with an arity of 3.
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
//   - r   = Red value
//   - g   = Green value
//   - b   = Blue value
//   - a   = Alpha value defaulted to 255
//   - err = Any error that occurred while attempting to parse this CSS function
//           string.
func CSSByte3Parser(input []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	r, g, b, err = cssByteParserCommon(input, offset, TailParen)
	a = 255

	return
}

// CSSByte4Parser parses a uint8 based css function with an arity of 4.
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
//   - r   = Red value
//   - g   = Green value
//   - b   = Blue value
//   - a   = Alpha value
//   - err = Any error that occurred while attempting to parse this CSS function
//           string.
func CSSByte4Parser(input []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	if r, g, b, err = cssByteParserCommon(input, offset, TailComma); err != nil {
		return
	}

	a, err = parseStrU8(input, offset, TailParen)

	return
}

func cssByteParserCommon(
	input []byte,
	offset *tally.UTally8,
	tail byte,
) (r, g, b uint8, err error) {
	if input[offset.Inc()] != LeadParen {
		return r, g, b, ErrMalformedCSSFuncString.WithValue(string(input), int(offset.Cur()-1))
	}

	if r, err = parseStrU8(input, offset, TailComma); err != nil {
		return
	}
	if g, err = parseStrU8(input, offset, TailComma); err != nil {
		return
	}
	if b, err = parseStrU8(input, offset, tail); err != nil {
		return
	}

	return
}

func parseStrU8(input []byte, index *tally.UTally8, tail byte) (val uint8, err error){
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
		return 0, ErrMalformedCSSFuncString.WithValue(string(input), int(index.Cur()-1))
	}

	val, err = ParseU8Base10(buffer[:read])
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