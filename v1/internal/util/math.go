package util

func MinU8(a, b uint8) uint8 {
	if a < b {
		return a
	}

	return b
}

func MinI(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func AbsF32(a float32) float32 {
	if a < 0 {
		return -a
	}
	return a
}

func Mod(a, b float32) float32 {
	neg := a < 0
	if neg {
		a = -a
	}
	if b < 0 {
		b = -b
	}

	for a >= b {
		a -= b
	}

	if neg {
		return -a
	}

	return a
}

func MinF32(floats ...float32) (out float32) {
	out = floats[0]

	for i := 1; i < len(floats); i++ {
		if floats[i] < out {
			out = floats[i]
		}
	}

	return
}

func MinF64(floats ...float64) (out float64) {
	out = floats[0]

	for i := 1; i < len(floats); i++ {
		if floats[i] < out {
			out = floats[i]
		}
	}

	return
}

func MaxF32(a ...float32) (out float32) {
	for i := range a {
		if a[i] > out {
			out = a[i]
		}
	}

	return
}

func MaxF64(a ...float64) (out float64) {
	for i := range a {
		if a[i] > out {
			out = a[i]
		}
	}

	return
}

func U8Pow(val, pow uint8) (out uint) {
	if pow == 0 {
		return 1
	}

	out = 1

	for i := uint8(0); i < pow; i++ {
		out *= uint(val)
	}

	return
}

func U16Pow(val, pow uint16) (out uint) {
	if pow == 0 {
		return 1
	}

	out = 1

	for i := uint16(0); i < pow; i++ {
		out *= uint(val)
	}

	return
}

func IPow(val, pow int) (out int) {
	if pow == 0 {
		return 1
	}

	out = 1

	for i := 0; i < pow; i++ {
		out *= val
	}

	return
}

// BluntRound performs a dumb round to a precision of 1/10, using a max
// precision of 1/1000 to perform the rounding.
func BluntRound(b float32) (out float32) {
	tmp := int32(b * 1_000)
	mod := tmp % 100
	out = float32(tmp - mod)

	if mod >= 50 {
		out += 100
	}

	return out / 1000
}

// BluntRoundByte64 performs a dumb round to an integral value using a max
// precision of 1/100 to perform the rounding.
func BluntRoundByte64(v float64) uint8 {
	t := uint32(v * 100)
	m := t % 100
	b := float64(t - m)

	if m >= 50 {
		b += 100
	}

	return uint8(b / 100)
}
