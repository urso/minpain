package ssa

import (
	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/constant"
	"github.com/urso/minpain/types"
)

type Program struct {
	Funcs map[string]*Func

	Main *Func
}

type Func struct {
	ID   ID
	Name string
	Pos  ast.Pos
	Type types.Signature

	Entry  *Block
	Blocks []*Block

	blockIDs idAlloc
	valueIDs idAlloc
}

type Block struct {
	ID     ID
	Parent *Func
	Pos    ast.Pos

	Preds []Edge
	Succs []Edge

	Instr []Instruction
}

type Edge struct {
	b *Block
	i int // index in Succs/Preds package of other block
}

type Instruction interface {
	instNode()
}

type instr struct {
	block *Block
}

type Value struct { // ssa value assignment
	instr

	ID  ID
	Pos Pos

	Op   Op
	Type types.Type
	Args []Instruction
}

type Const struct { // constant value
	instr

	ID  ID
	Pos Pos

	Value constant.Value
	Type  types.Type
}

func (_ *instr) instNode()     {}
func (i *instr) Block() *Block { return i.block }
func (i *instr) Parent() *Func { return i.block.Parent }

func (fn *Func) newBlock() *Block {
	blk := &Block{
		ID:     fn.blockIDs.next(),
		Parent: fn,
	}
	fn.Blocks = append(fn.Blocks, blk)
	return blk
}

func (b *Block) newValue(op Op, t types.Type, args ...Instruction) *Value {
	val := &Value{
		ID:   b.Parent.valueIDs.next(),
		Op:   op,
		Type: t,
		Args: args,
	}

	val.block = b
	b.Instr = append(b.Instr, val)
	return val
}

func (b *Block) NumPredecessors() int {
	return len(b.Preds)
}

func (b *Block) Predecssor(i int) *Block {
	return b.Preds[i].b
}

func (b *Block) newPhi(t types.Type, args ...Instruction) *Value {
	val := &Value{
		ID:   b.Parent.valueIDs.next(),
		Op:   OpPhi,
		Type: t,
		Args: args,
	}
	val.block = b
	b.Instr = append([]Instruction{val}, b.Instr...)
	return val
}

func (v *Value) AppendArg(arg Instruction) {
	v.Args = append(v.Args, arg)
}

func (v *Value) setPos(name string, p ast.Pos) {
	v.Pos = Pos{
		Symbol: name,
		Pos:    p,
	}
}
