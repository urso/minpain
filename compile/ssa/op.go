package ssa

import "fmt"

type Op uint

type opInfo struct {
	Name string
}

const (
	OpInvalid Op = iota

	// generic
	OpUndef // noop value
	OpArg   // mark value as argument access
	OpMov
	OpJmp
	OpPhi

	// arithmetic
	OpMul
	OpDiv
	OpRem
	OpAdd
	OpSub
	OpInc
	OpDec
	OpPos
	OpNeg

	// bit ops
	OpNot
	OpAnd
	OpOr
	OpXor
	OpLSH
	OpRSH
	OpUSH

	// comparison ops
	OpCmpEQ
	OpCmpEQR
	OpCmpNE
	OpCmpNER
	OpCmpLT
	OpCmpLTE
	OpCmpGT
	OpCmpGTE
)

func (op Op) String() string {
	info := op.info()
	if info == nil {
		return fmt.Sprintf("<unknown op %v>", uint(op))
	}
	return info.Name
}

func (op Op) info() *opInfo {
	i := int(op)
	if i < 0 || i >= len(opTable) {
		return nil
	}
	return &opTable[i]
}

var opTable = []opInfo{
	{Name: "Invalid"},

	// generic
	{Name: "Undef"},
	{Name: "Arg"},
	{Name: "Mov"},
	{Name: "Jmp"},
	{Name: "Phi"},

	// arithmetic
	{Name: "Mul"},
	{Name: "Div"},
	{Name: "Rem"},
	{Name: "Add"},
	{Name: "Sub"},
	{Name: "Inc"},
	{Name: "Dec"},
	{Name: "Pos"},
	{Name: "Neg"},

	// bit ops
	{Name: "Not"},
	{Name: "And"},
	{Name: "Or"},
	{Name: "Xor"},
	{Name: "LSH"},
	{Name: "RSH"},
	{Name: "USH"},

	// comparison ops
	{Name: "CmpEQ"},
	{Name: "CmpEQR"},
	{Name: "CmpNE"},
	{Name: "CmpNER"},
	{Name: "CmpLT"},
	{Name: "CmpLTE"},
	{Name: "CmpGT"},
	{Name: "CmpGTE"},
}
