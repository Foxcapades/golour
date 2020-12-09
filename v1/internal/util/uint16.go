package util

func ParseU16(v []byte) (uint16, error) {
	ln := len(v)

	if ln == 0 {
		return 0, ErrEmptyNumberVal.WithValue(string(v), 0)
	} else if ln > 3 {
		return 0, ErrU16TooBig.WithValue(string(v), 0)
	}

	ln--
	count := uint16(0)
	stage := uint(0)

	for ; ln >= 0; ln-- {
		d, err := DigitToU8(v[ln])

		if err != nil {
			return 0, ErrInvalidUint8Fmt.WithValue(string(v), int(count))
		}

		stage += uint(d) * U16Pow(10, count)
		count++
	}

	if stage > 255 {
		return 0, ErrU8TooBig.WithValue(string(v), 0)
	}

	return uint16(stage), nil
}
