package util

func NewError(msg string) *colorError {
	return &colorError{msg: msg}
}

type colorError struct {
	wrap  error
	value string
	ctx   string
	pos   int
	msg   string
}

func (c *colorError) Error() string {
	return c.String()
}

func (c *colorError) Unwrap() error {
	return c.wrap
}

func (c *colorError) Value() string {
	return c.value
}

func (c *colorError) Pos() int {
	return c.pos
}

func (c *colorError) String() string {
	if len(c.ctx) > 0 {
		return c.msg + ": " + c.ctx
	}

	return c.msg
}

func (c *colorError) WithContext(msg, value string, pos int) error {
	return &colorError{
		wrap:  c,
		value: value,
		ctx:   msg,
		pos:   pos,
		msg:   c.msg,
	}
}

func (c *colorError) WithValue(value string, pos int) error {
	return &colorError{
		wrap:  c,
		value: value,
		pos:   pos,
		msg:   c.msg,
	}
}

func (c *colorError) Is(err error) bool {
	if err == nil {
		if c == nil {
			return true
		}
		return false
	}

	for cur := err; cur != nil; {
		if c == cur {
			return true
		}

		if v, ok := cur.(interface{ Unwrap() error }); ok {
			cur = v.Unwrap()
		}
	}

	if v, ok := c.wrap.(interface { Is(error) bool }); ok {
		return v.Is(err)
	}

	return false
}