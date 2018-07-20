package ast

type Visitor func(Node) Visitor

func Walk(node Node, v Visitor) {
	if node == nil {
		return
	}

	v = v(node)
	if v == nil {
		return
	}

	switch n := node.(type) {
	case *Script:
		for _, stmts := range n.Program {
			Walk(stmts, v)
		}

	// declarations
	case *Decl:
		Walk(n.Type, v)
		walkExprs(v, n.Vars)
	case *FuncDecl:
		Walk(n.Ret, v)
		Walk(n.Name, v)
		Walk(n.Body, v)

	// expressions
	case *Ident, *Literal:
		// noop
	case *ListInit:
		walkExprs(v, n.Elts)
	case *MapInit:
		walkExprs(v, n.Values)
	case *Assign:
		Walk(n.LHS, v)
		Walk(n.RHS, v)
	case *CastExpr:
		Walk(n.To, v)
		Walk(n.Expr, v)
	case *CondExpr:
		Walk(n.Check, v)
		Walk(n.Then, v)
		Walk(n.Else, v)
	case *UnaryOp:
		Walk(n.X, v)
	case *BinOp:
		Walk(n.X, v)
		Walk(n.Y, v)
	case *Call:
		if n.Kind == CallMethod {
			Walk(n.Value, v)
		}
		Walk(n.ID, v)
		walkExprs(v, n.Args)
	case *NewArray:
		Walk(n.Typ, v)
		walkExprs(v, n.Init)
	case *FieldAccess:
		Walk(n.Value, v)
		Walk(n.ID, v)
	case *IdxAccess:
		Walk(n.Value, v)
	case *Access:
		Walk(n.Value, v)
		Walk(n.Accessor, v)
	case *Ref:
		Walk(n.Value, v)
		Walk(n.ID, v)
	case *Lambda:
		Walk(n.Body, v)

	// statements:
	case *BranchStmt:
		// noop
	case *BlockStmt:
		walkStmt(v, n.Stmts)
	case *Try:
		Walk(n.Block, v)
		for _, t := range n.Traps {
			Walk(t, v)
		}
	case *Trap:
		Walk(n.Type, v)
		Walk(n.ID, v)
		Walk(n.Block, v)
	case *IfStmt:
		Walk(n.Cond, v)
		Walk(n.Then, v)
		Walk(n.Else, v)
	case *ReturnStmt:
		Walk(n.Expr, v)
	case *ThrowStmt:
		Walk(n.Expr, v)
	case *Loop:
		Walk(n.Init, v)
		Walk(n.Cond, v)
		Walk(n.Post, v)
		Walk(n.Body, v)
	case *EachLoop:
		Walk(n.Type, v)
		Walk(n.ID, v)
		Walk(n.Expr, v)
		Walk(n.Body, v)

	// types
	case *ArrType:
		Walk(n.Len, v)
		Walk(n.Elt, v)
	}

	v(nil)
}

func walkExprs(v Visitor, exprs []Expr) {
	for _, e := range exprs {
		Walk(e, v)
	}
}

func walkStmt(v Visitor, stmts []Stmt) {
	for _, s := range stmts {
		Walk(s, v)
	}
}

func Inspect(fn func(Node) bool) Visitor {
	return func(n Node) Visitor {
		if !fn(n) {
			return nil
		}
		return Inspect(fn)
	}
}

func Each(fn func(Node)) Visitor {
	return Inspect(func(n Node) bool {
		fn(n)
		return true
	})
}

func Guard(pred func(Node) bool, fn func(Node)) Visitor {
	return Each(func(n Node) {
		if pred(n) {
			fn(n)
		}
	})
}

func Once(fn func(Node)) Visitor {
	return func(n Node) Visitor {
		fn(n)
		return nil
	}
}
