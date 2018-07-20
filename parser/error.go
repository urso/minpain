package parser

import (
	"fmt"
	"strconv"

	"github.com/urso/minpain/ast"
)

type Errors struct {
	errs []error
}

type ParseError struct {
	op  string
	pos ast.Pos
	msg string
}

func (e *Errors) Causes() []error {
	return e.errs
}

func (e *Errors) Error() string {
	buf := new([]byte)
	for _, err := range e.errs {
		str := err.Error()
		if str == "" {
			continue
		}

		pad(buf, "\n")
		appendString(buf, str)
	}
	return string(*buf)
}

func (p *ParseError) Op() string      { return p.op }
func (p *ParseError) Pos() ast.Pos    { return p.pos }
func (p *ParseError) Message() string { return p.msg }
func (p *ParseError) Error() string {
	buf := new([]byte)

	if p.op != "" {
		appendString(buf, p.op)
	}

	pos := p.pos
	pad(buf, ": ")
	if pos.Source != "" {
		appendString(buf, fmt.Sprintf("%v:%v:%v", pos.Source, pos.Line, pos.Column))
	} else {
		appendString(buf, fmt.Sprintf("at %v:%v", pos.Line, pos.Column))
	}

	if p.msg != "" {
		pad(buf, ": ")
		appendString(buf, p.msg)
	}

	return string(*buf)
}

func pad(buf *[]byte, pattern string) {
	if len(*buf) > 0 {
		appendString(buf, pattern)
	}
}

func appendString(buf *[]byte, str string) {
	*buf = append(*buf, str...)
}

func appendInt(buf *[]byte, i int) {
	*buf = strconv.AppendInt(*buf, int64(i), 10)
}
