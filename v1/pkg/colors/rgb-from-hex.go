package colors

import (
	"github.com/foxcapades/golour/v1/internal/util"
	"github.com/foxcapades/tally-go/v1/tally"
)

const (
	shortHexLen        = uint8(3)

	minValHexStr = shortHexLen
	maxByte      = uint8(0xFF)

	leadHex = '#'
)

var (
	ErrRGBHexTooShort  = NewColorError("input string is too short to be a valid rgb(a) hex value")
	ErrRGBHexBadFormat = NewColorError("invalid hex color format")
)

func NewRGBFromHex(hex string) (rgb RGB, err error) {
	length := uint8(len(hex))

	if length < minValHexStr {
		return rgb, ErrRGBHexTooShort
	}

	bytes := util.StringToReadOnlyBytes(&hex)
	offset := tally.UTally8(0)

	switch length {

	case 3: // RGB
		rgb.red, rgb.green, rgb.blue, rgb.alpha, err = Parse3DigitRGBHex(bytes, &offset)
	case 4: // RGB | #RGB
		rgb.red, rgb.green, rgb.blue, rgb.alpha, err = Parse4DigitRGBHex(bytes, &offset)
	case 5: // #RGB
		rgb.red, rgb.green, rgb.blue, rgb.alpha, err = Parse5DigitRGBHex(bytes, &offset)
	case 6: // RRGGBB
		rgb.red, rgb.green, rgb.blue, rgb.alpha, err = Parse6DigitRGBHex(bytes, &offset)
	case 7: // #RRGGBB
		rgb.red, rgb.green, rgb.blue, rgb.alpha, err = Parse7DigitRGBHex(bytes, &offset)
	case 8: // RRGGBBAA
		rgb.red, rgb.green, rgb.blue, rgb.alpha, err = Parse8DigitRGBHex(bytes, &offset)
	case 9: // #RRGGBBAA
		rgb.red, rgb.green, rgb.blue, rgb.alpha, err = Parse9DigitRGBHex(bytes, &offset)
	}

	return
}

// Parse3DigitRGBHex parses a given 3 digit hex string (RGB) into separate
// values.
func Parse3DigitRGBHex(value []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	var dum tally.UTally8
	var buf [6]byte

	buf[0] = value[offset.Cur()]
	buf[1] = value[offset.Inc()]

	buf[2] = value[offset.Cur()]
	buf[3] = value[offset.Inc()]

	buf[4] = value[offset.Cur()]
	buf[5] = value[offset.Inc()]

	return Parse6DigitRGBHex(buf[:], &dum)
}

// Parse4DigitRGBHex parses a given 4 digit hex string (#RGB or RGB) into
// separate values.
func Parse4DigitRGBHex(value []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	if value[*offset] == leadHex {
		offset.Inc()
		return Parse3DigitRGBHex(value, offset)
	}

	if r, g, b, a, err = Parse3DigitRGBHex(value, offset); err != nil {
		return
	}

	buf := [2]byte{value[*offset], value[*offset]}
	dum := tally.UTally8(0)

	offset.Add(2)
	a, err = util.ParseU8Base16(buf[:], &dum)

	return
}

// Parse5DigitRGBHex parses a given 5 digit hex string (#RGB) into separate
// values.
func Parse5DigitRGBHex(value []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	offset.Inc()
	return Parse4DigitRGBHex(value, offset)
}

// Parse6DigitRGBHex parses a given 6 digit hex string (RRGGBB) into separate
// values.
func Parse6DigitRGBHex(value []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	if r, err = util.ParseU8Base16(value, offset); err != nil {
		return
	}

	if g, err = util.ParseU8Base16(value, offset); err != nil {
		return
	}

	if b, err = util.ParseU8Base16(value, offset); err != nil {
		return
	}

	a = maxByte

	return
}

// Parse7DigitRGBHex parses a given 7 digit hex string (#RRGGBB) into separate
// values.
func Parse7DigitRGBHex(value []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	if value[offset.Inc()] != leadHex {
		err = ErrRGBHexBadFormat.WithValue(string(value), int(offset.Cur()-1))
		return
	}

	return Parse6DigitRGBHex(value, offset)
}

// Parse8DigitRGBHex parses a given 8 digit hex string (RRGGBBAA) into separate
// values.
func Parse8DigitRGBHex(value []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	if r, g, b, a, err = Parse6DigitRGBHex(value, offset); err != nil {
		return
	}

	if a, err = util.ParseU8Base16(value, offset); err != nil {
		return
	}

	return
}

// Parse9DigitRGBHex parses a given 9 digit hex string (#RRGGBBAA) into separate
// values.
func Parse9DigitRGBHex(value []byte, offset *tally.UTally8) (r, g, b, a uint8, err error) {
	if value[offset.Inc()] != leadHex {
		err = ErrRGBHexBadFormat.WithValue(string(value), int(offset.Cur()-1))
		return
	}

	return Parse8DigitRGBHex(value, offset)
}
