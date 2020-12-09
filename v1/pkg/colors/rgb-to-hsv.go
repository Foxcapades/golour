package colors

import "github.com/foxcapades/golour/v1/internal/util"

// ToHSV translates this RGB instance into an equivalent HSV value.
func (this *RGB) ToHSV() (out HSV) {
	r := float32(this.red) / 255
	g := float32(this.green) / 255
	b := float32(this.blue) / 255

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
	out.alpha = util.TruncateF32(float32(this.alpha) / 255, floatPrecision)

	return
}

// ToHSB is a convenience alias of ToHSV.
func (this *RGB) ToHSB() HSB {
	return this.ToHSV()
}
