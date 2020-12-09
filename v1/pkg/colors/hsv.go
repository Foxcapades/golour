package colors

import (
	"github.com/foxcapades/golour/v1/internal/util"
)

// NewHSV constructs a new HSV instance with the given values.
//
// The alpha value for the newly created HSV will be 1.
func NewHSV(h int16, s, v float32) (out HSV) {
	out.SetHue(h)
	out.SetSaturation(s)
	out.SetBrightness(v)
	out.SetAlpha(1)

	return
}

// NewHSVA constructs a new HSL instance with the given values.
func NewHSVA(h int16, s, v, a float32) (out HSV) {
	out.SetHue(h)
	out.SetSaturation(s)
	out.SetBrightness(v)
	out.SetAlpha(a)

	return
}

type HSV struct {
	hue        uint16
	saturation float32
	brightness float32
	alpha      float32
}

type HSB = HSV

// Hue returns the current hue value for this HSV instance.
//
// The returned value will be in the range [0, 360).
func (this HSV) Hue() uint16 {
	return this.hue
}

// SetHue overwrites the hue value for the this HSV instance.
//
// If the given value does not fall in the range [0, 360), it will be adjusted
// into that range.
//
// Adjustment Example:
//   foo.SetHue(450)
//   foo.Hue()  // Outputs: 90
func (this *HSV) SetHue(u int16) {
	this.hue = correctHue(u)
}

// Saturation returns the current saturation value for this HSV instance.
//
// The returned value will be in the range [0, 1].
func (this HSV) Saturation() float32 {
	return this.saturation
}

// SetSaturation overwrites the current saturation value for this HSV instance.
func (this *HSV) SetSaturation(f float32) {
	this.saturation = util.TruncateF32(util.MaxF32(util.MinF32(f, 1), 0), floatPrecision)
}

// Value is an alias of Brightness.
func (this HSV) Value() float32 {
	return this.brightness
}

// Brightness returns the current brightness value for this HSV instance.
//
// The returned value will be in the range [0, 1].
func (this HSV) Brightness() float32 {
	return this.brightness
}

// SetValue is an alias for SetBrightness.
func (this *HSV) SetValue(f float32) {
	this.brightness = util.TruncateF32(util.MaxF32(util.MinF32(f, 1), 0), floatPrecision)
}

// SetBrightness overwrites the current brightness value for this HSV instance.
func (this *HSV) SetBrightness(f float32) {
	this.brightness = util.TruncateF32(util.MaxF32(util.MinF32(f, 1), 0), floatPrecision)
}

// Alpha returns the current alpha value for this HSV instance.
//
// The returned value will be in the range [0, 1].
func (this HSV) Alpha() float32 {
	return this.alpha
}

// SetAlpha overwrites the current brightness value for this HSV instance.
func (this *HSV) SetAlpha(f float32) {
	this.alpha = util.TruncateF32(util.MaxF32(util.MinF32(f, 1), 0), floatPrecision)
}

// ToHSL returns a new HSL instance translated from the this HSV instance.
func (this HSV) ToHSL() (hsl HSL) {
	hsl.hue = this.hue
	hsl.lightness = util.TruncateF32(this.brightness * (1-this.saturation/2), floatPrecision)

	if hsl.lightness == 0 || hsl.lightness == 1 {
		hsl.saturation = 0
	} else {
		hsl.saturation = util.TruncateF32(
			(this.brightness-hsl.lightness) / util.MinF32(hsl.lightness, 1-hsl.lightness),
			floatPrecision)
	}
	hsl.alpha = this.alpha

	return
}

// String is an alias of CSSFuncHSVA.
func (this HSV) String() string {
	return this.CSSFuncHSVA()
}
