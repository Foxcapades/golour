package colors

import (
	"github.com/foxcapades/golour/v1/internal/util"
	"github.com/foxcapades/tally-go/v1/tally"
)

// HexRGB returns this RGB value as a standard 6 digit hex color code
// including the leading '#' character.
func (this RGB) HexRGB() string {
	off := tally.UTally8(0)
	out := [7]byte{}
	this.hexRgb(out[:], &off)

	return string(out[:])
}

// HexRGBA returns this RGB value as an 8 digit hex color code
// including the leading '#' character.
func (this RGB) HexRGBA() string {
	out := [9]byte{}
	off := tally.UTally8(0)
	this.hexRgb(out[:], &off)

	util.AppendHex(this.alpha, out[:], &off)

	return string(out[:])
}

func (this *RGB) hexRgb(val []byte, off *tally.UTally8) {
	val[off.Inc()] = '#'

	util.AppendHex(this.red, val, off)
	util.AppendHex(this.green, val, off)
	util.AppendHex(this.blue, val, off)
}