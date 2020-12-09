package colors

import "github.com/foxcapades/golour/v1/internal/util"

func (this *HSL) ToHSV() (out HSV) {
	out.hue = this.hue
	out.alpha = this.alpha

	out.brightness = util.TruncateF32(
		this.lightness + this.saturation*util.MinF32(this.lightness, 1-this.lightness),
		floatPrecision)

	if out.brightness == 0 {
		out.saturation = 0
	} else {
		out.saturation = util.TruncateF32(2 * (1 - this.lightness/out.brightness), floatPrecision)
	}


	return
}
