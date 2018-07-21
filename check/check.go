package check

import (
	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/ast/walk"
	"github.com/urso/minpain/types"
)

type checkCtx struct {
	info  *Info
	scope *Scope
	errs  *multiErr

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

func (ctx *checkCtx) recordErr(err error) { ctx.errs.add(err) }

func (ctx *checkCtx) withScope(scope *Scope, fn func()) {
	old := ctx.scope
	ctx.scope = old
	defer func() {
		ctx.scope = old
	}()
	fn()
}

func checkScript(errs *multiErr, types map[string]Type, script *ast.Script) {
	info := &Info{
		Functions: map[*ast.FuncDecl]*Function{},
		Scopes:    map[ast.Node]*Scope{},
		Decl:      map[*ast.Ident]Object{},
		Used:      map[*ast.Ident]Object{},
		Literals:  nil,
	}

	source := script.Pos().Source
	if source == "" {
		source = "<script>"
	}

	// TODO: support for parent scope
	scriptScope := NewScope(source, ScopeScript, nil)

	if !validateAST(errs, script) {
		return
	}

	indexScript(&idxCtx{
		info:  info,
		scope: scriptScope,
		errs:  errs,
		types: types,
	}, script)

	ctx := &checkCtx{info: info, scope: scriptScope, errs: errs}

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
		if obj := ctx.info.Used[n]; obj != nil {
			ctx.recordType(n, obj.Type())
		}

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
		checkTypes(ctx, types.Bool, ctx.typeOf(n.Cond))

	case *ast.ReturnStmt:
		retType := types.Void
		if n.Expr != nil {
			retType = ctx.typeOf(n.Expr)
		}

		if fn := ctx.function; fn != nil {
			checkTypes(ctx, fn.Return(), retType)
		}
		ctx.recordType(n, retType)

	case *ast.ThrowStmt:
		throwType := ctx.typeOf(n.Expr)
		checkTypes(ctx, types.Exception, throwType)
		ctx.recordType(n, throwType)

	case *ast.Loop:
		if n.Cond != nil {
			checkTypes(ctx, types.Bool, ctx.typeOf(n.Cond))
		}

	case *ast.EachLoop:
		declVar := ctx.info.Decl[n.ID]
		checkTypes(ctx, declVar.Type(), ctx.typeOf(n.Expr))
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
		ctx.recordType(lit, t)
		return t
	}

	ctx.recordErr(newNodeErrorf(lit, "bad literal"))
	return types.Void
}

func checkBinop(ctx *checkCtx, at ast.Node, op ast.OpKind, x, y ast.Expr) Type {
	switch op & ast.OPClassMask {
	case ast.OpClassArithmetic:
		return checkArithBinop(ctx, at, op, x, y)

	case ast.OpClassRegex:
		checkTypes(ctx, types.String, ctx.typeOf(x))
		checkTypes(ctx, types.Regex, ctx.typeOf(y))
		return types.Bool

	case ast.OpClassBits:
		ctx.recordErr(newNodeErrorf(at, "TODO: add support for bit ops"))
		return types.Bool

	case ast.OpClassBool:
		checkTypes(ctx, types.Bool, ctx.typeOf(x))
		checkTypes(ctx, types.Bool, ctx.typeOf(y))
		return types.Bool

	case ast.OpClassCmp:
		return checkCmpBinop(ctx, at, op, x, y)

	case ast.OpClassOther:
		switch op {
		case ast.OpAssign:
			return checkAssignType(ctx, ctx.typeOf(x), ctx.typeOf(y))
		case ast.OpInstanceOf:
			return types.Bool
		default:
			panic("invalid op type")
		}

	default:
		panic("invalid op type")
	}
}

func checkArithBinop(ctx *checkCtx, at ast.Node, op ast.OpKind, x, y ast.Expr) Type {
	var (
		sym      rune
		promotFn func(tx, ty types.Type) types.Type
	)

	switch op {
	case ast.OpMul:
		sym, promotFn = '*', types.PromoteNumbers
	case ast.OpDiv:
		sym, promotFn = '/', types.PromoteNumbers
	case ast.OpRem:
		sym, promotFn = '%', types.PromoteNumbers
	case ast.OpAdd:
		sym, promotFn = '+', types.PromoteAdd
	case ast.OpSub:
		sym, promotFn = '-', types.PromoteNumbers
	default:
		panic("invalid binary arithmetic operation")
	}

	tx, ty := ctx.typeOf(x), ctx.typeOf(y)
	promoted := promotFn(tx, ty)
	if promoted == nil {
		ctx.recordErr(newNodeErrorf(at, "can not apply '%v' to values of type %v and %v", sym, tx, ty))
		return types.Def // return some type, so checker can find more errors
	}

	ctx.recordExpectedType(x, promoted)
	ctx.recordExpectedType(y, promoted)
	return promoted
}

func checkCmpBinop(ctx *checkCtx, at ast.Node, op ast.OpKind, x, y ast.Expr) Type {
	panic("TODO")
}

func checkCall(ctx *checkCtx, n *ast.Call) Type {
	panic("TODO")
}

func checkAssign(ctx *checkCtx, n *ast.Assign) Type {
	panic("TODO")
}

func checkAssignType(ctx *checkCtx, to, from Type) Type {
	panic("TODO")
}

func checkTypes(ctx *checkCtx, expected, t Type) Type {
	panic("TODO")
}
