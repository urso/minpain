package check

import (
	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/ast/walk"
)

func scopeEnv(info *Info, scopeVar **Scope) walk.Visitor {
	isBlock := func(node ast.Node) bool {
		_, ok := node.(*ast.BlockStmt)
		return ok
	}

	openBlockScope := func(node ast.Node) (ok bool) {
		if ok = isBlock(node); ok {
			tmp := info.Scopes[node]
			if ok = (*scopeVar != tmp); ok {
				*scopeVar = tmp
			}
		}
		return
	}

	closeBlockScope := func(node ast.Node) {
		if isBlock(node) {
			parent := (*scopeVar).Parent()
			*scopeVar = parent
		}
	}
	return walk.EnvSetupWith(openBlockScope, closeBlockScope)
}

func findID(n ast.Node, i int, kind ast.ObjKind) *ast.Ident {
	n = findNode(n, i, func(n ast.Node) bool {
		id, ok := n.(*ast.Ident)
		if !ok {
			return false
		}
		return id.Obj.Kind == kind
	})

	if n == nil {
		return nil
	}
	return n.(*ast.Ident)
}

func findNode(n ast.Node, i int, fn func(ast.Node) bool) ast.Node {
	var found ast.Node
	ast.Walk(n, ast.Inspect(func(node ast.Node) bool {
		if fn(node) {
			if i == 0 {
				found = node
				return false
			}
			i--
		}
		return true
	}))
	return found
}
