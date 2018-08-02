package check

import (
	"fmt"

	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/ast/walk"
	"github.com/urso/minpain/types"
)

type checkCtx struct {
	info  *Info
	scope *Scope
	errs  multiErr

	trace func(string, ...interface{})

	function *Function // active function
}

func (ctx *checkCtx) recordType(n ast.Node, t Type) {
	ctx.info.Types.Actual[n] = t
}

func (ctx *checkCtx) recordExpectedType(n ast.Node, t Type) {
	ctx.info.Types.Expected[n] = t
}

func (ctx *checkCtx) typeOf(node ast.Node) Type {
	return ctx.info.Types.Actual[node]
}

func (ctx *checkCtx) recordErr(err error) { ctx.errs.Add(err) }

func (ctx *checkCtx) withScope(scope *Scope, fn func()) {
	old := ctx.scope
	ctx.scope = old
	defer func() {
		ctx.scope = old
	}()
	fn()
}

func Check(errs multiErr, types TypeTable, tree *ast.Script, debug bool) (*Info, error) {
	Validate(errs, tree)
	if err := errs.Err(); err != nil {
		return nil, err
	}

	info := NewInfo(types)
	scriptScope := Index(errs, tree.Pos().Source, info, tree)
	info.RootScope = scriptScope
	if err := errs.Err(); err != nil {
		return nil, err
	}

	Types(errs, info, scriptScope, tree, debug)
	if err := errs.Err(); err != nil {
		return nil, err
	}

	return info, nil
}

func Types(errs multiErr, info *Info, scope *Scope, script *ast.Script, debug bool) {
	checkScript(errs, info, scope, script, debug)
}

func checkScript(errs multiErr, info *Info, scope *Scope, script *ast.Script, debug bool) {
	source := script.Pos().Source
	if source == "" {
		source = "<script>"
	}

	ctx := &checkCtx{info: info, scope: scope, errs: errs}
	ctx.trace = func(_ string, _ ...interface{}) {}
	if debug {
		ctx.trace = func(s string, vs ...interface{}) {
			fmt.Println(fmt.Sprintf(s, vs...))
		}
	}

	activeFuncEnv := walk.EnvSetupWith(
		func(n ast.Node) (ok bool) {
			if decl, ok := n.(*ast.FuncDecl); ok {
				obj := info.Functions[decl]
				if ok = obj != nil; ok {
					ctx.function = obj
				}
			}
			return
		},
		func(_ ast.Node) { ctx.function = nil },
	)

	walk.Walk(script, walk.Rules(
		activeFuncEnv,
		scopeEnv(info, &ctx.scope),
		walk.FromBottom(func(node ast.Node) {
			checkNode(ctx, node)
		}),
	))
}

