package colors

import "github.com/foxcapades/golour/v1/internal/util"

// NewRGB constructs a new RGB instance with the given values.
//
// The alpha value for the newly created RGB will be 1.
func NewRGB(r, g, b uint8) RGB {
	return NewRGBA(r, g, b, 0xFF)
}

// NewRGB constructs a new RGB instance with the given values.
func NewRGBA(r, g, b, a uint8) RGB {
	return RGB{r, g, b, a}
}

// RGB defines an RGB color value including an alpha channel value.
//
// Note: Alpha values in RGB hex codes is in early stages of browser adoption
// and may not be supported for your use case.
type RGB struct {
	red   uint8
	green uint8
	blue  uint8
	alpha uint8
}

// Red returns the red channel value for this RGB(A) color.
func (R RGB) Red() uint8 {
	return R.red
}

// SetRed overwrites the current red value of this RGB(A) color with the given
// value.
func (R *RGB) SetRed(red uint8) {
	R.red = red
}

// RedF32 returns the red channel value for this RGB(A) color as a float32
// value.
//
// Value is translated as:
//   red / 255
func (R RGB) RedF32() float32 {
	return util.BluntRound(float32(R.red) / 255)
}

// SetRedF32 overwrites the current red value of this RGB(A) color with the
// given value.
//
// Value is translated as:
//   ⌊red * 255⌋
func (R *RGB) SetRedF32(red float32) {
	R.red = uint8(util.ClampF32(red, 0, 1) * 255)
}

// Green returns the green channel value for this RGB(A) color.
func (R RGB) Green() uint8 {
	return R.green
}

func (R *RGB) SetGreen(green uint8) {
	R.green = green
}

// GreenF32 returns the green channel value for this RGB(A) color as a float32
// value.
//
// Value is translated as:
//   green / 255
func (R RGB) GreenF32() float32 {
	return util.BluntRound(float32(R.green) / 255)
}

// SetGreenF32 overwrites the current green value of this RGB(A) color with the
// given value.
//
// Value is translated as:
//   ⌊green * 255⌋
func (R *RGB) SetGreenF32(green float32) {
	R.green = uint8(util.ClampF32(green, 0, 1) * 255)
}

// Blue returns the blue channel value for this RGB(A) color.
func (R RGB) Blue() uint8 {
	return R.blue
}

func (R *RGB) SetBlue(blue uint8) {
	R.blue = blue
}

// BlueF32 returns the blue channel value for this RGB(A) color as a float32
// value.
//
// Value is translated as:
//   blue / 255.
func (R RGB) BlueF32() float32 {
	return util.BluntRound(float32(R.blue) / 255)
}

// SetBlueF32 overwrites the current blue value of this RGB(A) color with the
// given value.
//
// Value is translated as:
//   ⌊blue * 255⌋
func (R *RGB) SetBlueF32(blue float32) {
	R.blue = uint8(util.ClampF32(blue, 0, 1) * 255)
}

// Alpha returns the alpha channel value for this RGB(A) color.
//
// Note: Alpha values in RGB hex codes is in early stages of browser adoption
// and may not be supported for your use case.
func (R RGB) Alpha() uint8 {
	return R.alpha
}

func (R *RGB) SetAlpha(alpha uint8) {
	R.alpha = alpha
}

// AlphaF32 returns the alpha channel value for this RGB(A) color as a float32
// value.
//
// Value is translated as:
//   alpha / 255.
func (R RGB) AlphaF32() float32 {
	return util.BluntRound(float32(R.alpha) / 255)
}

// SetAlphaF32 overwrites the current alpha value of this RGB(A) color with the
// given value.
//
// Value is translated as:
//   ⌊alpha * 255⌋
func (R *RGB) SetAlphaF32(alpha float32) {
	R.alpha = uint8(util.ClampF32(alpha, 0, 1) * 255)
}

func (R RGB) String() string {
	return R.CSSFuncRGBA()
}

func (R RGB) RawValueNoAlpha() uint32 {
	return uint32(R.red)<<16 | uint32(R.green)<<8 | uint32(R.blue)
}

func (R RGB) RawValueWithAlpha() uint32 {
	return uint32(R.red)<<24 |
		uint32(R.green)<<16 |
		uint32(R.blue)<<8 |
		uint32(R.alpha)
}
