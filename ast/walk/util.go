package walk

import "github.com/urso/minpain/ast"

type EnvSetuper interface {
	Enter(node ast.Node) bool
	Exit(node ast.Node)
}

type filterVisitor struct {
	pred func(ast.Node) bool
}

type envVisitor struct {
	enter func(ast.Node) bool
	exit  func(ast.Node)
}

type rulesVisitor struct {
	visitors []Visitor
}

type fromTop func(node ast.Node) bool
type fromBottom func(node ast.Node)

func Rules(vs ...Visitor) Visitor {
	return &rulesVisitor{vs}
}

func EnvSetup(s EnvSetuper) Visitor {
	return EnvSetupWith(s.Enter, s.Exit)
}
func EnvSetupWith(enter func(ast.Node) bool, exit func(ast.Node)) Visitor {
	return &envVisitor{enter: enter, exit: exit}
}
func (v *envVisitor) Enter(node ast.Node) (visit, callExit bool) { return true, v.enter(node) }
func (v *envVisitor) Exit(node ast.Node)                         { v.exit(node) }
func (v *envVisitor) FromTop(node ast.Node) (visit bool)         { return true }
func (v *envVisitor) FromBottom(node ast.Node)                   {}

func Filter(fn func(ast.Node) bool) Visitor {
	return &filterVisitor{fn}
}
func (f *filterVisitor) Enter(node ast.Node) (visit, callExit bool) { return f.pred(node), false }
func (f *filterVisitor) Exit(node ast.Node)                         {}
func (f *filterVisitor) FromTop(node ast.Node) (visit bool)         { return true }
func (f *filterVisitor) FromBottom(node ast.Node)                   {}

func FromTop(fn func(ast.Node) bool) Visitor                 { return fromTop(fn) }
func (f fromTop) Enter(node ast.Node) (visit, callExit bool) { return true, false }
func (f fromTop) Exit(node ast.Node)                         {}
func (f fromTop) FromTop(node ast.Node) (visit bool)         { return f(node) }
func (f fromTop) FromBottom(node ast.Node)                   {}

func FromBottom(fn func(ast.Node)) Visitor                      { return fromBottom(fn) }
func (f fromBottom) Enter(node ast.Node) (visit, callExit bool) { return true, false }
func (f fromBottom) Exit(node ast.Node)                         {}
func (f fromBottom) FromTop(node ast.Node) (visit bool)         { return true }
func (f fromBottom) FromBottom(node ast.Node)                   { f(node) }

func (r *rulesVisitor) Enter(node ast.Node) (visit, callExit bool) {
	visit = true
	callExit = false
	for _, v := range r.visitors {
		ok, withExit := v.Enter(node)
		if visit {
			visit = ok
		}
		callExit = callExit || withExit
	}
	return visit, callExit
}

func (r *rulesVisitor) Exit(node ast.Node) {
	for _, v := range r.visitors {
		v.Exit(node)
	}
}

func (r *rulesVisitor) FromTop(node ast.Node) bool {
	visit := true
	for _, v := range r.visitors {
		ok := v.FromTop(node)
		if visit {
			visit = ok
		}
	}
	return visit
}

func (r *rulesVisitor) FromBottom(node ast.Node) {
	for _, v := range r.visitors {
		v.FromBottom(node)
	}
}
