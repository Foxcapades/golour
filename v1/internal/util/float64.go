package util

// ParseFloat64 is a simplified float parser suitable for the specific needs of
// this library.
func ParseFloat64(v string) (float64, error) {
	ln := len(v)
	la := ln - 1

	if ln == 0 {
		return 0, ErrEmptyNumberVal.WithValue(v, 0)
	}

	count := uint8(0)
	stage := uint64(0)
	dpPos := -1
	for ; la >= 0; la-- {
		// We've hit the decimal marker, record it's position and move on.
		if v[la] == '.' {
			if dpPos > -1 {
				return 0, ErrInvalidFloatFmt.WithValue(v, la)
			}

			dpPos = la
			continue
		}

		d, err := DigitToU8(v[la])

		if err != nil {
			return 0, ErrInvalidUint8Fmt.WithValue(v, int(count))
		}

		stage += uint64(d) * uint64(U8Pow(10, count))
		count++
	}

	div := 1.0
	if dpPos > 0 {
		div = float64(IPow(10, ln-1-dpPos))
	}
	return float64(stage) / div, nil
}
