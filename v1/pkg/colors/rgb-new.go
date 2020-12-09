package colors

// RawToRGB converts a raw uint32 value into an RGB instance by only parsing
// the 24 least significant bits.
//
// This method supports converting from `0xFF_FF_FF` (6 digits) into an RGB
// value with an alpha value set to 100%.
//
// The breakdown of the input is as follows:
//
//   Input Value   | Mask          | Becomes
//   0xFF_FF_FF_FF & 0xFF_00_00_00 = ignored
//   0xFF_FF_FF_FF & 0x00_FF_00_00 = Red
//   0xFF_FF_FF_FF & 0x00_00_FF_00 = Green
//   0xFF_FF_FF_FF & 0x00_00_00_FF = Blue
func RawToRGB(v uint32) (out RGB) {
	out.red = uint8(v >> 16) & 0xFF
	out.green = uint8(v >> 8) & 0xFF
	out.blue = uint8(v >> 0) & 0xFF
	out.alpha = 0xFF

	return
}

// RawToRGBAlpha converts a raw uint32 value into an RGB instance by parsing
// all 32 bits.
//
// This method supports converting from `0xFF_FF_FF_FF (8 digits) into an RGB
// value.  For most use cases, RawToRGB is likely the function you want.
//
// The breakdown of the input is as follows:
//
//   Input Value   | Mask          | Used for
//   0xFF_FF_FF_FF & 0xFF_00_00_00 = Red
//   0xFF_FF_FF_FF & 0x00_FF_00_00 = Green
//   0xFF_FF_FF_FF & 0x00_00_FF_00 = Blue
//   0xFF_FF_FF_FF & 0x00_00_00_FF = Alpha
func RawToRGBAlpha(v uint32) (out RGB) {
	out.red = uint8(v >> 24) & 0xFF
	out.green = uint8(v >> 16) & 0xFF
	out.blue = uint8(v >> 8) & 0xFF
	out.alpha = uint8(v) & 0xFF

	return
}
