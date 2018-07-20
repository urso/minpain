package ast

type (
	Decl struct {
		At   Pos
		Type TypeExpr
		Vars []Expr
	}

	FuncDecl struct {
		At   Pos
		Ret  TypeExpr
		Name *Ident
		Args []*Field
		Body *BlockStmt
	}

	Field struct {
		FieldPos Pos
		Type     TypeExpr
		ID       *Ident
	}
)

func (n *Decl) Pos() Pos  { return n.At }
func (_ *Decl) stmtNode() {}
func (_ *Decl) declNode() {}

func (n *FuncDecl) Pos() Pos  { return n.At }
func (_ *FuncDecl) declNode() {}

func (n *Field) Pos() Pos { return n.FieldPos }
