package util

var (
	ErrMalformedCSSFuncString = NewError("malformed css color function string, unable to parse")
	ErrBadCssFuncName         = NewError("incorrect or unexpected css function name")
)

const (
	TailComma   byte = ','
	TailParen   byte = ')'
	TailPercent byte = '%'
	LeadParen   byte = '('
)
