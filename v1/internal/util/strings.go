package util

import (
	"github.com/foxcapades/tally-go/v1/tally"
)

const (
	minPer = 0.0001
)

func ToggleCase(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	}

	return c + 32
}

func FirstParen(input []byte) (position uint8, found bool) {
	ln := uint8(len(input))
	for ; position < ln; position++ {
		if input[position] == LeadParen {
			return position, true
		}
	}

	return
}

// SkipSpace iterates over the characters of the given string until it hits a
// non-space character.
//
// The given counter `i` will be incremented for each skipped character.
//
// Returns whether there are more characters left in the string.
func SkipSpace(v []byte, i *tally.UTally8) (hasMore bool){
	ln := uint8(len(v))
	for ; i.Cur() < ln && v[*i] == ' '; i.Inc() {
		// do nothing
	}
	return i.Cur() < ln
}

// WriteSeparator appends a function arg separator to the given byte slice.
func WriteSeparator(v []byte, i *tally.UTally8) {
	v[i.Inc()] = ','
	v[i.Inc()] = ' '
}

func WritePercentSeparator(v []byte, i *tally.UTally8) {
	v[i.Inc()] = '%'
	v[i.Inc()] = ','
	v[i.Inc()] = ' '
}

func IsDecimalDigit(c byte) bool {
	return IsIntegralDigit(c) || c == '.'
}

func IsIntegralDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func PercentBytes(f uint8, buf []byte) (written uint8) {
	if f == 0 {
		buf[0] = '0'
		return 1
	} else if f >= 100 {
		buf[0] = '1'
		return 1
	}

	buf[0] = '0'
	buf[1] = '.'

	if t := f % 10; t > 0 {
		buf[3] = t + '0'
		f -= t
		written = 4
	}

	f /= 10
	if t := f % 10; t > 0 {
		buf[2] = t + '0'
	} else {
		buf[2] = '0'
	}

	if written == 0 {
		written = 3
	}

	return
}