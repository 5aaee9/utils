package errors

import (
	"fmt"
	"io"
	"runtime"

	"github.com/pkg/errors"
)

func Trace(err error) error {
	if err == nil {
		return nil
	}

	if HasStack(err) {
		return err
	}

	return WithStack(err)
}

func WithStack(err error) error {
	if err == nil {
		return nil
	}

	return &ErrorWithStack{
		err,
		callers(),
	}
}

func HasStack(err error) bool {
	_, ok := err.(ErrorStack)
	return ok
}

type ErrorStack interface {
	StackTrace() errors.StackTrace
}

var _ ErrorStack = (*ErrorWithStack)(nil)

type ErrorWithStack struct {
	error
	*stack
}

func (w *ErrorWithStack) Cause() error { return w.error }

// Unwrap provides compatibility for Go 1.13 error chains.
func (w *ErrorWithStack) Unwrap() error { return w.error }

func (w *ErrorWithStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.Cause())
			w.stack.Format(s, verb)
			return
		}
		fallthrough
	case 's':
		io.WriteString(s, w.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.Error())
	}
}

type stack []uintptr

func (s *stack) Format(st fmt.State, verb rune) {
	switch verb {
	case 'v':
		switch {
		case st.Flag('+'):
			for _, pc := range *s {
				f := errors.Frame(pc)
				fmt.Fprintf(st, "\n%+v", f)
			}
		}
	}
}

func (s *stack) StackTrace() errors.StackTrace {
	f := make([]errors.Frame, len(*s))
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame((*s)[i])
	}
	return f
}

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}
