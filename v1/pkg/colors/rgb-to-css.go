package colors

import (
	"github.com/foxcapades/go-bytify/v0/bytify"
	"github.com/foxcapades/golour/v1/internal/util"
	"github.com/foxcapades/tally-go/v1/tally"
)

func (R RGB) CSSFuncRGBA() string {
	r := R.Red()
	g := R.Green()
	b := R.Blue()
	a := R.Alpha()

	buf := make([]byte, bytify.Uint8StringSize(r)+
		bytify.Uint8StringSize(g)+
		bytify.Uint8StringSize(b)+
		bytify.Uint8StringSize(a)+12)
	ind := tally.UTally8(0)

	buf[ind.Inc()], buf[ind.Inc()],
		buf[ind.Inc()], buf[ind.Inc()],
		buf[ind.Inc()] = 'r', 'g', 'b', 'a', '('

	ind.Add(bytify.Uint8ToBytes(r, buf[ind:]))
	util.WriteSeparator(buf, &ind)
	ind.Add(bytify.Uint8ToBytes(g, buf[ind:]))
	util.WriteSeparator(buf, &ind)
	ind.Add(bytify.Uint8ToBytes(b, buf[ind:]))
	util.WriteSeparator(buf, &ind)
	ind.Add(bytify.Uint8ToBytes(a, buf[ind:]))
	buf[ind] = ')'

	return string(buf)
}

func (R RGB) CSSFuncRGB() string {
	r := R.Red()
	g := R.Green()
	b := R.Blue()

	// rgb(, , )
	buf := make([]byte, bytify.Uint8StringSize(r)+
		bytify.Uint8StringSize(g)+
		bytify.Uint8StringSize(b)+9)
	ind := tally.UTally8(0)

	buf[ind.Inc()], buf[ind.Inc()],
		buf[ind.Inc()], buf[ind.Inc()] = 'r', 'g', 'b', '('

	ind.Add(bytify.Uint8ToBytes(r, buf[ind:]))
	util.WriteSeparator(buf, &ind)
	ind.Add(bytify.Uint8ToBytes(g, buf[ind:]))
	util.WriteSeparator(buf, &ind)
	ind.Add(bytify.Uint8ToBytes(b, buf[ind:]))
	buf[ind] = ')'

	return string(buf)
}
