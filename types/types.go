package types

type Type interface {
	Extends() Type
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

func (t *Array) Extends() Type { return Def }
