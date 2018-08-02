package check

import (
	"fmt"

	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/types"
)

type Info struct {
	Functions map[*ast.FuncDecl]*Function

	RootScope *Scope
	Scopes    map[ast.Node]*Scope

	// link variable declarations to objects
	Decl map[*ast.Ident]Object

	// link variable usage to original object
	Used map[*ast.Ident]Object

	FunCalls map[*ast.Call]Object

	// record types of expressions, calls and functions
	Types TypeMap

	// table of available types
	DeclTypes TypeTable

	Literals []*ast.Literal
}

type TypeMap struct {
	Actual   map[ast.Node]Type // actual type of a node
	Expected map[ast.Node]Type // expected type (e.g. on implicit type casts)
}

type (
	Object interface {
		Name() string
		Type() Type
		Parent() *Scope
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

func NewInfo(types TypeTable) *Info {
	return &Info{
		Functions: map[*ast.FuncDecl]*Function{},
		Scopes:    map[ast.Node]*Scope{},
		Decl:      map[*ast.Ident]Object{},
		Used:      map[*ast.Ident]Object{},
		FunCalls:  map[*ast.Call]Object{},
		Types: TypeMap{
			Actual:   map[ast.Node]Type{},
			Expected: map[ast.Node]Type{},
		},
		DeclTypes: types,
		Literals:  nil,
	}
}

func (o *object) Name() string {
	return o.name
}

func (o *object) Type() Type {
	return o.typ
}

func (o *object) Pos() ast.Pos {
	return o.node.Pos()
}

func (o *object) Parent() *Scope {
	return o.parent
}

func (o *object) String() string {
	return fmt.Sprintf("%v: %v", FullObjectName(o), o.Type())

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

func FullObjectName(obj Object) string {
	parent := obj.Parent()
	if parent == nil {
		return obj.Name()
	}
	return fmt.Sprintf("%v.%v", parent.FullName(), obj.Name())
}
