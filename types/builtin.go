package types

type Primitive interface {
	Type
}

type primType struct {
	bits  uint8
	float bool
}

type voidType struct{}

type defType struct{}

type stringType struct {
	isRegex bool
}

var (
	Void Type = &voidType{}
	Def  Type = &defType{}

	Bool   Primitive = &primType{bits: 1}
	Byte   Primitive = &primType{bits: 8}
	Short  Primitive = &primType{bits: 16}
	Int    Primitive = &primType{bits: 32}
	Long   Primitive = &primType{bits: 64}
	Float  Primitive = &primType{bits: 32, float: true}
	Double Primitive = &primType{bits: 64, float: true}

	String Type = &stringType{isRegex: false}
	Regex  Type = &stringType{isRegex: true}
)
