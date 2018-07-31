package check

import (
	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/ast/walk"
	"github.com/urso/minpain/types"
)

type idxCtx struct {
	info  *Info
	scope *Scope
	errs  multiErr

	types TypeTable
}

type TypeTable interface {
	Lookup(string) Type
}

func (ctx *idxCtx) openScriptScope(name string, n *ast.Script) *Scope {
	return ctx.openScope(name, ScopeScript, n)
}
func (ctx *idxCtx) openBlockScope(name string, n ast.Node) *Scope {
	return ctx.openScope(name, ScopeBlock, n)
}
func (ctx *idxCtx) openFuncScope(name string, n *ast.FuncDecl) *Scope {
	return ctx.openScope(name, ScopeFunction, n)
}
func (ctx *idxCtx) openLambdaScope(n *ast.Lambda) *Scope {
	return ctx.openScope("<lambda>", ScopeFunction, n)
}
func (ctx *idxCtx) openScope(name string, kind ScopeKind, node ast.Node) *Scope {
	scope := NewScope(name, kind, ctx.scope)
	ctx.scope = scope
	ctx.recordScope(node, scope)
	return scope
}

func (ctx *idxCtx) closeScope() {
	ctx.scope = ctx.scope.Parent()
}

func (ctx *idxCtx) activeScope() *Scope {
	return ctx.scope
}

func (ctx *idxCtx) withScope(scope *Scope, fn func()) {
	old := ctx.scope
	ctx.scope = scope
	defer func() {
		ctx.scope = old
	}()
	fn()
}

func (ctx *idxCtx) recordErr(err error) { ctx.errs.Add(err) }

func (ctx *idxCtx) recordRedecl(id *ast.Ident) {
	ctx.recordErr(newNodeErrorf(id, "%v redeclared", id.Name))
}

func (ctx *idxCtx) recordUnknownFunc(id *ast.Ident) {
	ctx.recordErr(newNodeErrorf(id, "unknown function %v", id.Name))
}

func (ctx *idxCtx) recordUnknownVar(id *ast.Ident) {
	ctx.recordErr(newNodeErrorf(id, "unknown variable %v", id.Name))
}

func (ctx *idxCtx) recordInvalidType(node ast.Node) {
	ctx.recordErr(newTypeErrorf(node, "no valid type expression"))
}

func (ctx *idxCtx) recordUnknownType(node ast.Node, name string) {
	ctx.recordErr(newTypeErrorf(node, "unknown type '%v'", name))
}

func (ctx *idxCtx) recordScope(node ast.Node, scope *Scope)     { ctx.info.Scopes[node] = scope }
func (ctx *idxCtx) recordFunc(node *ast.FuncDecl, fn *Function) { ctx.info.Functions[node] = fn }
func (ctx *idxCtx) recordDecl(node *ast.Ident, obj Object)      { ctx.info.Decl[node] = obj }
func (ctx *idxCtx) recordUse(node *ast.Ident, obj Object)       { ctx.info.Used[node] = obj }
func (ctx *idxCtx) recordLit(lit *ast.Literal)                  { ctx.info.Literals = append(ctx.info.Literals, lit) }
func (ctx *idxCtx) recordCall(call *ast.Call, obj Object)       { ctx.info.FunCalls[call] = obj }

func Index(errs multiErr, source string, info *Info, script *ast.Script) *Scope {
	topScope := NewScope("<top>", ScopeScript, nil)

	ctx := &idxCtx{
		info:  info,
		scope: topScope,
		errs:  errs,
		types: info.DeclTypes,
	}

	indexScript(ctx, script)
	return topScope
}

func indexScript(ctx *idxCtx, script *ast.Script) {
	// 1. phase: create and link block scopes
	buildScopes(ctx, script)

	// 2. phase: collect function names into the script scope
	//    As functions are defined before any other statements, but function can
	//    call each other, we have to collect the names into the scope before checking
	//    any variable access is the main script.
	indexFnNames(ctx, script)

	// index all variable declarations and uses
	indexVars(ctx, script)
}

func buildScopes(ctx *idxCtx, script *ast.Script) {
	walk.Walk(script, walk.Rules(
		walk.EnvSetupWith(
			func(node ast.Node) bool { return buildNodeScope(ctx, node) },
			func(_ ast.Node) { ctx.closeScope() },
		)))
}

func buildNodeScope(ctx *idxCtx, node ast.Node) bool {
	switch n := node.(type) {
	case *ast.Script:
		ctx.openScriptScope(n.Pos().Source, n)

	case *ast.FuncDecl:
		scope := ctx.openFuncScope(n.Name.Name, n)
		ctx.recordScope(n.Body, scope)

	case *ast.Lambda:
		scope := ctx.openLambdaScope(n)
		if _, ok := n.Body.(*ast.BlockStmt); ok {
			ctx.recordScope(n.Body, scope)
		}

	case *ast.Try:
		scope := ctx.openBlockScope("<try>", n)
		ctx.recordScope(n.Block, scope)

	case *ast.Trap:
		scope := ctx.openBlockScope("<trap>", n)
		ctx.recordScope(n.Block, scope)

	case *ast.Loop:
		scope := ctx.openBlockScope("<loop>", n)
		if _, ok := n.Body.(*ast.BlockStmt); ok {
			ctx.recordScope(n.Body, scope)
		}

	case *ast.EachLoop:
		scope := ctx.openBlockScope("<each>", n)
		if _, ok := n.Body.(*ast.BlockStmt); ok {
			ctx.recordScope(n.Body, scope)
		}

	case *ast.BlockStmt:
		if scope := ctx.info.Scopes[node]; scope != nil {
			return false // parent node did create scope -> do not cleanup scope
		}
		ctx.openBlockScope("<block>", n) // standalone block

	default:
		return false
	}

	return true
}

