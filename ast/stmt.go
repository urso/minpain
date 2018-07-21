package ast

type (
	BlockStmt struct {
		BlockPos Pos
		Stmts    []Stmt
	}

	Try struct {
		At    Pos
		Block *BlockStmt
		Traps []*Trap
	}

	Trap struct {
		TrapPos Pos
		Type    TypeExpr
		ID      *Ident
		Block   *BlockStmt
	}

	IfStmt struct {
		IfPos Pos
		Cond  Expr
		Then  Stmt
		Else  Stmt
	}

	BranchStmt struct {
		At   Pos
		Kind BranchKind
	}

	ReturnStmt struct {
		At   Pos
		Expr Expr
	}

	ThrowStmt struct {
		At   Pos
		Expr Expr
	}

	Loop struct {
		LoopPos Pos
		Kind    LoopKind
		Init    Node // can be an expressions or declaration
		Cond    Expr
		Post    Expr
		Body    Stmt
	}

	EachLoop struct {
		LoopPos Pos
		Kind    LoopKind
		Type    TypeExpr
		ID      *Ident
		Expr    Expr
		Body    Stmt
	}
)

type BranchKind uint8

type LoopKind uint8

//go:generate stringer -type=BranchKind -linecomment=true
const (
	BranchInvalid  BranchKind = iota // invalid
	BranchBreak                      // break
	BranchContinue                   // continue
)

//go:generate stringer -type=BranchKind
const (
	LoopUnknown LoopKind = iota
	LoopFor
	LoopWhile
	LoopDo
	LoopEach
	LoopInEach
)

func (n *BlockStmt) Pos() Pos  { return n.BlockPos }
func (_ *BlockStmt) stmtNode() {}

func (n *Try) Pos() Pos  { return n.At }
func (_ *Try) stmtNode() {}
func (n *Trap) Pos() Pos { return n.TrapPos }

func (n *IfStmt) Pos() Pos  { return n.IfPos }
func (_ *IfStmt) stmtNode() {}

func (n *Loop) Pos() Pos  { return n.LoopPos }
func (_ *Loop) stmtNode() {}

func (n *EachLoop) Pos() Pos  { return n.LoopPos }
func (_ *EachLoop) stmtNode() {}

func (n *ThrowStmt) Pos() Pos  { return n.At }
func (_ *ThrowStmt) stmtNode() {}

func (n *BranchStmt) Pos() Pos  { return n.At }
func (_ *BranchStmt) stmtNode() {}

func (n *ReturnStmt) Pos() Pos  { return n.At }
func (_ *ReturnStmt) stmtNode() {}
