package util

var (
	ErrU8TooBig        = NewError("string value would overflow u8")
	ErrInvalidUint8Fmt = NewError("invalid uint8 format")
	ErrU16TooBig       = NewError("string value would overflow u16")
	ErrEmptyNumberVal  = NewError("cannot parse empty string as a number")
	ErrInvalidFloatFmt = NewError("invalid float format")
)


