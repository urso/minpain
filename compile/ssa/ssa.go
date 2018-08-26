package ssa

import (
	"io"

	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/constant"
	"github.com/urso/minpain/types"
)

type Unit interface {
	genValueID() ID
	genBlockID() ID
}

type Program struct {
	Funcs map[string]*Func

	Entry  *Block
	Blocks []*Block

	blockIDs idAlloc
	valueIDs idAlloc
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
	Parent Unit
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
	HasID() bool
	GetID() ID
	Format(io.Writer) error
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
	Pos ast.Pos

	Value constant.Value
	Type  types.Type
}

func (_ *instr) instNode()     {}
func (i *instr) Block() *Block { return i.block }
func (i *instr) Parent() Unit  { return i.block.Parent }

func (_ *Value) HasID() bool { return true }
func (v *Value) GetID() ID   { return v.ID }

func (_ *Const) HasID() bool { return true }
func (c *Const) GetID() ID   { return c.ID }

func (prog *Program) genBlockID() ID {
	return prog.blockIDs.next()
}

func (prog *Program) genValueID() ID {
	return prog.valueIDs.next()
}

func (prog *Program) newBlock() *Block {
	blk := &Block{
		ID:     prog.genBlockID(),
		Parent: prog,
	}
	prog.Blocks = append(prog.Blocks, blk)
	return blk
}

func (fn *Func) newBlock() *Block {
	blk := &Block{
		ID:     fn.genBlockID(),
		Parent: fn,
	}
	fn.Blocks = append(fn.Blocks, blk)
	return blk
}

func (fn *Func) genBlockID() ID {
	return fn.blockIDs.next()
}

func (fn *Func) genValueID() ID {
	return fn.valueIDs.next()
}

func (b *Block) newValue(op Op, t types.Type, args ...Instruction) *Value {
	val := &Value{
		ID:   b.Parent.genValueID(),
		Op:   op,
		Type: t,
		Args: args,
	}

	val.block = b
	b.Instr = append(b.Instr, val)
	return val
}

func (b *Block) newConst(t types.Type, val constant.Value) *Const {
	c := &Const{
		ID:    b.Parent.genValueID(),
		Type:  t,
		Value: val,
	}

	c.block = b
	b.Instr = append(b.Instr, c)
	return c
}

func (b *Block) NumPredecessors() int {
	return len(b.Preds)
}

func (b *Block) Predecssor(i int) *Block {
	return b.Preds[i].b
}

func (b *Block) newPhi(t types.Type, args ...Instruction) *Value {
	val := &Value{
		ID:   b.Parent.genValueID(),
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
