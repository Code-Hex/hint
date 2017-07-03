package hint

import (
	"fmt"
	"io"
)

type withHint struct {
	cause error
	hint  string
}

func Wrap(err error, suggest string) error {
	if err == nil {
		return nil
	}
	return &withHint{
		cause: err,
		hint:  "hint: " + suggest,
	}
}

func Wrapf(err error, format string, a ...interface{}) error {
	return Wrap(err, fmt.Sprintf(format, a...))
}

func (h *withHint) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v\n", h.Cause())
			io.WriteString(s, h.hint)
			return
		}
		fallthrough
	case 's', 'q':
		io.WriteString(s, h.Error())
	}
}

func (h *withHint) Hint() string  { return h.hint }
func (h *withHint) Error() string { return h.cause.Error() + "\n" + h.hint }
func (h *withHint) Cause() error  { return h.cause }
