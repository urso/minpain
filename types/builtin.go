package types

type Primitive interface {
	Type
}

type primType struct {
	name    string
	bits    uint8
	extends Type
}

type constType struct{}

type voidType struct{}

type objectType struct{}

type defType struct{}

type stringType struct {
	isRegex bool
}

var (
	Void      Type = &voidType{}
	Object    Type = &objectType{}
	Def       Type = &defType{}
	Null      Type = Def
	Exception Type = Def

	Bool Primitive = &primType{name: "boolean", bits: 1}

	Numeric Primitive = &primType{name: "numeric", bits: 0xff}
	Decimal Primitive = &primType{name: "decimal", bits: 0xff, extends: Numeric}
	Byte    Primitive = &primType{name: "byte", extends: Numeric, bits: 8}
	Short   Primitive = &primType{name: "short", extends: Numeric, bits: 16}
	Int     Primitive = &primType{name: "int", extends: Numeric, bits: 32}
	Long    Primitive = &primType{name: "long", extends: Numeric, bits: 64}
	Float   Primitive = &primType{name: "float", extends: Decimal, bits: 32}
	Double  Primitive = &primType{name: "double", extends: Decimal, bits: 64}

	String Type = &stringType{isRegex: false}
	Regex  Type = &stringType{isRegex: true}
)

func (_ *voidType) String() string { return "void" }
func (_ *voidType) Extends() Type  { return nil }

func (_ *objectType) String() string { return "Object" }
func (_ *objectType) Extends() Type  { return nil }

func (_ *defType) String() string { return "def" }
func (_ *defType) Extends() Type  { return Object }

func (p *primType) String() string { return p.name }
func (p *primType) Extends() Type  { return p.extends }

func (s *stringType) Extends() Type { return Def }
func (s *stringType) String() string {
	if s.isRegex {
		return "Regex"
	}
	return "String"
}

func IsInteger(t Type) bool {
	return IsNumeric(t) && !IsDecimal(t)
}

func IsDecimal(t Type) bool {
	return checkExtends(isType(Decimal), t)
}

func IsNumeric(t Type) bool {
	return checkExtends(isType(Numeric), t)
}

func isType(t Type) func(Type) bool {
	return func(in Type) bool {
		return t == in
	}
}

func checkExtends(pred func(Type) bool, in Type) bool {
	if pred(in) {
		return true
	}
	if e := in.Extends(); e != nil {
		return checkExtends(pred, e)
	}
	return false
}
