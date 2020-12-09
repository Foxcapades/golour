package util

import "github.com/foxcapades/tally-go/v1/tally"

// ParseU8Base10 attempts to parse the input bytes into a uint8 value.
func ParseU8Base10(v []byte) (uint8, error) {
	ln := len(v)

	if ln == 0 {
		return 0, ErrEmptyNumberVal.WithValue(string(v), 0)
	} else if ln > 3 {
		return 0, ErrU8TooBig.WithValue(string(v), 0)
	}

	ln--
	count := uint8(0)
	stage := uint(0)

	for ; ln >= 0; ln-- {
		d, err := DigitToU8(v[ln])

		if err != nil {
			return 0, ErrInvalidUint8Fmt.WithValue(string(v), int(count))
		}

		stage += uint(d) * U8Pow(10, count)
		count++
	}

	if stage > 255 {
		return 0, ErrU8TooBig.WithValue(string(v), 0)
	}

	return uint8(stage), nil
}

func ParseU8Base16(value []byte, offset *tally.UTally8) (out uint8, err error) {
	return HexBlock(value[offset.Inc()], value[offset.Inc()])
}

// DigitToU8 converts a single ASCII digit to a uint8 value.
//
// Handles base 10 and 16.
func DigitToU8(a byte) (uint8, error) {
	if a >= '0' && a <= '9' {
		return a - '0', nil
	}

	if a >= 'a' && a <= 'f' {
		return a - 87, nil
	}

	if a >= 'A' && a <= 'F' {
		return a - 55, nil
	}

	return 0, ErrInvalidUint8Fmt
}


// HexBlock converts a "block" of 2 ASCII hex characters into a uint8 value.
func HexBlock(a, b byte) (out uint8, err error) {
	tmp, err := DigitToU8(a)
	if err != nil {
		return
	}
	out = tmp << 4

	tmp, err = DigitToU8(b)
	out |= tmp

	return
}

const (
	hexStartLow = 48
	hexStartHigh = 55
)
func DigitToHex(val uint8) byte {
	if val < 10 {
		return val + hexStartLow
	}

	return val + hexStartHigh
}

func AppendHex(val uint8, buf []byte, off *tally.UTally8) {
	buf[off.Inc()] = DigitToHex(val >> 4)
	buf[off.Inc()] = DigitToHex(val & 0xF)
}
