package colors

import (
	"github.com/foxcapades/go-bytify/v0/bytify"
	"github.com/foxcapades/golour/v1/internal/util"
	"github.com/foxcapades/tally-go/v1/tally"
)

var leadHsv = [4]byte{'h', 's', 'v', '('}
func (this HSV) CSSFuncHSV() string {
	// hsv(, , )
	size := 9 + bytify.Uint16StringSize(this.hue) +
		util.PercentStringSize(this.saturation) +
		util.PercentStringSize(this.brightness)

	out := make([]byte, size)
	copy(out, leadHsv[:])
	ind := tally.UTally8(4)

	ind.Add(bytify.Uint16ToBytes(this.hue, out[ind:]))
	util.WriteSeparator(out, &ind)
	ind.Add(util.PrecisionPercentToBytes(this.saturation, out[ind:]))
	util.WritePercentSeparator(out, &ind)
	ind.Add(util.PrecisionPercentToBytes(this.brightness, out[ind:]))
	out[ind.Inc()] = '%'
	out[ind.Inc()] = ')'

	return string(out)
}

var leadHsva = [5]byte{'h', 's', 'v', 'a', '('}
func (this HSV) CSSFuncHSVA() string {
	// len("hsva(, , , )")
	size := 12 + bytify.Uint16StringSize(this.hue) +
		util.PercentStringSize(this.saturation) +
		util.PercentStringSize(this.brightness) +
		util.F32StringSize(this.alpha)

	out := make([]byte, size)
	copy(out, leadHsva[:])
	ind := tally.UTally8(5)

	ind.Add(bytify.Uint16ToBytes(this.hue, out[ind:]))
	util.WriteSeparator(out, &ind)
	ind.Add(util.PrecisionPercentToBytes(this.saturation, out[ind:]))
	util.WritePercentSeparator(out, &ind)
	ind.Add(util.PrecisionPercentToBytes(this.brightness, out[ind:]))
	util.WritePercentSeparator(out, &ind)
	util.AppendF32(this.alpha, out, &ind)
	out[ind.Inc()] = ')'

	return string(out)
}
