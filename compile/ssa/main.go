// +build ignore

package main

import (
	"fmt"
	"os"

	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/check"
	"github.com/urso/minpain/compile/ssa"
	"github.com/urso/minpain/errors"
	"github.com/urso/minpain/parser"
	"github.com/urso/minpain/types"
)

type typeMap map[string]check.Type

var typeTable = typeMap{
	"void": types.Void,
	"def":  types.Def,

	// objects
	"Object": types.Object,
	"Regex":  types.Regex,
	"String": types.String,

	// exception types
	"Exception": types.Exception,

	// primitive types:
	"boolean": types.Bool,
	"byte":    types.Byte,
	"short":   types.Short,
	"int":     types.Int,
	"long":    types.Long,

	"float":  types.Float,
	"double": types.Double,
}

func main() {
	test("true")
	test(`"test"`)
	test("1")
	// test("null")
	// test("int x = 1")
	// test("boolean x = true")
	// test("def x = 1")
}

func test(in string) {
	source := "<test>"
	fmt.Printf("input: ```\n%v\n```\n\n", in)

	errs := errors.NewLimitedMultiError(5)
	defer func() {
		if r := recover(); r != nil {
			me, ok := r.(*errors.LimitedMultiError)
			if !ok {
				panic(r)
			}

			fmt.Println("Fail:\n", me.Err())
		}
	}()

	p := parser.NewParser(errs, typeTable.Exists)
	tree, err := p.Parse(source, in, false)
	if err != nil {
		fmt.Println("Fail:", err)
		return
	}

	script := tree.(*ast.Script)
	info, err := check.Check(errs, typeTable, script, false)
	if err != nil {
		fmt.Println("Fail check:", err)
		return
	}

	progam, err := ssa.Build(info, script)
	if err != nil {
		fmt.Println("Fail build ssa:", err)
		return
	}

	progam.Format(os.Stdout)
}

func (t typeMap) Exists(name string) bool       { return t.Lookup(name) != nil }
func (t typeMap) Lookup(name string) check.Type { return t[name] }
