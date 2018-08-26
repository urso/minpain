package ast

// expression nodes
type (
	Ident struct {
		NamePos Pos
		Name    string
		Obj     *Object
	}

	ListInit struct {
		Start Pos
		Elts  []Expr
	}

	MapInit struct {
		Start  Pos
		Fields []Expr
		Values []Expr
	}

	Assign struct {
		OpPos    Pos
		Op       OpKind
		LHS, RHS Expr
	}

	CastExpr struct {
		CastPos Pos
		To      TypeExpr
		Expr    Expr
	}

	CondExpr struct {
		Check Expr
		Then  Expr
		Else  Expr
	}

	UnaryOp struct {
		OpPos Pos
		X     Expr
		Pre   bool
		Op    OpKind
	}

	BinOp struct {
		X, Y  Expr
		OpPos Pos
		Op    OpKind
	}

	Literal struct {
		ValuePos Pos
		Kind     LiteralKind
		Value    string
	}

	Call struct {
		CallPos Pos
		Kind    CallKind

		Value   Expr   // method invocation
		ID      *Ident // function/new call
		SafeDot bool

		Args []Expr
	}

	NewArray struct {
		Typ  *ArrType
		Init []Expr
	}

	FieldAccess struct {
		At      Pos
		Value   Expr
		SafeDot bool
		ID      *Ident
	}

	IdxAccess struct {
		At      Pos
		Value   Expr
		SafeDot bool
		Idx     int
	}

	Access struct {
		At       Pos
		Value    Expr
		Accessor Expr
	}

	Ref struct {
		At    Pos
		Value Expr
		ID    *Ident
	}

	Lambda struct {
		At   Pos
		Args []*Field
		Body Node
	}
)

type LiteralKind uint8

type CallKind uint8

type OpKind uint16

//go:generate stringer -type=LiteralKind
const (
	BadLiteral LiteralKind = iota
	LitBool
	LitNull
	LitRegex
	LitString
	LitInt
	LitDecimal
)

//go:generate stringer -type=CallKind
const (
	CallFunction CallKind = iota
	CallMethod
	CallNew
)

//go:generate stringer -type=OpKind
const (
	OpBad OpKind = 0

	// op classes:
	OpClassArithmetic = (iota << 8)
	OpClassRegex
	OpClassBits
	OpClassBool
	OpClassCmpEq
	OpClassCmpNum
	OpClassOther
	OPClassMask = 0xff00

	// arithmetic

	OpMul OpKind = OpClassArithmetic + iota
	OpDiv
	OpRem
	OpAdd
	OpSub
	OpInc
	OpDec
	OpPos
	OpNeg

	// regex ops

	OpFind OpKind = OpClassRegex + iota
	OpMatch

	// bit ops

	OpLSH OpKind = OpClassBits + iota
	OpRSH
	OpUSH
	OpBitsAnd
	OpBitsXOR
	OpBitsOR
	OpBitsNot

	// bool ops
	OpAND OpKind = OpClassBool + iota
	OpOR
	OpBoolNot

	// comparison ops
	OpCmpEQ OpKind = OpClassCmpEq + iota
	OpCmpEQR
	OpCmpNE
	OpCmpNER

	OpCmpLT OpKind = OpClassCmpNum + iota
	OpCmpLTE
	OpCmpGT
	OpCmpGTE

	// other ops

	OpAssign OpKind = OpClassOther + iota
	OpInstanceOf
)

var opSyms = map[OpKind]string{
	OpMul:        "*",
	OpDiv:        "/",
	OpRem:        "%",
	OpAdd:        "+",
	OpSub:        "-",
	OpInc:        "++",
	OpDec:        "--",
	OpPos:        "+",
	OpNeg:        "-",
	OpFind:       "=~",
	OpLSH:        "<<",
	OpRSH:        ">>",
	OpUSH:        ">>>",
	OpBitsAnd:    "&",
	OpBitsXOR:    "^",
	OpBitsOR:     "|",
	OpBitsNot:    "~",
	OpAND:        "&&",
	OpOR:         "||",
	OpBoolNot:    "!",
	OpMatch:      "==~",
	OpCmpEQ:      "==",
	OpCmpEQR:     "===",
	OpCmpNE:      "!=",
	OpCmpNER:     "!==",
	OpCmpLT:      "<",
	OpCmpLTE:     "<=",
	OpCmpGT:      ">",
	OpCmpGTE:     ">=",
	OpAssign:     "=",
	OpInstanceOf: "<instanceof>",
}

func (n *Ident) Pos() Pos  { return n.NamePos }
func (_ *Ident) exprNode() {}
func (_ *Ident) typeNode() {}
func (n *Ident) Storable() bool {
	return n.Obj != nil && n.Obj.Kind == VarObj
}

func (n *ListInit) Pos() Pos  { return n.Start }
func (_ *ListInit) exprNode() {}
func (_ *ListInit) stmtNode() {} // TODO: check if these should be expressions only

func (n *MapInit) Pos() Pos  { return n.Start }
func (_ *MapInit) exprNode() {}
func (_ *MapInit) stmtNode() {} // TODO: check if these should be expressions only

func (n *Assign) Pos() Pos  { return n.OpPos }
func (_ *Assign) exprNode() {}
func (_ *Assign) stmtNode() {}

func (n *CastExpr) Pos() Pos  { return n.CastPos }
func (_ *CastExpr) exprNode() {}

func (n *CondExpr) Pos() Pos  { return n.Check.Pos() }
func (_ *CondExpr) exprNode() {}

func (n *UnaryOp) Pos() Pos  { return n.OpPos }
func (_ *UnaryOp) exprNode() {}
func (_ *UnaryOp) stmtNode() {}

func (n *BinOp) Pos() Pos  { return n.OpPos }
func (_ *BinOp) exprNode() {}

func (n *Literal) Pos() Pos  { return n.ValuePos }
func (_ *Literal) exprNode() {}
func (_ *Literal) stmtNode() {} // TODO: remove me

func (n *Call) Pos() Pos  { return n.CallPos }
func (_ *Call) exprNode() {}
func (_ *Call) stmtNode() {}

func (n *FieldAccess) Pos() Pos  { return n.At }
func (_ *FieldAccess) exprNode() {}
func (n *FieldAccess) Storable() bool {
	return true
}

func (n *IdxAccess) Pos() Pos  { return n.At }
func (_ *IdxAccess) exprNode() {}
func (_ *IdxAccess) Storable() bool {
	return true
}

func (n *Access) Pos() Pos  { return n.At }
func (_ *Access) exprNode() {}
func (_ *Access) Storable() bool {
	return true
}

func (n *ArrType) Pos() Pos  { return n.Start }
func (_ *ArrType) exprNode() {}

func (n *NewArray) Pos() Pos  { return n.Typ.Pos() }
func (_ *NewArray) exprNode() {}

func (n *Lambda) Pos() Pos  { return n.At }
func (_ *Lambda) exprNode() {}

func (n *Ref) Pos() Pos  { return n.At }
func (_ *Ref) exprNode() {}

func (o OpKind) Symbol() string {
	if s := opSyms[o]; s != "" {
		return s
	}
	return "<" + o.String() + ">"
}
