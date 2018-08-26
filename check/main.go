// +build ignore

package main

import (
	"fmt"

	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/ast/print"
	"github.com/urso/minpain/check"
	"github.com/urso/minpain/errors"
	"github.com/urso/minpain/parser"
	"github.com/urso/minpain/types"
)

type typeMap map[string]check.Type

func (t typeMap) Exists(name string) bool       { return t.Lookup(name) != nil }
func (t typeMap) Lookup(name string) check.Type { return t[name] }

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
	test("null")
	test("int x = 1")
	test("boolean x = true")
	test("def x = 1")
	test(`
		int fn(int a) {
			int x = a + 1;
			return x + 2
		}
		int y = fn(3)
	`)
	test(`
	  if (true) {}
	`)
	test(`
	  boolean b = true; if (b) {}
	`)
	test(`
	  def b = true; if (b) {}
	`)
	test(`
	  for (int i = 0; i < 10; i++) {}
	`)
	test(`
	  for (def i = 0; i < 10; i++) {}
	`)
	test(`
	  int i = 0;
	  for (; i < 10; i++) {}
	`)
	test(`
	  for (;;) {}
	`)
	test(`
	  int[] list;
	  for (i in list) {}
	`)
	test(`
	  int[] list;
		for (int i : list) {}
	`)
	test(`
	  def foo(int a) {
			for (int i = 0; i < a; i++) {
				for (int j = 0; j < a; j++) {
				}
			}

			return a
		}
	`)
	test(`
	int fib(int n) {
		if (n == 0) {
			return 0
		}
		if (n == 1) {
			return 1
		}

		int a = 0, b = 1;
		for (int i = 2; i < n; i++) {
			int c = a + b;
			a = b;
			b = c
		}
		return b
	}

	fib(100)
	`)

	// TODO: add support for exceptions
	//test(`
	//	try {
	//		def b = 2
	//	} catch(def e) {}
	//`)

	// test some errors
	test("break")
	test("continue")
	test("boolean x = 1")
	test("int b = 1; if (b) {}")
	test(`for (int i = 5; i; i++) {}`)
}

func test(in string) {
	source := "<test>"

	fmt.Printf("input: ```\n%v\n```\n", in)

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

	fmt.Println("tree:")
	print.Print(tree)

	check.Validate(errs, tree)
	if err := errs.Err(); err != nil {
		fmt.Println("Validation failed:\n", err)
		return
	}

	info := check.NewInfo(typeTable)
	scriptScope := check.Index(errs, source, info, tree.(*ast.Script), true)
	if err := errs.Err(); err != nil {
		fmt.Println("index failed:\n", err)
	}

	fmt.Println("scope entries")
	scriptScope.Each(func(obj check.Object) {
		fmt.Println(obj)
	})

	if len(info.Decl) > 0 {
		fmt.Println("declarations")
		for id, obj := range info.Decl {
			fmt.Printf("(%v) %v\n", id.Pos(), obj)
		}
	}

	if len(info.Used) > 0 {
		fmt.Println("symbol access")
		for id, obj := range info.Used {
			fmt.Printf("(%v) %v\n", id.Pos(), obj)
		}
	}

	check.Types(errs, info, scriptScope, tree.(*ast.Script), true)
	if err := errs.Err(); err != nil {
		fmt.Println("check failed:\n", err)
	}

	fmt.Println("")
}
