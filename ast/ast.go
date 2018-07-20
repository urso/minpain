package ast

type (
	Script struct {
		Start   Pos
		Program []Node
	}

	Node interface {
		Pos() Pos
	}

	Expr interface {
		Node
		exprNode()
	}

	TypeExpr interface {
		Node
		typeNode()
	}

	DeclExpr interface {
		Node
		declNode()
	}

	Stmt interface {
		Node
		stmtNode()
	}

	Object struct {
		Kind ObjKind
		Name string
	}
)

type ObjKind uint8

//go:generate stringer -type=ObjKind
const (
	BadObj ObjKind = iota
	TypObj
	VarObj
	FunObj
	FieldObj
)

func IsDecl(n Node) bool {
	_, ok := n.(DeclExpr)
	return ok
}

func IsExpr(n Node) bool {
	_, ok := n.(Expr)
	return ok
}

func IsStmt(n Node) bool {
	_, ok := n.(Stmt)
	return ok
}

func (s *Script) Pos() Pos       { return s.Start }
func (s *Script) Source() string { return s.Start.Source }
