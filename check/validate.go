package check

import (
	"github.com/urso/minpain/ast"
	"github.com/urso/minpain/ast/walk"
)

// validateAST checks if the AST is a valid script.  Errors are reporter to
// errs. The function returns false in case very critical errors have been
// identified, that make type checking impossible.
func validateAST(errs *multiErr, node ast.Node) (cont bool) {
	validateLoopOnlyStmts(errs, node)

	return true
}

// validateLoopOnlyStmts checks that `break` and `continue` statements
// only appear within a loop
func validateLoopOnlyStmts(errs *multiErr, node ast.Node) {
	levels := []int{0}

	walk.Walk(node, walk.Rules(
		walk.EnvSetupWith(
			func(n ast.Node) bool {
				switch n.(type) {
				case *ast.EachLoop, *ast.Loop:
					levels[len(levels)-1]++

				case *ast.Lambda:
					levels = append(levels, 0)

				default:
					return false
				}
				return true
			},
			func(n ast.Node) {
				switch n.(type) {
				case *ast.EachLoop, *ast.Loop:
					levels[len(levels)-1]--

				case *ast.Lambda:
					levels = levels[:len(levels)-1]
				}
			},
		),
		walk.FromBottom(func(node ast.Node) {
			if stmt, ok := node.(*ast.BranchStmt); ok && levels[len(levels)-1] == 0 {
				err := newNodeErrorf(node, "branching statement %v not in loop", stmt.Kind)
				errs.add(err)
			}
		}),
	))
}
