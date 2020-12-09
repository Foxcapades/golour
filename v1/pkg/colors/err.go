package colors

import "github.com/foxcapades/golour/v1/internal/util"

func NewColorError(msg string) ColorError {
	return util.NewError(msg)
}

type ColorError interface {
	error

	// String is an Alias of Error()
	String() string

	// Unwrap returns the wrapped parent error.
	Unwrap() error

	// Value returns the string value related to this error.
	Value() string

	// Pos returns the position in Value related to this error.
	Pos() int

	// WithValue returns a copy of this error with the given value and pos instead
	// of the current instance's value and pos.
	WithValue(val string, pos int) error

	// WithContext returns a copy of this error with the given context, value, and
	// position instead of the current instance's values.
	WithContext(message, value string, pos int) error

	// Is returns whether or not the given error is from the same base error as
	// this error instance.
	Is(err error) bool
}

// IsColorError is a convenience method that returns whether the given error
// is a ColorError instance.
func IsColorError(e error) bool {
	if _, ok := e.(ColorError); ok {
		return true
	}

	return false
}