package check

import (
	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/types"
)

type Info struct {
	Functions map[*ast.FuncDecl]*Function

	Scopes map[ast.Node]*Scope

	// link variable declarations to objects
	Decl map[*ast.Ident]Object

	// link variable usage to original object
	Used map[*ast.Ident]Object

	// record type for every expression
	Types map[ast.Expr]Type

	Literals []*ast.Literal
}

type (
	Object interface {
		Name() string
		Type() Type
	}

	object struct {
		parent *Scope
		name   string
		node   ast.Node
		typ    Type
	}

	Function struct {
		object
		scope *Scope
	}

	Variable struct {
		object
	}
)

func (o *object) Name() string {
	return o.name
}

func (o *object) Type() Type {
	return o.typ
}

func (o *object) Pos() ast.Pos {
	return o.node.Pos()
}

func newFuncObj(fn *ast.FuncDecl, typ Type, scope *Scope) *Function {
	return &Function{
		object: object{parent: scope.parent, name: fn.Name.Name, typ: typ, node: fn},
		scope:  scope,
	}
}

func (fn *Function) Signature() types.Signature {
	return fn.Type().(types.Signature)
}

func (fn *Function) Return() Type {
	return fn.Signature().Return()
}

func newVarObj(decl ast.Node, name string, typ Type, parent *Scope) *Variable {
	return &Variable{object{parent: parent, name: name, typ: typ, node: decl}}
}

func newParam(id *ast.Ident, typ Type, parent *Scope) *Variable {
	return &Variable{object{parent: parent, name: id.Name, typ: typ, node: id}}
}

func isVariable(obj Object) bool {
	_, ok := obj.(*Variable)
	return ok
}