func checkNode(ctx *checkCtx, node ast.Node) {
	switch n := node.(type) {
	// check expressions:
	case *ast.Ident:
		if n.Obj.Kind == ast.TypObj {
			break
		}

		var typ Type
		if obj := ctx.info.Used[n]; obj != nil {
			typ = obj.Type()
		} else if obj := ctx.info.Decl[n]; obj != nil {
			typ = obj.Type()
		}

		if typ == nil {
			ctx.trace("check: can not resolve id type: %v (%v)", n.Name, n.Pos())
			break
		}

		ctx.trace("check: resolve id type: %v (%v) -> %v", n.Name, n.Pos(), typ)
		ctx.recordType(n, typ)

	case *ast.ListInit:
		ctx.recordErr(newNodeErrorf(n, "TODO: define type for list"))

	case *ast.MapInit:
		ctx.recordErr(newNodeErrorf(n, "TODO: define type for map"))

	case *ast.Assign:
		ctx.recordType(n, checkAssign(ctx, n))

	case *ast.CastExpr:
		ctx.recordErr(newNodeErrorf(n, "TODO: add support for cast expressions"))

	case *ast.CondExpr:
		ctx.recordErr(newNodeErrorf(n, "TODO: add support for cond expressions"))

	case *ast.UnaryOp:
		ctx.recordType(n, ctx.typeOf(n.X))

	case *ast.BinOp:
		ctx.recordType(n, checkBinop(ctx, n, n.Op, n.X, n.Y))

	case *ast.Literal:
		checkLit(ctx, n)

	case *ast.Call:
		ctx.recordType(n, checkCall(ctx, n))

	case *ast.NewArray:
		ctx.recordErr(newNodeErrorf(n, "TODO: add support for arrays"))

	case *ast.FieldAccess:
		ctx.recordErr(newNodeErrorf(n, "TODO: add support for field access"))

	case *ast.IdxAccess:
		ctx.recordErr(newNodeErrorf(n, "TODO: add support for idx access"))

	case *ast.Access:
		ctx.recordErr(newNodeErrorf(n, "TODO: add support for map/array access"))

	case *ast.Ref:
		ctx.recordErr(newNodeErrorf(n, "TODO: add support for method"))

	case *ast.Lambda:
		ctx.recordErr(newNodeErrorf(n, "TODO: add support for lambda"))

	// check statements
	case *ast.Trap: // check type can be an Exception
		ctx.recordErr(newNodeErrorf(n, "TODO: define required types for exceptions"))

	case *ast.IfStmt:
		checkTypes(ctx, n.Cond, types.Bool, ctx.typeOf(n.Cond))
		ctx.recordExpectedType(n.Cond, types.Bool)

	case *ast.ReturnStmt:
		retType := types.Void
		if n.Expr != nil {
			retType = ctx.typeOf(n.Expr)
		}

		ctx.recordType(n, retType)
		if fn := ctx.function; fn != nil {
			fnRet := fn.Return()
			checkTypes(ctx, n, fnRet, retType)
			ctx.recordExpectedType(n, fnRet)
		}

	case *ast.ThrowStmt:
		throwType := ctx.typeOf(n.Expr)
		ctx.recordType(n, checkTypes(ctx, n, types.Exception, throwType))

	case *ast.Loop:
		if n.Cond != nil {
			checkTypes(ctx, n.Cond, types.Bool, ctx.typeOf(n.Cond))
			ctx.recordExpectedType(n.Cond, types.Bool)
		}

	case *ast.EachLoop:
		declVar := ctx.info.Decl[n.ID]
		declType := declVar.Type()

		exprType := ctx.typeOf(n.Expr)
		componentType := exprType.ComponentType()
		if componentType == nil {
			ctx.recordErr(newNodeErrorf(n.Expr, "Type '%v' is not iterable", exprType))
		} else {
			checkTypes(ctx, n.Expr, declType, componentType)
		}
		ctx.recordExpectedType(n.Expr, declType)
	}
}

func checkLit(ctx *checkCtx, lit *ast.Literal) (t Type) {
	litTypes := map[ast.LiteralKind]Type{
		ast.LitBool:    types.Bool,
		ast.LitNull:    types.Null,
		ast.LitString:  types.String,
		ast.LitRegex:   types.Regex,
		ast.LitInt:     types.Numeric,
		ast.LitDecimal: types.Decimal,
	}

	if t, ok := litTypes[lit.Kind]; ok {
		ctx.trace("literal %v(%v) -> %v", lit.Value, lit.Pos(), t)
		ctx.recordType(lit, t)
		return t
	}

	err := newNodeErrorf(lit, "bad literal")
	ctx.trace("%v", err)
	ctx.recordErr(err)
	return types.Void
}

func checkBinop(ctx *checkCtx, at ast.Node, op ast.OpKind, x, y ast.Expr) Type {
	switch op & ast.OPClassMask {
	case ast.OpClassArithmetic:
		return checkArithBinop(ctx, at, op, x, y)

	case ast.OpClassRegex:
		checkTypes(ctx, x, types.String, ctx.typeOf(x))
		checkTypes(ctx, y, types.Regex, ctx.typeOf(y))
		return types.Bool

	case ast.OpClassBits:
		ctx.recordErr(newNodeErrorf(at, "TODO: add support for bit ops"))
		return types.Bool

	case ast.OpClassBool:
		checkTypes(ctx, x, types.Bool, ctx.typeOf(x))
		checkTypes(ctx, y, types.Bool, ctx.typeOf(y))
		return types.Bool

	case ast.OpClassCmpEq, ast.OpClassCmpNum:
		return checkCmpBinop(ctx, at, op, x, y)

	case ast.OpClassOther:
		switch op {
		case ast.OpInstanceOf:
			ctx.recordErr(newNodeErrorf(at, "TODO: add support for instanceof"))
			return types.Bool
		default:
			panic("invalid op type")
		}

	default:
		panic("invalid op type")
	}
}

func checkArithBinop(ctx *checkCtx, at ast.Node, op ast.OpKind, x, y ast.Expr) Type {
	return checkPromotingBinop(ctx, at, op, x, y)
}

func checkCmpBinop(ctx *checkCtx, at ast.Node, op ast.OpKind, x, y ast.Expr) Type {
	// check operands and record expected types for comparison operator
	checkPromotingBinop(ctx, at, op, x, y)
	return types.Bool // but comparisons always return bool
}

