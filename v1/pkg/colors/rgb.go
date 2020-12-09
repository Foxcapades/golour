package colors

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
func (this RGB) Red() uint8 {
	return this.red
}

func (this *RGB) SetRed(red uint8) {
	this.red = red
}

// Green returns the green channel value for this RGB(A) color.
func (this RGB) Green() uint8 {
	return this.green
}

func (this *RGB) SetGreen(green uint8) {
	this.green = green
}

// Blue returns the blue channel value for this RGB(A) color.
func (this RGB) Blue() uint8 {
	return this.blue
}

func (this *RGB) SetBlue(blue uint8) {
	this.blue = blue
}

// Alpha returns the alpha channel value for this RGB(A) color.
//
// Note: Alpha values in RGB hex codes is in early stages of browser adoption
// and may not be supported for your use case.
func (this RGB) Alpha() uint8 {
	return this.alpha
}

func (this *RGB) SetAlpha(alpha uint8) {
	this.alpha = alpha
}

func (this RGB) String() string {
	return this.CSSFuncRGBA()
}
