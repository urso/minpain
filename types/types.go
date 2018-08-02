package types

import "fmt"

type Type interface {
	String() string
	Extends() Type
	InstanceOf(t Type) bool
	ComponentType() Type
}

type Class interface {
	Type // a class is a type

	NumConstructors() int
	Constructor(i int) Constructor

	NumFields() int
	Field(i int) Field

	NumMethods() int
	Method(i int) Method
}

type Interface interface {
	Type // an interface is a type

	NumFields() int
	Field(i int) Field

	NumMethods() int
	Method(i int) Method
}

type Constructor interface {
	Type() Signature
}

// Signature of functions/methods.
type Signature interface {
	Type // a signature is a type

	Return() Type
	NumArguments() int
	Argument(int) Type
}

type Field interface {
	Name() string
	IsStatic() bool
	Type() Type
}

type Method interface {
	Name() string
	IsStatic() bool
	Type() Signature
}

type Array struct {
	Elt Type
}

func (a *Array) Extends() Type       { return Def }
func (a *Array) String() string      { return fmt.Sprintf("%v[]", a.Elt) }
func (a *Array) ComponentType() Type { return a.Elt }
func (a *Array) InstanceOf(t Type) bool {
	if other, ok := t.(*Array); ok {
		return a.Elt.InstanceOf(other.Elt)
	}
	return t == Def || t == Object
}
