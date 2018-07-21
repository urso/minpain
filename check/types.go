package check

import "github.com/urso/minpain/types"

type Type interface {
	types.Type
}

type fnSignature struct {
	ret  Type
	args []Type
}

func (fn *fnSignature) Extends() types.Type { return types.Def }
func (fn *fnSignature) Return() Type        { return fn.ret }
func (fn *fnSignature) NumArguments() int   { return len(fn.args) }
func (fn *fnSignature) Argument(i int) Type { return fn.args[i] }
