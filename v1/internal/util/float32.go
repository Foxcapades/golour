package util

import (
	"github.com/foxcapades/go-bytify/v0/bytify"
	"github.com/foxcapades/tally-go/v1/tally"
)

func ClampF32(value, min, max float32) float32 {
	if value > max {
		return max
	} else if value < min {
		return min
	} else {
		return value
	}
}

func TruncateF32(value float32, precision int) float32 {
	mag := IPow(10, precision)
	return float32(int64(value*float32(mag))) / float32(mag)
}

// ParseFloat32 is a simplified float parser suitable for the specific needs of
// this library.
func ParseFloat32(v []byte) (float32, error) {
	ln := len(v)
	la := ln - 1

	if ln == 0 {
		return 0, ErrEmptyNumberVal.WithValue(string(v), 0)
	}

	count := uint8(0)
	stage := uint32(0)
	dpPos := -1
	for ; la >= 0; la-- {
		// We've hit the decimal marker, record it's position and move on.
		if v[la] == '.' {
			if dpPos > -1 {
				return 0, ErrInvalidFloatFmt.WithValue(string(v), la)
			}

			dpPos = la
			continue
		}

		d, err := DigitToU8(v[la])

		if err != nil {
			return 0, ErrInvalidUint8Fmt.WithValue(string(v), int(count))
		}

		stage += uint32(d) * uint32(U8Pow(10, count))
		count++
	}

	div := float32(1.0)
	if dpPos > 0 {
		div = float32(IPow(10, ln-1-dpPos))
	}
	return float32(stage) / div, nil
}

func PrecisionPercentToBytes(per float32, buf []byte) (written uint8) {
	if per < 0 {
		per = -per
	}

	if per <= minPer {
		buf[0] = '0'
		return 1
	} else if per >= 1 {
		buf[0] = '1'
		buf[1] = '0'
		buf[2] = '0'
		return 3
	}

	big := BluntRound(per * 100)
	tmp := uint8(big)
	written = bytify.Uint8ToBytes(tmp, buf)
	if t := big - float32(tmp); t == 0 {
		return
	} else {
		buf[written] = '.'
		written++
		buf[written] = uint8(t*10) + '0'
		written++
	}

	return
}

func AppendF32(val float32, buf []byte, off *tally.UTally8) {
	val = AbsF32(val)

	if val == 0 {
		buf[off.Inc()] = '0'
	} else if val >= 1 {
		buf[off.Inc()] = '1'
	} else {
		buf[off.Inc()] = '0'
		buf[off.Inc()] = '.'

		floor := int32(val * 100)

		buf[off.Inc()] = byte(floor/10) + '0'

		if mod := floor % 10; mod > 0 {
			buf[off.Inc()] = byte(mod) + '0'
		}
	}
}

func F32StringSize(val float32) (size uint8) {
	val = AbsF32(val)

	if val == 0 || val >= 1 {
		return 1
	}

	floor := int(val * 100)
	size = 2 // "0."

	// If we have something in the hundredths
	if floor%10 > 0 {
		size += 2
	} else {
		size++
	}

	return
}

func PercentStringSize(per float32) (size uint8) {
	if per < 0 {
		per = -per
	}

	if per <= minPer {
		return 2
	} else if per >= 1 {
		return 4
	}

	rounded := BluntRound(per * 100)
	floor := uint8(rounded)
	size = bytify.Uint8StringSize(floor) + 1 // (+1 for '%' character)

	// If rounded - floor == 0 then we have no decimal places.
	if t := rounded - float32(floor); t == 0 {
		return size
	}

	// If we do have decimal places, add 2 to account for the leading `0.`
	return size + 2
}
