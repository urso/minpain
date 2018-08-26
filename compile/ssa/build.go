package ssa

import (
	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/check"
	"github.com/urso/minpain/types"
)

// Build constructs the SSA directly from the AST according to [1].
// The bldState type holds the local build state, as is required the
// construction algorithm. Field namings and methods follow the papers namings
// as much as possible, so it becomes easier to compare the implementation with
// the original paper.
//
// [1]: Simple and Efficient Construction of Static Single Assignment Form by M.Braun et. al

type bldState struct {
	info *check.Info

	ids idAlloc

	// active function/block
	fn    *Func
	block *Block

	// list of all generated functions
	funcs []*Func

	// variable construction/lookup related fields:
	// --------------------------------------------

	// variable mapping for identifier to current block local SSA value
	currentDef     map[*ast.Ident]map[*Block]*Value
	sealedBlocks   map[*Block]struct{}
	incompletePhis map[*Block]map[*ast.Ident]*Value
	phiUsers       map[*Value][]*Value
}

// Build constructs the SSA directly from the AST according to [1].
//
func Build(info *check.Info, script *ast.Script) (*Program, error) {
	st := &bldState{
		info:           info,
		currentDef:     map[*ast.Ident]map[*Block]*Value{},
		sealedBlocks:   map[*Block]struct{}{},
		incompletePhis: map[*Block]map[*ast.Ident]*Value{},
		phiUsers:       map[*Value][]*Value{},
	}

	for _, fn := range info.Functions {
		if err := buildFn(st, fn); err != nil {
			return nil, err
		}
	}

	panic("TODO")
}

func buildFn(st *bldState, info *check.Function) error {
	fn := st.newFunc(info)
	entry := fn.newBlock()
	fn.Entry = entry

	// prepare build state for current function
	st.fn = fn
	st.block = entry

	// add function parameters to entry block variable list
	decl := info.Declaration()
	sig := info.Signature()
	for i, arg := range decl.Args {
		t := sig.Argument(i)

		value := st.newValue(entry, OpAdd, t)
		value.setPos(arg.ID.Name, arg.FieldPos)
		st.writeVar(arg.ID, entry, value)
	}

	panic("TODO")
}

func (st *bldState) newValue(b *Block, op Op, t types.Type, args ...Instruction) *Value {
	v := b.newValue(op, t, args...)
	st.indexPhiUsers(v)
	return v
}

func (st *bldState) newPhi(b *Block, t types.Type, args ...Instruction) *Value {
	phi := b.newPhi(t, args...)
	st.indexPhiUsers(phi)
	return phi
}

func (st *bldState) indexPhiUsers(v *Value) {
	args := v.Args
	for _, a := range args {
		arg, ok := a.(*Value)
		if !ok {
			continue
		}
		if arg.Op != OpPhi {
			continue
		}

		st.phiUsers[arg] = append(st.phiUsers[arg], v)
	}
}

func (st *bldState) newFunc(info *check.Function) *Func {
	fn := &Func{
		ID:   st.ids.next(),
		Name: info.Name(),
		Pos:  info.Pos(),
		Type: info.Signature(),
	}

	st.funcs = append(st.funcs, fn)
	return fn
}

func (st *bldState) writeVar(variable *ast.Ident, blk *Block, val *Value) {
	m := st.currentDef[variable]
	if m == nil {
		m = map[*Block]*Value{}
		st.currentDef[variable] = m
	}
	m[blk] = val
}

func (st *bldState) readVar(variable *ast.Ident, blk *Block) *Value {
	blocks := st.currentDef[variable]
	if blocks != nil {
		if val := blocks[blk]; val != nil {
			return val
		}
	}

	return st.readVarRecursive(variable, blk)
}

func (st *bldState) readVarRecursive(variable *ast.Ident, blk *Block) (val *Value) {
	if !st.isBlockSealed(blk) {
		// incomplete CFG
		val = st.newPhi(blk, nil)
		st.addIncompletePhi(variable, blk, val)
	} else if blk.NumPredecessors() == 1 {
		// If block has only one predecessor, no phi is needed
		val = st.readVar(variable, blk.Predecssor(0))
	} else {
		// Blocks is already sealed, but variable is unknown:
		//   -> break potential cycle with operandless phi
		val = st.newPhi(blk, nil)
		st.writeVar(variable, blk, val)        // write phi variable, so to break recursion on visited nodes
		val = st.addPhiOperands(variable, val) // <- recurse into predecessor blocks
	}

	st.writeVar(variable, blk, val)
	return val
}

func (st *bldState) addPhiOperands(variable *ast.Ident, phi *Value) *Value {
	blk := phi.Block()
	for i := 0; i < blk.NumPredecessors(); i++ {
		pred := blk.Predecssor(i)
		val := st.readVar(variable, pred) // recursively read variable from predecessors
		phi.AppendArg(val)
	}

	// check if we can cleanup block after all phi instructions have been generated
	return st.tryRemoveTrivialPhi(phi)
}

func (st *bldState) tryRemoveTrivialPhi(phi *Value) *Value {
	trivial, orig := isTrivialPhi(phi)
	if !trivial {
		return phi
	}

	// Phi instruction is unreachable or in function entry block if 'orig == nil'
	// Otherwise orig is the only unique variable used by phi.
	if orig == nil {
		orig = &Value{
			ID: phi.Block().Parent.valueIDs.next(),
			Op: OpUndef,
		}
	}

	// replace phi with orig
	users := st.phiUsers[phi]
	delete(st.phiUsers, phi)

	for _, user := range users {
		if user == phi {
			continue
		}

		for i, arg := range user.Args {
			if arg == phi {
				user.Args[i] = orig
			}
		}
	}

	// phi instructions using the current phi might have become trivial as well
	// -> recursively try to remove these phi instructions as well
	for _, user := range users {
		if user.Op != OpPhi {
			continue
		}

		st.tryRemoveTrivialPhi(user)
	}

	return orig
}

// isTrivialPhi checks if phi is trivial. Phi is trivial iff it references itself and one
// other value v any number of times.
func isTrivialPhi(phi *Value) (bool, *Value) {
	var other *Value
	for _, arg := range phi.Args {
		if arg == phi {
			continue // self-reference
		}
		if arg == other {
			continue // repeated variable reference
		}

		// phi merges at least 2 unique values -> not trivial
		return false, nil
	}

	return true, other
}

func (st *bldState) isBlockSealed(b *Block) bool {
	_, exists := st.sealedBlocks[b]
	return exists
}

func (st *bldState) sealBlock(b *Block) {
	if st.isBlockSealed(b) {
		return
	}

	phis := st.incompletePhis[b]
	delete(st.incompletePhis, b)
	for id, phi := range phis {
		st.addPhiOperands(id, phi)
	}

	st.sealedBlocks[b] = struct{}{}
}

func (st *bldState) addIncompletePhi(variable *ast.Ident, b *Block, phi *Value) {
	blk := st.incompletePhis[b]
	if blk == nil {
		blk = map[*ast.Ident]*Value{}
		st.incompletePhis[b] = blk
	}
	blk[variable] = phi
}
