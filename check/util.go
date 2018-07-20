package check

import (
	"github.com/urso/minpain/ast"
)

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
