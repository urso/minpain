package check

import (
	"fmt"

	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/types"
)

type checkCtx struct {
	info  *Info
	scope *Scope

	// active function
	function *Function

	errs *multiErr
}

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

	indexScript(&idxCtx{
		info:  info,
		scope: scriptScope,
		errs:  errs,
		types: types,
	}, script)

	ctx := &checkCtx{info: info, scope: scriptScope, errs: errs}
	for _, node := range script.Program {
		checkNode(ctx, node)
	}
}

func checkNode(ctx *checkCtx, node ast.Node) Type {
	switch n := node.(type) {
	case ast.DeclExpr:
		return checkDecl(ctx, n)

	// some expressions can be statements too, check expressions first
	case ast.Expr:
		return checkExpr(ctx, n)

	case ast.Stmt:
		return checkStmt(ctx, n)

	default:
		panic("unexpected node type")
	}
}

func checkDecl(ctx *checkCtx, node ast.DeclExpr) Type {
	switch decl := node.(type) {
	case *ast.Decl:
		return checkVarDecl(ctx, decl)
	case *ast.FuncDecl:
		return checkFunc(ctx, decl)
	default:
		panic(fmt.Sprint("invalid decl node: ", node))
	}
}

func checkVarDecl(ctx *checkCtx, node *ast.Decl) Type {
	var reportType Type

	for _, expr := range node.Vars {
		switch decl := expr.(type) {
		case *ast.Ident:
			reportType = ctx.info.Decl[decl].Type()

		case *ast.Assign:
			id := decl.LHS.(*ast.Ident)
			reportType = ctx.info.Decl[id]
			checkTypes(reportType, checkNode(ctx, decl.RHS))

		default:
			panic(fmt.Sprint("invalid variable declaration node: ", node))
		}
	}

	return reportType
}

func checkFunc(ctx *checkCtx, node *ast.FuncDecl) Type {
	panic("TODO")
}

func checkStmt(ctx *checkCtx, node ast.Stmt) Type {
	switch stmt := node.(type) {
	case *ast.BlockStmt:
		panic("TODO")

	case *ast.Try:
		panic("TODO")

	case *ast.IfStmt:
		checkTypes(types.Bool, checkExpr(ctx, stmt.Cond))
		checkNode(ctx, stmt.Then)
		if stmt.Else != nil {
			checkNode(ctx, stmt.Else)
		}
		return types.Void

	case *ast.Loop:
		panic("TODO")

	case *ast.EachLoop:
		panic("TODO")

	case *ast.ThrowStmt:
		return checkExpr(ctx, stmt.Expr)

	case *ast.BranchStmt:
		// nothing
		return types.Void

	case *ast.ReturnStmt:
		var retType Type
		if stmt.Expr == nil {
			retType = types.Void
		} else {
			retType = checkExpr(ctx, stmt.Expr)
		}

		if ctx.function == nil {
			return retType
		}
		return checkTypes(ctx.function.Return(), retType)

	default:
		panic(fmt.Sprint("unexpected statement node:", node))
	}
}

func checkExpr(ctx *checkCtx, node ast.Expr) Type {
	switch node.(type) {
	case *ast.Ident:
		panic("TODO")

	case *ast.ListInit:
		panic("TODO")

	case *ast.MapInit:
		panic("TODO")

	case *ast.Assign:
		panic("TODO")

	case *ast.CastExpr:
		panic("TODO")

	case *ast.CondExpr:
		panic("TODO")

	case *ast.UnaryOp:
		panic("TODO")

	case *ast.BinOp:
		panic("TODO")

	case *ast.Literal:
		panic("TODO")

	case *ast.Call:
		panic("TODO")

	case *ast.FieldAccess:
		panic("TODO")

	case *ast.IdxAccess:
		panic("TODO")

	case *ast.Access:
		panic("TODO")

	case *ast.ArrType:
		panic("TODO")

	case *ast.NewArray:
		panic("TODO")

	case *ast.Lambda:
		panic("TODO")

	case *ast.Ref:
		panic("TODO")

	default:
		panic(fmt.Sprint("invalid decl node: ", node))
	}
}

func checkTypes(t1, t2 Type) Type {
	panic("TODO")
}
