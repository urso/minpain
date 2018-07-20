package print

import (
	"fmt"
	"reflect"

	"github.com/urso/minpain/ast"
)

func toSexp(node ast.Node) interface{} {
	switch n := node.(type) {

	// top-level
	case *ast.Script:
		return newSexp("script", allToSexp(n.Program)...)
	case *ast.Decl:
		return newSexp("decl", typeStr(n.Type)).Add(allToSexp(n.Vars)...)
	case *ast.FuncDecl:
		return newSexp("func", typeStr(n.Ret), n.Name.Name,
			fieldsSexp(n.Args), toSexp(n.Body))

	// statements:
	case *ast.BlockStmt:
		return newSexp("block", allToSexp(n.Stmts)...)
	case *ast.Try:
		return newSexp("try", toSexp(n.Block)).Add(allToSexp(n.Traps)...)
	case *ast.Trap:
		return newSexp("catch", newSexp("", typeStr(n.Type), n.ID.Name), toSexp(n.Block))
	case *ast.IfStmt:
		e := newSexp("if", toSexp(n.Cond), toSexp(n.Then))
		if n.Else != nil {
			e.Add(toSexp(n.Else))
		}
		return e
	case *ast.BranchStmt:
		return n.Kind.String()
	case *ast.ReturnStmt:
		return newSexp("return", toSexp(n.Expr))
	case *ast.ThrowStmt:
		return newSexp("throw", toSexp(n.Expr))
	case *ast.Loop:
		e := newSexp("loop")
		if n.Init != nil {
			e.Add(kw{"init", toSexp(n.Init)})
		}
		if n.Cond != nil {
			e.Add(kw{"cond", toSexp(n.Cond)})
		}
		if n.Post != nil {
			e.Add(kw{"post", toSexp(n.Post)})
		}
		if n.Body != nil {
			e.Add(toSexp(n.Body))
		}
		return e
	case *ast.EachLoop:
		e := newSexp("each")
		if n.Type != nil {
			e.Add(typeStr(n.Type))
		}
		e.Add(n.ID, toSexp(n.Expr))
		if n.Body != nil {
			e.Add(toSexp(n.Body))
		}

	// expressions
	case *ast.Ident:
		return n.Name
	case *ast.ListInit:
		return newSexp("list", allToSexp(n.Elts)...)
	case *ast.MapInit:
		e := newSexp("map")
		for i, f := range n.Fields {
			e.Add(newSexp("", toSexp(f), toSexp(n.Values[i])))
		}
		return e
	case *ast.Assign:
		return newSexp(n.Op.String(), toSexp(n.LHS), toSexp(n.RHS))
	case *ast.CastExpr:
		return newSexp("cast", typeStr(n.To), toSexp(n.Expr))
	case *ast.CondExpr:
		e := newSexp(":?", toSexp(n.Check), toSexp(n.Then))
		if n.Else != nil {
			e.Add(toSexp(n.Else))
		}
		return e
	case *ast.UnaryOp:
		name := "post"
		if n.Pre {
			name = "pre"
		}
		return newSexp(name, n.Op.String(), toSexp(n.X))
	case *ast.BinOp:
		return newSexp(n.Op.String(), toSexp(n.X), toSexp(n.Y))
	case *ast.Literal:
		return n.Value
	case *ast.Call:
		switch n.Kind {
		case ast.CallFunction:
			return newSexp("call", n.ID.Name).Add(argsSexp(n.Args)...)
		case ast.CallNew:
			return newSexp("new", n.ID.Name).Add(argsSexp(n.Args)...)
		case ast.CallMethod:
			dot := "."
			if n.SafeDot {
				dot = "?."
			}
			return newSexp(dot, toSexp(n.Value), n.ID.Name).Add(argsSexp(n.Args)...)
		}

	case *ast.NewArray:
		return newSexp("newArray", typeStr(n.Typ)).Add(argsSexp(n.Init)...)
	case *ast.FieldAccess:
		dot := "."
		if n.SafeDot {
			dot = "?."
		}
		return newSexp(dot, toSexp(n.Value), n.ID.Name)
	case *ast.IdxAccess:
		dot := "."
		if n.SafeDot {
			dot = "?."
		}
		return newSexp(dot, toSexp(n.Value), fmt.Sprintf("%v", n.Idx))
	case *ast.Access:
		return newSexp("access", toSexp(n.Value), toSexp(n.Accessor))
	case *ast.Ref:
		return newSexp("ref", toSexp(n.Value), n.ID.Name)
	case *ast.Lambda:
		return newSexp("lambda", fieldsSexp(n.Args), toSexp(n.Body))

	default:
		panic("unhandled type")
	}
	return nil
}

func posValue(p ast.Pos) kw {
	return kw{"pos", fmt.Sprintf("(%v %v)", p.Line, p.Column)}
}

func allToSexp(vs ...interface{}) []interface{} {
	var to []interface{}
	for _, v := range vs {
		iterNodes(v, func(node ast.Node) {
			to = append(to, toSexp(node))
		})
	}
	return to
}

func collectSexp(arr *[]interface{}) func(ast.Node) {
	return func(node ast.Node) {
		*arr = append(*arr, toSexp(node))
	}
}

func iterNodes(in interface{}, fn func(ast.Node)) {
	tNode := reflect.TypeOf((*ast.Node)(nil)).Elem()

	v := reflect.ValueOf(in)
	if v.IsNil() {
		return
	}

	if k := v.Kind(); k != reflect.Array && k != reflect.Slice {
		node := v.Convert(tNode).Interface().(ast.Node)
		fn(node)
		return
	}

	l := v.Len()
	for i := 0; i < l; i++ {
		elem := v.Index(i)
		// if elem.Type().ConvertibleTo(tNode) {
		elem = elem.Convert(tNode)
		node := elem.Interface().(ast.Node)
		fn(node)
		// }
	}
}

func typeStr(n ast.TypeExpr) string {
	switch t := n.(type) {
	case *ast.ArrType:
		return typeStr(t.Elt.(ast.TypeExpr)) + "[]"
	case *ast.Ident:
		return t.Name
	}

	return ""
}

func argsSexp(args []ast.Expr) []interface{} {
	res := allToSexp(args)
	if len(res) == 0 {
		return []interface{}{newSexp("")}
	}
	return res
}

func fieldsSexp(fields []*ast.Field) *sexp {
	lst := make([]interface{}, len(fields))
	for i, field := range fields {
		if field.Type == nil {
			lst[i] = field.ID.Name
		} else {
			lst[i] = newSexp("", typeStr(field.Type), field.ID.Name)
		}
	}
	return newSexp("", lst...)
}
