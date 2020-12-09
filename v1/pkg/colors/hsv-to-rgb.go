package colors

import "github.com/foxcapades/golour/v1/internal/util"

func (this HSV) ToRGB() (out RGB) {
	c := this.brightness * this.saturation
	h := float32(this.hue) / 60
	x := c * (1 - util.AbsF32(util.Mod(h, 2) - 1))

	r, g, b := float32(0), float32(0), float32(0)
	switch true {
	case h <= 1:
		r, g, b = c, x, 0
	case h <= 2:
		r, g, b = x, c, 0
	case h <= 3:
		r, g, b = 0, c, x
	case h <= 4:
		r, g, b = 0, x, c
	case h <= 5:
		r, g, b = x, 0, c
	case h <= 6:
		r, g, b = c, 0, x
	}

	m := this.brightness - c

	out.red = uint8((r+m)*255)
	out.green = uint8((g+m)*255)
	out.blue = uint8((b+m)*255)
	out.alpha = uint8(this.alpha * 255)

	return
}