func checkPromotingBinop(ctx *checkCtx, at ast.Node, op ast.OpKind, x, y ast.Expr) Type {
	var promotFn func(tx, ty types.Type) types.Type
	switch op & ast.OPClassMask {
	case ast.OpClassArithmetic:
		if op == ast.OpAdd {
			promotFn = types.PromoteAdd
		} else {
			promotFn = types.PromoteNumbers
		}
	case ast.OpClassCmpEq:
		promotFn = types.PromoteEq
	case ast.OpClassCmpNum:
		promotFn = types.PromoteNumbers
	}

	tx, ty := ctx.typeOf(x), ctx.typeOf(y)
	promoted := promotFn(tx, ty)
	if promoted == nil {
		ctx.recordErr(newNodeErrorf(at, "can not apply '%v' to values of type %v and %v", op.Symbol(), tx, ty))
		return types.Def // return some type, so checker can find more errors
	}

	ctx.recordExpectedType(x, promoted)
	ctx.recordExpectedType(y, promoted)
	return promoted
}

func checkCall(ctx *checkCtx, n *ast.Call) Type {
	if n.Value != nil {
		return checkInvoke(ctx, n)
	}

	obj := ctx.info.FunCalls[n]
	fn, ok := obj.(*Function)
	if !ok {
		ctx.recordErr(newNodeErrorf(n, "%v is no function", n.ID.Name))
		return types.Def // return def so type checking can continue with extra errors
	}

	sig := fn.Signature()

	nArgs := sig.NumArguments()
	if nArgs != len(n.Args) {
		ctx.recordErr(newNodeErrorf(n, "function %v expected %v arguments", fn.Name(), nArgs))

		if nArgs > len(n.Args) {
			nArgs = len(n.Args)
		}
	}

	ctx.trace("check call: %v", sig)
	for i := 0; i < nArgs; i++ {
		arg := n.Args[i]
		tSig := sig.Argument(i)
		ctx.trace("arg %v -> %v", arg, tSig)
		checkTypes(ctx, arg, tSig, ctx.typeOf(arg))
		ctx.recordExpectedType(arg, tSig)
	}

	return sig.Return()
}

func checkInvoke(ctx *checkCtx, n *ast.Call) Type {
	ctx.recordErr(newNodeErrorf(n, "TODO: add support for method invokations"))
	return types.Def
}

func checkAssign(ctx *checkCtx, n *ast.Assign) Type {
	if !ast.Storable(n.LHS) {
		// invalid AST, already caught by validator
		return types.Def
	}

	if n.Op != ast.OpAssign {
		return checkCompoundAssign(ctx, n)
	}

	// simple assignment
	res := ctx.typeOf(n.RHS)
	if res == types.Void {
		ctx.recordErr(newNodeError(n, "Right-hand side cannot be a [void] type for assignment."))
	}
	expected := checkTypes(ctx, n.RHS, ctx.typeOf(n.LHS), res) // check compatibility
	ctx.recordExpectedType(n.RHS, expected)
	return expected
}

func checkCompoundAssign(ctx *checkCtx, n *ast.Assign) Type {
	ctx.recordErr(newNodeError(n, "TODO: compound assignment"))
	return ctx.typeOf(n.LHS)
}

func checkTypes(ctx *checkCtx, at ast.Node, expected, t Type) Type {
	if t == nil {
		panic("type is <nil>")
	}
	if expected == nil {
		panic("expected type is nil")
	}

	ctx.trace("checkTypes: %v == %v", expected, t)

	if expected == t { // bool, string, regex, void must match by type
		return expected
	}

	if t == types.Void {
		ctx.recordErr(newNodeErrorf(at, "Did not expected 'void', but %v", expected))
		return types.Def // return 'def', so type checker can continue
	}

	if expected == types.Def {
		return expected // def always wins
	}

	if types.IsNumeric(expected) { // try to 'promote' common numeric types
		var promoted Type
		if types.IsNumeric(t) {
			promoted = types.PromoteNumbers(expected, t)
		}

		if promoted == nil {
			ctx.recordErr(newNodeErrorf(at, "type mismatch. Expected: %v, but did find: %v", expected, t))
		}
		return promoted
	}

	if t.InstanceOf(expected) {
		return expected
	}

	// ohoh, type mismatch
	ctx.recordErr(newNodeErrorf(at, "type mismatch. Expected: %v, but did find: %v", expected, t))
	return types.Def // return 'def', so type checker can continue finding more potential errors
}
