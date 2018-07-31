// +build ignore

package main

import (
	"fmt"

	"github.com/urso/minpain/errors"
	"github.com/urso/minpain/parser"
	"github.com/urso/minpain/print"
)

var debug = false

func main() {
	test(`null`)
	test(`return true`)
	test(`return false`)
	test(`"string1"`)
	test(`'string2'`)
	test(`0700`)
	test(`0x42`)
	test(`23`)
	test(`3.14`)
	test(`/match/`)
	test(`[]`)
	test(`[:]`)
	test(`abc`)
	test(`foo()`)
	test(`new def()`)
	test(`foo.bar()`)
	test(`foo?.bar()`)
	test(`foo.bar`)
	test(`foo?.1`)
	test(`foo["bar"]`)
	test(`String.test()`)
	test(`new int[5][5];`)
	test(`new def[] {1,2};`)
	test(`new def[] {1,2}.size()`)
	test(`10 / 2`)
	test(`try {
		a = b
	} catch(Exception e) {}`)
	test(`if (a > 0) { x = i }`)
	test(`for (def i = 0; i < a; i++) { x = i }`)
	test(`for (;;) { }`)
	test(`for (;true;) { }`)
	test(`for (;i < 10; i++) { }`)
	test(`for (i = 0;i < 10; ) { }`)
	test(`for (i = 0;i < 10; i++ ) { }`)
	test(`for (def i = 0;i < 10; i++ ) { }`)
	test(`
	if (a > 5) {
    for (def i=0;i < 10; i++) {
			if (a < i) {
				break
			}
		}
	}`)
	test(`def emptyFn() {}`)
	test(`def add(def a, def b) { return a + b }`)
	test(`abc(x -> x + 1)`)

	test(`
	int fib(int n) {
		if (n == 0) {
			return 0
		}
		if (n <= 2) {
			return 1
		}

		int a=1;
		int b=1;
		for (int i = 2; i < n; i++) {
			int c = a + b;
			a = b;
			b = c
		}
		return b
	}

	fib(100)
	`)
}

func test(in string) {
	fmt.Printf("input: ```\n%v\n```\n", in)

	errs := errors.NewLimitedMultiError(10)

	p := parser.NewParser(errs, func(str string) bool {
		switch str {
		case "def", "boolean", "int", "String", "Exception":
			return true
		}
		return false
	})

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Fail:\n", r)
		}
	}()
	v, err := p.Parse("<test>", in, debug)

	if err != nil {
		fmt.Println("Fail:", err)
	} else {
		print.Print(v)
	}
	fmt.Println()
}
