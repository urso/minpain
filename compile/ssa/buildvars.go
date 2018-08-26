package ssa

import "github.com/urso/minpain/ast"

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
			ID: phi.Block().Parent.genValueID(),
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

func (st *bldState) addIncompletePhi(variable *ast.Ident, b *Block, phi *Value) {
	blk := st.incompletePhis[b]
	if blk == nil {
		blk = map[*ast.Ident]*Value{}
		st.incompletePhis[b] = blk
	}
	blk[variable] = phi
}
