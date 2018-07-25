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
