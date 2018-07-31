package errors

import "strings"

type MultiError struct {
	errs []error
}

type LimitedMultiError struct {
	max  int
	errs []error
}

func (e *MultiError) Len() int {
	return len(e.errs)
}

func (e *MultiError) Add(err error) {
	e.errs = append(e.errs, err)
}

func (e *MultiError) Err() error {
	if len(e.errs) == 0 {
		return nil
	}
	return e
}

func (e *MultiError) Error() string {
	if len(e.errs) == 0 {
		return "<no error>"
	}
	strs := make([]string, len(e.errs))
	for i, err := range e.errs {
		strs[i] = err.Error()
	}
	return strings.Join(strs, "\n")
}

func NewLimitedMultiError(max int) *LimitedMultiError {
	if max <= 0 {
		panic("max must be > 0")
	}
	return &LimitedMultiError{max: max, errs: nil}
}

func (e *LimitedMultiError) Len() int {
	return len(e.errs)
}

func (e *LimitedMultiError) Add(err error) {
	e.errs = append(e.errs, err)
	if len(e.errs) == e.max {
		panic(e)
	}
}

func (e *LimitedMultiError) Err() error {
	if len(e.errs) == 0 {
		return nil
	}
	return e
}

func (e *LimitedMultiError) Error() string {
	if len(e.errs) == 0 {
		return "<no error>"
	}
	strs := make([]string, len(e.errs))
	for i, err := range e.errs {
		strs[i] = err.Error()
	}
	return strings.Join(strs, "\n")
}
