package check

import (
	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/ast/walk"
	"github.com/urso/minpain/types"
)

type idxCtx struct {
	info  *Info
	scope *Scope
	errs  *multiErr

	types map[string]Type
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

func (ctx *idxCtx) recordErr(err error) { ctx.errs.add(err) }

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

func indexScript(ctx *idxCtx, script *ast.Script) {
	// 1. phase: create and link block scopes
	buildScopes(ctx, script)

	// 2. phase: collect function names into the script scope
	//    As functions are defined before any other statements, but function can
	//    call each other, we have to collect the names into the scope before checking
	//    any variable access is the main script.
	indexFnNames(ctx, script)

	// index all variable declarations and uses
	doIndexScript(ctx, script)
}

func buildScopes(ctx *idxCtx, script *ast.Script) {
	walk.Walk(script, walk.Rules(
		walk.EnvSetupWith(
			func(node ast.Node) bool {
				return buildNodeScope(ctx, node)
			},
			ctx.closeScope,
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
		sig := fnSignature{ret: getType(ctx, decl.Ret), args: args}
		obj := newFuncObj(decl, sig, fnScope)
		addFields(ctx, fnScope, args, decl.Args)

		// index function declaration in script scope:
		scope.add(obj)
	}
}

func doIndexScript(ctx *idxCtx, script *ast.Script) {
	ctx.scope = ctx.info.Scopes[script]
	defer ctx.closeScope()

	for _, n := range script.Program {
		indexNode(ctx, n)
	}
}

func indexNode(ctx *idxCtx, node ast.Node) {
	walk.Walk(node, walk.Rules(
		walk.EnvSetupWith(openBlockScopes(ctx), ctx.closeScope),
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
		typ := getType(ctx, n.Type)
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
		scope := ctx.info.Scopes[node]
		typ := types.Def
		if n.Type != nil {
			typ = getType(ctx, n.Type)
		}

	case *ast.Lambda:

		// usage:
	case *ast.Ident:

	case *ast.Literal:

	case *ast.Call:

	case *ast.FieldAccess:

	case *ast.IdxAccess:

	case *ast.Ref:

	}

	return true
}

func indexVarDecl(ctx *idxCtx, decl *ast.Decl) {
	typ := getType(ctx, decl.Type)

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

func openBlockScopes(ctx *idxCtx) func(ast.Node) bool {
	return func(node ast.Node) (ok bool) {
		if _, ok = node.(*ast.BlockStmt); ok {
			ctx.scope = ctx.info.Scopes[node]
		}
		return
	}
}

func argTypes(ctx *idxCtx, fields []*ast.Field) []Type {
	if len(fields) == 0 {
		return nil
	}

	args := make([]Type, len(fields))
	for i, expr := range fields {
		args[i] = getType(ctx, expr.Type)
	}
	return args
}

func getType(ctx *idxCtx, node ast.TypeExpr) Type {
	switch expr := node.(type) {
	case *ast.Ident:
		t := ctx.types[expr.Name]
		if t == nil {
			ctx.recordUnknownType(node, expr.Name)
		}
		return t

	case *ast.ArrType:
		t := getType(ctx, expr.Elt)
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

/*
func indexScript(ctx *idxCtx, script *ast.Script) {
	ctx.openScriptScope(script.Pos().Source, script)
	defer ctx.closeScope()
	for _, stmt := range script.Program {
		indexNode(ctx, stmt)
	}

	for decl, obj := range ctx.info.Functions {
		ctx.withScope(obj.scope, func() {
			indexBody(ctx, "", decl.Body)
		})
	}
}

func indexNode(ctx *idxCtx, node ast.Node) {
	if node == nil {
		return
	}

	switch n := node.(type) {
	case *ast.Script:
		panic("unexpected script node")
	case *ast.Decl:
		indexVarDecl(ctx, n)
	case *ast.FuncDecl:
		indexFuncDecl(ctx, n)
	case *ast.BlockStmt:
		indexBlock(ctx, "<block>", n)
	case *ast.Try:
		indexBody(ctx, "<try>", n.Block)
		for _, trap := range n.Traps {
			indexTrap(ctx, trap)
		}
	case *ast.IfStmt:
		indexNode(ctx, n.Cond)
		indexBody(ctx, "<if>", n.Then)
		indexBody(ctx, "<else>", n.Else)
	case *ast.ReturnStmt:
		indexNode(ctx, n.Expr)
	case *ast.ThrowStmt:
		indexNode(ctx, n.Expr)
	case *ast.Loop:
		indexLoop(ctx, n)
	case *ast.EachLoop:
		indexEachLoop(ctx, n)
	case *ast.Ident:
		indexID(ctx, n)
	case *ast.ListInit:
		indexExprs(ctx, n.Elts)
	case *ast.MapInit:
		indexExprs(ctx, n.Fields)
		indexExprs(ctx, n.Values)
	case *ast.Assign:
		indexNode(ctx, n.LHS)
		indexNode(ctx, n.RHS)
	case *ast.CastExpr:
		indexNode(ctx, n.Expr)
	case *ast.CondExpr:
		indexNode(ctx, n.Check)
		indexNode(ctx, n.Then)
		indexNode(ctx, n.Else)
	case *ast.UnaryOp:
		indexNode(ctx, n.X)
	case *ast.BinOp:
		indexNode(ctx, n.X)
		indexNode(ctx, n.Y)
	case *ast.Literal:
		ctx.recordLit(n)
	case *ast.Call:
		indexCall(ctx, n)
	case *ast.NewArray:
		indexExprs(ctx, n.Init)
	case *ast.FieldAccess:
		indexNode(ctx, n.Value)
	case *ast.IdxAccess:
		indexNode(ctx, n.Value)
	case *ast.Access:
		indexNode(ctx, n.Value)
		indexNode(ctx, n.Accessor)
	case *ast.Ref:
		indexNode(ctx, n.Value)
	case *ast.Lambda:
		indexLambda(ctx, n)
	default:
		panic("index: unhandled ast node")
	}
}

func indexFuncDecl(ctx *idxCtx, decl *ast.FuncDecl) {
	// TODO: create function signature

	id := decl.Name
	scope := ctx.activeScope()
	if scope.LookupDirect(id.Name) != nil {
		ctx.recordRedecl(id)
		return
	}

	args := argTypes(ctx, decl.Args)
	sig := fnSignature{ret: getType(ctx, decl.Ret), args: args}
	obj := newFuncObj(decl, sig, scope)
	ctx.recordFunc(decl, obj)

	fnScope := ctx.openFuncScope(id.Name, decl)
	defer ctx.closeScope()

	fnScope.add(obj)
	addFields(ctx, fnScope, args, decl.Args)

	// postpose processing function bodies until we have collected all function
	// declarations
}

func indexVarDecl(ctx *idxCtx, decl *ast.Decl) {
	typ := getType(ctx, decl.Type)

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

func indexID(ctx *idxCtx, id *ast.Ident) {
	if k := id.Obj.Kind; k != ast.VarObj && k != ast.FunObj {
		return
	}

	// variable usage only
	scope := ctx.activeScope()
	obj := scope.Lookup(id.Name)
	if obj == nil {
		ctx.recordUnknownVar(id)
	} else {
		ctx.recordUse(id, obj)
	}
}

func indexTrap(ctx *idxCtx, trap *ast.Trap) {
	scope := ctx.openBlockScope("<trap>", trap)
	defer ctx.closeScope()

	typ := getType(ctx, trap.Type)
	scope.add(newVarObj(trap, trap.ID.Name, typ, scope))
	indexStmts(ctx, trap.Block.Stmts)
}

func indexLoop(ctx *idxCtx, stmt *ast.Loop) {
	ctx.openBlockScope("<loop>", stmt)
	defer ctx.closeScope()

	indexNode(ctx, stmt.Init)
	indexNode(ctx, stmt.Cond)
	indexNode(ctx, stmt.Post)
	indexBody(ctx, "", stmt.Body)
}

func indexEachLoop(ctx *idxCtx, stmt *ast.EachLoop) {
	indexNode(ctx, stmt.Expr)

	scope := ctx.openBlockScope("<loop>", stmt)
	defer ctx.closeScope()

	typ := types.Def
	if stmt.Type != nil {
		typ = getType(ctx, stmt.Type)
	}
	scope.add(newVarObj(stmt, stmt.ID.Name, typ, scope))
	indexBody(ctx, "", stmt.Body)
}

func indexBody(ctx *idxCtx, name string, node ast.Stmt) {
	switch body := node.(type) {
	case *ast.BlockStmt:
		if name == "" {
			ctx.recordScope(node, ctx.activeScope())
			indexStmts(ctx, body.Stmts)
		} else {
			indexBlock(ctx, name, body)
		}
	default:
		indexNode(ctx, node)
	}
}
func indexBlock(ctx *idxCtx, name string, stmt *ast.BlockStmt) {
	if stmt != nil {
		ctx.openBlockScope(name, stmt)
		defer ctx.closeScope()
		indexStmts(ctx, stmt.Stmts)
	}
}

func indexCall(ctx *idxCtx, expr *ast.Call) {
	if expr.Value != nil {
		indexNode(ctx, expr.Value)
	} else {
		scope := ctx.activeScope()
		if id := expr.ID; scope.Lookup(id.Name) == nil {
			ctx.recordUnknownFunc(id)
		}
	}
	indexExprs(ctx, expr.Args)
}

func indexLambda(ctx *idxCtx, lambda *ast.Lambda) {
	scope := ctx.openLambdaScope(lambda)
	defer ctx.closeScope()

	args := argTypes(ctx, lambda.Args)
	addFields(ctx, scope, args, lambda.Args)
	if stmt, ok := lambda.Body.(ast.Stmt); ok {
		indexBody(ctx, "", stmt)
	} else {
		indexNode(ctx, lambda.Body)
	}
}

func indexExprs(ctx *idxCtx, exprs []ast.Expr) {
	for _, e := range exprs {
		indexNode(ctx, e)
	}
}

func indexStmts(ctx *idxCtx, stmts []ast.Stmt) {
	for _, stmt := range stmts {
		indexNode(ctx, stmt)
	}
}

func addFields(ctx *idxCtx, scope *Scope, types []Type, fields []*ast.Field) {
	// TODO: extract type from argument

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

func argTypes(ctx *idxCtx, fields []*ast.Field) []Type {
	if len(fields) == 0 {
		return nil
	}

	args := make([]Type, len(fields))
	for i, expr := range fields {
		args[i] = getType(ctx, expr.Type)
	}
	return args
}

func getType(ctx *idxCtx, node ast.TypeExpr) Type {
	switch expr := node.(type) {
	case *ast.Ident:
		t := ctx.types[expr.Name]
		if t == nil {
			ctx.recordUnknownType(node, expr.Name)
		}
		return t

	case *ast.ArrType:
		t := getType(ctx, expr.Elt)
		if t != nil {
			t = &types.Array{t}
		}
		return t

	default:
		ctx.recordInvalidType(node)
		return nil
	}
}
*/
