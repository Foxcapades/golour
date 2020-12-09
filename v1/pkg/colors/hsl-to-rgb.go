package colors

import "github.com/foxcapades/golour/v1/internal/util"

func (this HSL) ToRGB() (out RGB) {
	H := float32(this.Hue())
	S := this.Saturation()
	L := this.Lightness()
	A := this.Alpha()

	if S == 0 {
		tmp := uint8(L*255)

		out.SetRed(tmp)
		out.SetGreen(tmp)
		out.SetBlue(tmp)
		out.SetAlpha(uint8(A * 100))

		return
	}

	C := (1 - util.AbsF32(2 * L - 1)) * S
	X := C * (1 - util.AbsF32(util.Mod(H/60, 2) - 1))
	m := L - C/2

	Rʼ, Gʼ, Bʼ := float32(0), float32(0), float32(0)
	switch true {
	case H < 60:
		Rʼ, Gʼ, Bʼ = C, X, 0
	case H < 120:
		Rʼ, Gʼ, Bʼ = X, C, 0
	case H < 180:
		Rʼ, Gʼ, Bʼ = 0, C, X
	case H < 240:
		Rʼ, Gʼ, Bʼ = 0, X, C
	case H < 300:
		Rʼ, Gʼ, Bʼ = X, 0, C
	case H < 360:
		Rʼ, Gʼ, Bʼ = C, 0, X
	}

	R, G, B := (Rʼ+m)*255, (Gʼ+m)*255, (Bʼ+m)*255

	out.SetRed(uint8(R))
	out.SetGreen(uint8(G))
	out.SetBlue(uint8(B))
	out.SetAlpha(uint8(A * 100))

	return
}
