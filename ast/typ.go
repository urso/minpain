package ast

type (
	ArrType struct {
		Start Pos
		Len   Expr
		Elt   TypeExpr
	}
)

func (_ *ArrType) typeNode() {}
