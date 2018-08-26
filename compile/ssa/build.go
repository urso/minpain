package ssa

import (
	"fmt"
	"strconv"

	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/check"
	"github.com/urso/minpain/constant"
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

	// build functions
	for _, fn := range info.Functions {
		if err := buildFn(st, fn); err != nil {
			return nil, err
		}
	}

	// build main program
	funcs := make(map[string]*Func, len(st.funcs))
	for _, fn := range st.funcs {
		funcs[fn.Name] = fn
	}

	progam := &Program{Funcs: funcs}
	progam.Entry = progam.newBlock()

	st.fn = nil
	st.block = progam.Entry
	for _, node := range script.Program {
		stmt, ok := node.(ast.Stmt)
		if !ok {
			continue // filter out function definitions
		}

		if err := buildStmt(st, stmt); err != nil {
			return nil, err
		}
	}

	return progam, nil
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

	return buildStmt(st, decl.Body)
}

func buildStmts(st *bldState, stmts []ast.Stmt) error {
	for _, stmt := range stmts {
		if err := buildStmt(st, stmt); err != nil {
			return err
		}
	}
	return nil
}

func buildStmt(st *bldState, node ast.Stmt) error {
	switch stmt := node.(type) {
	case *ast.BlockStmt:
		return buildStmts(st, stmt.Stmts)

	case *ast.Try:
		panic("TODO")

	case *ast.IfStmt:
		panic("TODO: build if stmt")

	case *ast.BranchStmt:
		panic("TODO: build break/continue stmt")

	case *ast.ReturnStmt:
		panic("TODO: build return stmt")

	case *ast.ThrowStmt:
		panic("TOOD: build throw stmt")

	case *ast.Loop:
		panic("TODO: build loop stmt")

	case *ast.EachLoop:
		panic("TODO: build eachloop stmt")

	case *ast.Decl:
		panic("TODO: init declararation")

	case ast.Expr:
		_, err := buildExpr(st, stmt)
		return err

	default:
		panic(fmt.Errorf("unhandled stmt: %v", stmt))
	}
}

func buildExpr(st *bldState, node ast.Expr) (Instruction, error) {
	switch expr := node.(type) {
	case *ast.ListInit:
		panic("TODO: list init")

	case *ast.MapInit:
		panic("TODO: map init")

	case *ast.Ident: // variable access
		panic("TODO")

	case *ast.Assign:
		panic("TODO: assign value")

	case *ast.CastExpr:
		panic("TODO: cast expr")

	case *ast.CondExpr:
		panic("TODO: cond expr")

	case *ast.UnaryOp:
		panic("TODO: unary op")

	case *ast.BinOp:
		panic("TODO: binop")

	case *ast.Literal:
		return buildLit(st, expr)

	case *ast.Call:
		panic("TODO: call")

	case *ast.FieldAccess:
		panic("TODO: field access")

	case *ast.IdxAccess:
		panic("TODO: index access")

	case *ast.Access:
		panic("TODO: access")

	case *ast.NewArray:
		panic("TODO: new array")

	case *ast.Lambda:
		panic("TODO: lambda")

	case *ast.Ref:
		panic("TODO: ref")

	default:
		panic(fmt.Errorf("unhandled expression type: %v", node))
	}
}

func buildLit(st *bldState, lit *ast.Literal) (Instruction, error) {
	var val constant.Value

	switch lit.Kind {
	case ast.LitBool:
		val = constant.Boolean(lit.Value == "true")

	case ast.LitInt:
		i, err := strconv.ParseInt(lit.Value, 10, 64)
		if err != nil {
			return nil, err
		}
		val = constant.Number(i)

	case ast.LitString:
		val = constant.Str(lit.Value)

	default:
		panic(fmt.Errorf("unhandled literal type: %v", lit.Kind))
	}

	t := st.info.Types.Expected[lit]
	if t == nil {
		t = st.info.Types.Actual[lit]
	}

	c := st.newConst(st.block, t, val)
	c.Pos = lit.Pos()
	return c, nil
}

func (st *bldState) newValue(b *Block, op Op, t types.Type, args ...Instruction) *Value {
	v := b.newValue(op, t, args...)
	st.indexPhiUsers(v)
	return v
}

func (st *bldState) newConst(b *Block, t types.Type, val constant.Value) *Const {
	return b.newConst(t, val)
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
