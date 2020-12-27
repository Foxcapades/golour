package colors

import "github.com/foxcapades/golour/v1/internal/util"

// ToHSV translates this RGB instance into an equivalent HSV value.
func (R *RGB) ToHSV() (out HSV) {
	r := float32(R.red) / 255
	g := float32(R.green) / 255
	b := float32(R.blue) / 255

	max := util.MaxF32(r, g, b)
	min := util.MinF32(r, g, b)

	Δ := max - min

	h := float32(0)
	if Δ != 0 {
		switch max {
		case r:
			h = util.Mod((g-b)/Δ, 6)
		case g:
			h = ((b - r) / Δ) + 2
		case b:
			h = ((r - g) / Δ) + 4
		}
		h *= 60
	}

	s := float32(0)
	if max != 0 {
		s = Δ / max
	}

	v := max

	out.hue = uint16(h)
	out.saturation = s
	out.brightness = v
	out.alpha = util.TruncateF32(float32(R.alpha) / 255, floatPrecision)

	return
}

// ToHSB is a convenience alias of ToHSV.
func (R *RGB) ToHSB() HSB {
	return R.ToHSV()
}