func indexFnNames(ctx *idxCtx, script *ast.Script) {
	scope := ctx.info.Scopes[script]

	for _, node := range script.Program {
		decl, ok := node.(*ast.FuncDecl)
		if !ok {
			continue
		}

		id := decl.Name
		if scope.LookupDirect(id.Name) != nil {
			ctx.recordRedecl(id)
			continue
		}

		fnScope := ctx.info.Scopes[node]
		args := argTypes(ctx, decl.Args)
		sig := &fnSignature{ret: createType(ctx, decl.Ret), args: args}
		obj := newFuncObj(decl, sig, fnScope)
		addFields(ctx, fnScope, args, decl.Args)

		// index function declaration in script scope:
		scope.add(obj)
	}
}

func indexVars(ctx *idxCtx, script *ast.Script) {
	ctx.scope = ctx.info.Scopes[script]
	defer ctx.closeScope()

	for _, n := range script.Program {
		indexNode(ctx, n)
	}
}

func indexNode(ctx *idxCtx, node ast.Node) {
	walk.Walk(node, walk.Rules(
		scopeEnv(ctx.info, &ctx.scope),
		walk.FromTop(func(node ast.Node) bool {
			return indexNodeVars(ctx, node)
		}),
	))
}

func indexNodeVars(ctx *idxCtx, node ast.Node) bool {
	switch n := node.(type) {
	// declarations
	case *ast.Decl:
		indexVarDecl(ctx, n)
		return false

	case *ast.Trap: // record 'catch' variable
		scope := ctx.info.Scopes[node]
		typ := createType(ctx, n.Type)
		scope.add(newVarObj(n, n.ID.Name, typ, scope))

	case *ast.Loop:
		// manually handle Loop here, as the loop can declare a varialbe in Init.
		// By remembering the old scope entry, we ensure the correct scope is in place
		// on function exit, as the loop body may open the current scope as well
		old := ctx.scope
		ctx.scope = ctx.info.Scopes[node]
		defer func() {
			ctx.scope = old
		}()

		indexNode(ctx, n.Init)
		indexNode(ctx, n.Cond)
		indexNode(ctx, n.Post)

		return false

	case *ast.EachLoop:
		typ := types.Def
		if n.Type != nil {
			typ = createType(ctx, n.Type)
		}
		scope := ctx.info.Scopes[node]
		v := newVarObj(n, n.ID.Name, typ, scope)
		scope.add(v)
		ctx.recordDecl(n.ID, v)

	case *ast.Lambda:
		scope := ctx.info.Scopes[node]
		args := argTypes(ctx, n.Args)
		addFields(ctx, scope, args, n.Args)

	// direct usage:
	case *ast.Ident:
		if k := n.Obj.Kind; k != ast.VarObj && k != ast.FunObj {
			break
		}

		obj := ctx.activeScope().Lookup(n.Name)
		if obj == nil {
			ctx.recordUnknownVar(n)
		} else {
			ctx.recordUse(n, obj)
		}

	case *ast.Call:
		if n.Value != nil {
			// method call -> need to index method receiver.
			// But defer name lookup to the type checker or even runtime.
			break
		}
		obj := ctx.activeScope().Lookup(n.ID.Name)
		if obj == nil {
			ctx.recordUnknownFunc(n.ID)
		} else {
			ctx.recordCall(n, obj)
		}

	// other
	case *ast.Literal:
		ctx.recordLit(n)
	}

	return true
}

func indexVarDecl(ctx *idxCtx, decl *ast.Decl) {
	typ := createType(ctx, decl.Type)

	for _, exp := range decl.Vars {
		scope := ctx.activeScope()
		var id *ast.Ident

		switch varDecl := exp.(type) {
		case *ast.Ident:
			id = varDecl

		case *ast.Assign:
			id = varDecl.LHS.(*ast.Ident)
			indexNode(ctx, varDecl.RHS)
		}

		if scope.LookupDirect(id.Name) != nil {
			ctx.recordRedecl(id)
			continue
		}

		obj := newVarObj(exp, id.Name, typ, scope)
		ctx.recordDecl(id, obj)
		scope.add(obj)
	}
}

func argTypes(ctx *idxCtx, fields []*ast.Field) []Type {
	if len(fields) == 0 {
		return nil
	}

	args := make([]Type, len(fields))
	for i, expr := range fields {
		args[i] = createType(ctx, expr.Type)
	}
	return args
}

func createType(ctx *idxCtx, node ast.TypeExpr) Type {
	switch expr := node.(type) {
	case *ast.Ident:
		t := ctx.types.Lookup(expr.Name)
		if t == nil {
			ctx.recordUnknownType(node, expr.Name)
		}
		return t

	case *ast.ArrType:
		t := createType(ctx, expr.Elt)
		if t != nil {
			t = &types.Array{Elt: t}
		}
		return t

	default:
		ctx.recordInvalidType(node)
		return nil
	}
}

func addFields(ctx *idxCtx, scope *Scope, types []Type, fields []*ast.Field) {
	for i, field := range fields {
		id := field.ID
		typ := types[i]
		if scope.LookupDirect(id.Name) != nil {
			ctx.recordRedecl(id)
			continue
		}
		scope.add(newParam(id, typ, scope))
	}
}
