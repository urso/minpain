package constant

import "fmt"

type Value interface {
	Kind() Kind
	String() string
}

type Kind uint

//go:generate stringer -type=Kind
const (
	Unknown Kind = iota
	Bool
	Numeric
	Decimal
	String
)

type unknownVal struct {
}

type boolVal struct {
	b bool
}

type numVal struct {
	i int64
}

type decVal struct {
	f float64
}

type strVal struct {
	s string
}

func (_ *unknownVal) Kind() Kind     { return Unknown }
func (_ *unknownVal) String() string { return "<unknown>" }

func (v *boolVal) Kind() Kind { return Bool }
func (v *boolVal) String() string {
	if v.b {
		return "true"
	}
	return "false"
}

func (v *numVal) Kind() Kind     { return Numeric }
func (v *numVal) String() string { return fmt.Sprintf("%v", v.i) }

func (v *decVal) Kind() Kind     { return Decimal }
func (v *decVal) String() string { return fmt.Sprintf("%v", v.f) }

func (v *strVal) Kind() Kind     { return String }
func (v *strVal) String() string { return fmt.Sprintf("%q", v.s) }
