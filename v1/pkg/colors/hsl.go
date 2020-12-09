package colors

import "github.com/foxcapades/golour/v1/internal/util"

const floatPrecision = 3

// NewHSL constructs a new HSL instance with the given values.
//
// The alpha value for the newly created HSL will be 1.
func NewHSL(h int16, s, l float32) (out HSL) {
	out.SetHue(h)
	out.SetSaturation(util.TruncateF32(util.MinF32(s, 1), floatPrecision))
	out.SetLightness(util.TruncateF32(util.MinF32(l, 1), floatPrecision))
	out.SetAlpha(1)

	return
}

// NewHSLA constructs a new HSL instance with the given values.
func NewHSLA(h int16, s, l, a float32) (out HSL) {
	out.SetHue(h)
	out.SetSaturation(util.TruncateF32(util.MinF32(s, 1), floatPrecision))
	out.SetLightness(util.TruncateF32(util.MinF32(l, 1), floatPrecision))
	out.SetAlpha(util.TruncateF32(util.MinF32(a, 1), floatPrecision))

	return
}

// HSL represents a color value in the hsl format with an alpha value.
type HSL struct{
	hue        uint16
	saturation float32
	lightness  float32
	alpha      float32
}

// Hue returns the current hue value as a uint16 value.
func (this HSL) Hue() uint16 {
	return this.hue
}

// SetHue sets the HSL(a) Hue value to the given input value.
//
// Values are expected to be in the range [0..360), if they are not they will
// be adjusted to fit within that range.
func (this *HSL) SetHue(val int16) {
	this.hue = correctHue(val)
}

// Saturation returns the saturation value as a float value between 0 and 1.
func (this HSL) Saturation() float32 {
	return this.saturation
}

// SetSaturation sets the HSL(a) saturation value to the given float.
//
// Input values are expected to be between the range 0 and 1.  Any value greater
// than 1 passed to this value will be treated as 1.  Any value
func (this *HSL) SetSaturation(saturation float32) {
	this.saturation = util.TruncateF32(util.MaxF32(util.MinF32(saturation, 1), 0), floatPrecision)
}

func (this HSL) Lightness() float32 {
	return this.lightness
}

// SetLightness sets the HSL(a) lightness value to the given float.
//
// Input values are expected to be between the range 0 and 1.  Any value greater
// than 1 passed to this value will be treated as 1.
func (this *HSL) SetLightness(lightness float32) {
	this.lightness = util.TruncateF32(util.MaxF32(util.MinF32(lightness, 1), 0), floatPrecision)
}

// Alpha returns the current alpha value for this HSL(a) instance.
//
// This value is a percentage in the range [0, 1].
func (this HSL) Alpha() float32 {
	return this.alpha
}

// SetAlpha sets the HSL(a) alpha value to the given float.
//
// Input values are expected to be in the range [0, 1].  Any value greater than
// 1 passed to this value will be treated as 1.
func (this *HSL) SetAlpha(alpha float32) {
	this.alpha = util.TruncateF32(util.MaxF32(util.MinF32(alpha, 1), 0), floatPrecision)
}

// String returns the output of CSSFuncHSLA.
func (this HSL) String() string {
	return this.CSSFuncHSLA()
}

func correctHue(val int16) uint16 {
	for val < 0 {
		val += 360
	}
	for val >= 360 {
		val -= 360
	}

	return uint16(val)
}
