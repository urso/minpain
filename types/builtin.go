package types

type Primitive interface {
	Type
}

type primType struct {
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

	Bool Primitive = &primType{bits: 1}

	Numeric Primitive = &primType{bits: 0xff}
	Decimal Primitive = &primType{bits: 0xff, extends: Numeric}
	Byte    Primitive = &primType{extends: Numeric, bits: 8}
	Short   Primitive = &primType{extends: Numeric, bits: 16}
	Int     Primitive = &primType{extends: Numeric, bits: 32}
	Long    Primitive = &primType{extends: Numeric, bits: 64}
	Float   Primitive = &primType{extends: Decimal, bits: 32}
	Double  Primitive = &primType{extends: Decimal, bits: 64}

	String Type = &stringType{isRegex: false}
	Regex  Type = &stringType{isRegex: true}
)

func (_ *voidType) Extends() Type {
	return nil
}

func (_ *objectType) Extends() Type {
	return nil
}

func (_ *defType) Extends() Type {
	return Object
}

func (p *primType) Extends() Type {
	return p.extends
}

func (s *stringType) Extends() Type {
	return Def
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
