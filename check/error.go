package check

import (
	"fmt"

	"github.com/urso/minpain/ast"
)

type multiErr interface {
	Add(err error)
}

type nodeError struct {
	node ast.Node
	msg  string
}

type typeError struct {
	node ast.Node
	msg  string
}

func newNodeError(node ast.Node, msg string) *nodeError {
	return &nodeError{node, msg}
}
func newNodeErrorf(node ast.Node, s string, vs ...interface{}) *nodeError {
	return newNodeError(node, fmt.Sprintf(s, vs...))
}
func (e *nodeError) Node() ast.Node { return e.node }
func (e *nodeError) Pos() ast.Pos   { return e.node.Pos() }
func (e *nodeError) Error() string {
	pos := e.Pos()
	return fmt.Sprintf("%v:%v: %v", pos.Line, pos.Column, e.msg)
}

func newTypeError(node ast.Node, msg string) *typeError {
	return &typeError{node, msg}
}
func newTypeErrorf(node ast.Node, s string, vs ...interface{}) *typeError {
	return newTypeError(node, fmt.Sprintf(s, vs...))
}
func (e *typeError) Node() ast.Node { return e.node }
func (e *typeError) Pos() ast.Pos   { return e.node.Pos() }
func (e *typeError) Error() string {
	pos := e.Pos()
	return fmt.Sprintf("%v:%v: %v", pos.Line, pos.Column, e.msg)
}

/*
func (m *multiErr) append(errs []error) {
	*m = append(*m, errs...)
}
*/
