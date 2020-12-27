package colors

import (
	"github.com/foxcapades/golour/v1/internal/util"
)

// ToHSL returns an HSL value translated from this RGB value.
func (R RGB) ToHSL() (out HSL) {
	r := float32(R.Red()) / 255
	g := float32(R.Green()) / 255
	b := float32(R.Blue()) / 255

	max := util.MaxF32(r, g, b)
	min := util.MinF32(r, g, b)

	L := (max + min) / 2
	Δ := max - min

	S := float32(0)
	if Δ != 0 {
		S = Δ / (1 - util.AbsF32(2*L-1))
	}

	H := float32(0)
	if Δ > 0 {
		switch max {
		case r:
			H = util.Mod((g - b) / Δ, 6)
		case g:
			H = 2 + (b-r)/Δ
		case b:
			H = 4 + (r-g)/Δ
		}
	}
	H *= 60

	out.SetHue(int16(util.BluntRound(H/10)*10))
	out.SetSaturation(S)
	out.SetLightness(L)
	out.SetAlpha(util.TruncateF32(float32(R.Alpha()) / 255, floatPrecision))

	return
}
