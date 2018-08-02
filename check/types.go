package check

import (
	"fmt"
	"strings"

	"github.com/urso/minpain/types"
)

type Type interface {
	types.Type
}

type fnSignature struct {
	ret  Type
	args []Type
}

func (fn *fnSignature) Extends() types.Type          { return types.Def }
func (fn *fnSignature) ComponentType() types.Type    { return nil }
func (fn *fnSignature) InstanceOf(t types.Type) bool { return t == types.Def }
func (fn *fnSignature) Return() types.Type           { return fn.ret }
func (fn *fnSignature) NumArguments() int            { return len(fn.args) }
func (fn *fnSignature) Argument(i int) types.Type    { return fn.args[i] }

func (fn *fnSignature) String() string {
	args := make([]string, len(fn.args))
	for i, t := range fn.args {
		args[i] = t.String()
	}
	return fmt.Sprintf("%v(%v)", fn.ret, strings.Join(args, ", "))
}

var _ types.Signature = &fnSignature{}
