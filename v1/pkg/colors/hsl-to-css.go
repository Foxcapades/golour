package colors

import (
	"github.com/foxcapades/go-bytify/v0/bytify"
	"github.com/foxcapades/golour/v1/internal/util"
	"github.com/foxcapades/tally-go/v1/tally"
)

var leadHsl = [4]byte{'h', 's', 'l', '('}
func (this HSL) CSSFuncHSL() string {
	hue := this.Hue()
	sat := this.Saturation()
	lig := this.Lightness()

	// "hsl(, , )"
	size := 9 + bytify.Uint16StringSize(hue) +
		util.PercentStringSize(sat) +
		util.PercentStringSize(lig)

	out := make([]byte, size)
	copy(out, leadHsl[:])
	ind := tally.UTally8(4)

	ind.Add(bytify.Uint16ToBytes(hue, out[ind:]))
	util.WriteSeparator(out, &ind)
	ind.Add(util.PrecisionPercentToBytes(sat, out[ind:]))
	util.WritePercentSeparator(out, &ind)
	ind.Add(util.PrecisionPercentToBytes(lig, out[ind:]))
	out[ind.Inc()] = '%'
	out[ind] = ')'

	return string(out)
}

var leadHsla = [5]byte{'h', 's', 'l', 'a', '('}

func (this HSL) CSSFuncHSLA() string {
	hue := this.Hue()
	sat := this.Saturation()
	lig := this.Lightness()
	alp := this.Alpha()

	// 14 = len("hsla(, , , )")
	size := 12 + bytify.Uint16StringSize(hue) +
		util.PercentStringSize(sat) +
		util.PercentStringSize(lig) +
		util.F32StringSize(alp)

	ind := tally.UTally8(0)
	out := make([]byte, size)

	copy(out, leadHsla[:])
	ind.Add(5)

	ind.Add(bytify.Uint16ToBytes(hue, out[ind:]))
	util.WriteSeparator(out, &ind)
	ind.Add(util.PrecisionPercentToBytes(sat, out[ind:]))
	util.WritePercentSeparator(out, &ind)
	ind.Add(util.PrecisionPercentToBytes(lig, out[ind:]))
	util.WritePercentSeparator(out, &ind)
	util.AppendF32(alp, out, &ind)
	out[ind] = ')'

	return string(out)
}
