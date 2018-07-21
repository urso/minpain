package walk

import "github.com/urso/minpain/ast"

type Visitor interface {
	// Enter/Exit notify the visitor of nodes being visited, top-down.  These
	// calls should be used to initialize/finalize some additional context in the
	// background while walking.
	// Enter must set visit to true if the node should be followed by the walker.
	// Exit will only be called if callExit is true.
	Enter(node ast.Node) (visit, callExit bool)
	Exit(node ast.Node)

	// FromTop is called when a node is visited. Children nodes will not be
	// visited if FromTop returns false
	FromTop(node ast.Node) (visit bool)

	// FromBottom is called after the child nodes have been visited.  The AST can
	// be visited from from bottom, in a left-to-right manner by using
	// FromBottom.
	FromBottom(node ast.Node)
}

func Walk(node ast.Node, v Visitor) {
	if node == nil {
		return
	}

	visit, callExit := v.Enter(node)
	if callExit {
		defer v.Exit(node)
	}
	if !visit {
		return
	}

	visit = v.FromTop(node)
	if visit {
		walkChildren(node, v)
	}
	v.FromBottom(node)
}

func walkChildren(node ast.Node, v Visitor) {
	switch n := node.(type) {
	case *ast.Script:
		for _, stmts := range n.Program {
			Walk(stmts, v)
		}

	// declarations
	case *ast.Decl:
		Walk(n.Type, v)
		walkExprs(v, n.Vars)
	case *ast.FuncDecl:
		Walk(n.Ret, v)
		Walk(n.Body, v)

	// expressions
	case *ast.Ident, *ast.Literal:
		// noop
	case *ast.ListInit:
		walkExprs(v, n.Elts)
	case *ast.MapInit:
		walkExprs(v, n.Values)
	case *ast.Assign:
		Walk(n.LHS, v)
		Walk(n.RHS, v)
	case *ast.CastExpr:
		Walk(n.To, v)
		Walk(n.Expr, v)
	case *ast.CondExpr:
		Walk(n.Check, v)
		Walk(n.Then, v)
		Walk(n.Else, v)
	case *ast.UnaryOp:
		Walk(n.X, v)
	case *ast.BinOp:
		Walk(n.X, v)
		Walk(n.Y, v)
	case *ast.Call:
		if n.Kind == ast.CallMethod {
			Walk(n.Value, v)
		}
		walkExprs(v, n.Args)
	case *ast.NewArray:
		Walk(n.Typ, v)
		walkExprs(v, n.Init)
	case *ast.FieldAccess:
		Walk(n.Value, v)
	case *ast.IdxAccess:
		Walk(n.Value, v)
	case *ast.Access:
		Walk(n.Value, v)
		Walk(n.Accessor, v)
	case *ast.Ref:
		Walk(n.Value, v)
	case *ast.Lambda:
		Walk(n.Body, v)

	// statements:
	case *ast.BranchStmt:
		// noop
	case *ast.BlockStmt:
		walkStmt(v, n.Stmts)
	case *ast.Try:
		Walk(n.Block, v)
		for _, t := range n.Traps {
			Walk(t, v)
		}
	case *ast.Trap:
		Walk(n.Type, v)
		Walk(n.Block, v)
	case *ast.IfStmt:
		Walk(n.Cond, v)
		Walk(n.Then, v)
		Walk(n.Else, v)
	case *ast.ReturnStmt:
		Walk(n.Expr, v)
	case *ast.ThrowStmt:
		Walk(n.Expr, v)
	case *ast.Loop:
		Walk(n.Init, v)
		Walk(n.Cond, v)
		Walk(n.Post, v)
		Walk(n.Body, v)
	case *ast.EachLoop:
		Walk(n.Type, v)
		Walk(n.Expr, v)
		Walk(n.Body, v)

	// types
	case *ast.ArrType:
		Walk(n.Len, v)
		Walk(n.Elt, v)
	}
}

func walkExprs(v Visitor, exprs []ast.Expr) {
	for _, e := range exprs {
		Walk(e, v)
	}
}

func walkStmt(v Visitor, stmts []ast.Stmt) {
	for _, s := range stmts {
		Walk(s, v)
	}
}
