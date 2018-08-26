package ssa

import "github.com/urso/minpain/ast"

type Pos struct {
	ast.Pos
	Symbol string
}
