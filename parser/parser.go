package parser

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/urso/minpain/ast"
	pantlr "github.com/urso/minpain/parser/internal/parser"
	"github.com/urso/minpain/print"
)

type Parser struct {
	errListener multiError
	isType      func(string) bool
}

type parseListener struct {
	pantlr.BasePainlessParserListener

	source string
	debug  bool

	st *state
}

type multiError interface {
	Err() error
	Add(err error)
}

func NewParser(errListener multiError, isType func(string) bool) *Parser {
	if isType == nil {
		isType = func(name string) bool { return name == "def" }
	}

	return &Parser{errListener: errListener, isType: isType}
}

func (p *Parser) Parse(origin, input string, debug bool) (ast.Node, error) {
	is := antlr.NewInputStream(input)
	lex := pantlr.NewLexer(is, p.isType)
	toks := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	parser := pantlr.NewPainlessParser(toks)

	st := &state{}

	pl := &parseListener{source: origin, debug: debug, st: st}

	// install our own error listener
	errs := &errorListener{errs: p.errListener, op: "painless/parse", source: origin}
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errs)

	// let's parse
	tree := parser.Source() // get grammar root node
	if err := errs.Error(); err != nil {
		return nil, err
	}

	// parse complete, let's walk the tree and build the actual ast
	walker := antlr.NewParseTreeWalker()
	walker.Walk(pl, tree)
	return st.Pop(), nil
}

var _ pantlr.PainlessParserListener = &parseListener{}

func (l *parseListener) EnterSource(ctx *pantlr.SourceContext) { l.st.PushStack() }
func (l *parseListener) ExitSource(ctx *pantlr.SourceContext) {
	prog := l.st.PopStack()
	l.st.Push(&ast.Script{
		Start:   l.startPos(ctx),
		Program: prog,
	})
}

// trace handler:

func (l *parseListener) VisitTerminal(node antlr.TerminalNode) {
	if l.debug {
		fmt.Println("visit terminal: ", node.GetSymbol())
	}
}

func (l *parseListener) VisitErrorNode(node antlr.ErrorNode) {
	if l.debug {
		fmt.Println("visit error node: ", node.GetSymbol())
	}
}

func (l *parseListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if l.debug {
		fmt.Println("EnterEveryRule:", ctx.GetText())
	}
}

func (l *parseListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	if l.debug {
		fmt.Println("ExitEveryRule:", ctx.GetText())
	}
}

//
// handle function:
// ----------------

func (l *parseListener) EnterFunction(ctx *pantlr.FunctionContext) { l.st.PushStack() }
func (l *parseListener) ExitFunction(ctx *pantlr.FunctionContext) {
	stack := l.st.PopStack()
	body := stack.Pop().(*ast.BlockStmt)
	declType := stack[0].(ast.TypeExpr)

	paramTypes := stack[1:]
	paramNames := ctx.Parameters().(*pantlr.ParametersContext).AllID()
	if len(paramTypes) != len(paramNames) {
		panic("function parameters mis-count")
	}

	fields := make([]*ast.Field, len(paramTypes))
	for i, node := range paramTypes {
		fields[i] = &ast.Field{
			FieldPos: node.Pos(),
			Type:     node.(ast.TypeExpr),
			ID:       l.newIdent(paramNames[i], ast.VarObj),
		}
	}

	l.st.Push(&ast.FuncDecl{
		At:   l.startPos(ctx),
		Ret:  declType,
		Name: l.newIdent(ctx.ID(), ast.FunObj),
		Args: fields,
		Body: body,
	})
}

//
// handle rstatement:
// ------------------

func (l *parseListener) ExitIf(ctx *pantlr.IfContext) {
	var otherwise ast.Stmt
	if len(ctx.AllTrailer()) > 1 {
		otherwise = l.st.Pop().(ast.Stmt)
	}
	then := l.st.Pop().(ast.Stmt)
	cond := l.st.Pop().(ast.Expr)

	l.st.Push(&ast.IfStmt{
		IfPos: l.startPos(ctx),
		Cond:  cond,
		Then:  then,
		Else:  otherwise,
	})
}

func (l *parseListener) ExitWhile(ctx *pantlr.WhileContext) {
	var body ast.Stmt
	expr, trailer := l.st.Pop2()
	if trailer != nil {
		body = trailer.(ast.Stmt)
	}

	l.st.Push(&ast.Loop{
		LoopPos: l.startPos(ctx),
		Kind:    ast.LoopWhile,
		Cond:    expr.(ast.Expr),
		Body:    body,
	})
}

func (l *parseListener) EnterFor(ctx *pantlr.ForContext) { l.st.PushStack() }
func (l *parseListener) ExitFor(ctx *pantlr.ForContext) {
	nodes := l.st.PopStack()
	fmt.Println("for nodes:")
	for _, n := range nodes {
		print.Print(n)
	}

	nodeWhen := func(ctx antlr.ParserRuleContext) ast.Node {
		if ctx == nil {
			return nil
		}
		return nodes.Pop()
	}

	toExpr := func(n ast.Node) ast.Expr {
		if n == nil {
			return nil
		}
		return n.(ast.Expr)
	}

	body := nodes.Pop().(ast.Stmt)
	post := nodeWhen(ctx.Afterthought())
	cond := nodeWhen(ctx.Expression())
	init := nodeWhen(ctx.Initializer())
	if len(nodes) > 0 {
		panic("invalid for context")
	}

	l.st.Push(&ast.Loop{
		LoopPos: l.startPos(ctx),
		Kind:    ast.LoopFor,
		Init:    init,
		Cond:    toExpr(cond),
		Post:    toExpr(post),
		Body:    body,
	})
}

func (l *parseListener) ExitEach(ctx *pantlr.EachContext) {
	expr, body := l.st.Pop2()
	typ := l.st.Pop()
	l.st.Push(&ast.EachLoop{
		LoopPos: l.startPos(ctx),
		Kind:    ast.LoopEach,
		Type:    typ.(ast.TypeExpr),
		ID:      l.newIdent(ctx.ID(), ast.VarObj),
		Expr:    expr.(ast.Expr),
		Body:    body.(ast.Stmt),
	})
}

func (l *parseListener) ExitIneach(ctx *pantlr.IneachContext) {
	expr, body := l.st.Pop2()
	l.st.Push(&ast.EachLoop{
		LoopPos: l.startPos(ctx),
		Kind:    ast.LoopInEach,
		ID:      l.newIdent(ctx.ID(), ast.VarObj),
		Expr:    expr.(ast.Expr),
		Body:    body.(ast.Stmt),
	})
}

func (l *parseListener) ExitEmpty(ctx *pantlr.EmptyContext) { l.st.Push(nil) }

//
// handle dstatement
// -----------------

func (l *parseListener) ExitDo(ctx *pantlr.DoContext) {
	body, cond := l.st.Pop2()
	l.st.Push(&ast.Loop{
		LoopPos: l.startPos(ctx),
		Kind:    ast.LoopDo,
		Cond:    cond.(ast.Expr),
		Body:    body.(ast.Stmt),
	})
}

func (l *parseListener) ExitBreak(ctx *pantlr.BreakContext) {
	l.st.Push(&ast.BranchStmt{
		At:   l.startPos(ctx),
		Kind: ast.BranchBreak,
	})
}

func (l *parseListener) ExitContinue(ctx *pantlr.ContinueContext) {
	l.st.Push(&ast.BranchStmt{
		At:   l.startPos(ctx),
		Kind: ast.BranchContinue,
	})
}

func (l *parseListener) ExitReturn(ctx *pantlr.ReturnContext) {
	l.st.Push(&ast.ReturnStmt{
		At:   l.startPos(ctx),
		Expr: l.st.Pop().(ast.Expr),
	})
}

func (l *parseListener) ExitThrow(ctx *pantlr.ThrowContext) {
	l.st.Push(&ast.ThrowStmt{
		At:   l.startPos(ctx),
		Expr: l.st.Pop().(ast.Expr),
	})
}

//
// handle declaration
// ------------------

func (l *parseListener) EnterDeclaration(ctx *pantlr.DeclarationContext) { l.st.PushStack() }
func (l *parseListener) ExitDeclaration(ctx *pantlr.DeclarationContext) {
	stack := l.st.PopStack()
	typ := stack[0].(ast.TypeExpr)

	vars := make([]ast.Expr, len(stack)-1)
	for i, node := range stack[1:] {
		vars[i] = node.(ast.Expr)
	}

	l.st.Push(&ast.Decl{
		At:   l.startPos(ctx),
		Type: typ,
		Vars: vars,
	})
}

func (l *parseListener) ExitDeclvar(ctx *pantlr.DeclvarContext) {
	id := l.newIdent(ctx.ID(), ast.VarObj)

	if terminal := ctx.ASSIGN(); terminal != nil {
		expr := l.st.Pop().(ast.Expr)
		l.st.Push(&ast.Assign{
			OpPos: tokPos(l.source, terminal.GetSymbol()),
			Op:    ast.OpAssign,
			LHS:   id,
			RHS:   expr,
		})
	} else {
		l.st.Push(id)
	}
}

//
// handle try:
// ------------

func (l *parseListener) EnterTry(ctx *pantlr.TryContext) { l.st.PushStack() }
func (l *parseListener) ExitTry(ctx *pantlr.TryContext) {
	raw := l.st.PopN(len(l.st.ActiveStack()) - 1)
	block := l.st.Pop().(*ast.BlockStmt)
	l.st.PopStack() // cleanup local state

	traps := make([]*ast.Trap, len(raw))
	for i, node := range raw {
		traps[i] = node.(*ast.Trap)
	}

	l.st.Push(&ast.Try{
		At:    l.startPos(ctx),
		Block: block,
		Traps: traps,
	})
}

func (l *parseListener) ExitTrap(ctx *pantlr.TrapContext) {
	l.st.Push(&ast.Trap{
		TrapPos: l.startPos(ctx),
		Type:    l.newIdent(ctx.TYPE(), ast.TypObj),
		ID:      l.newIdent(ctx.ID(), ast.VarObj),
		Block:   l.st.Pop().(*ast.BlockStmt),
	})
}

//
// handle block:
// -------------

func (l *parseListener) EnterBlock(ctx *pantlr.BlockContext) { l.st.PushStack() }
func (l *parseListener) ExitBlock(ctx *pantlr.BlockContext) {
	stmts := l.popStmts()
	l.st.Push(&ast.BlockStmt{
		BlockPos: l.startPos(ctx),
		Stmts:    stmts,
	})
}

func (l *parseListener) popStmts() []ast.Stmt {
	raw := l.st.PopStack()
	if len(raw) == 0 {
		return nil
	}

	stmts := make([]ast.Stmt, len(raw))
	for i, node := range raw {
		stmts[i] = node.(ast.Stmt)
	}
	return stmts
}

//
// handle rule: expression
//

func (l *parseListener) ExitBinary(ctx *pantlr.BinaryContext) {
	switch {
	// arithmetic
	case l.makeBinOp(ast.OpMul, ctx.MUL()):
	case l.makeBinOp(ast.OpDiv, ctx.DIV()):
	case l.makeBinOp(ast.OpRem, ctx.REM()):
	case l.makeBinOp(ast.OpAdd, ctx.ADD()):
	case l.makeBinOp(ast.OpSub, ctx.SUB()):

	// regex
	case l.makeBinOp(ast.OpFind, ctx.FIND()):
	case l.makeBinOp(ast.OpMatch, ctx.MATCH()):

	// bitwise
	case l.makeBinOp(ast.OpLSH, ctx.LSH()):
	case l.makeBinOp(ast.OpRSH, ctx.RSH()):
	case l.makeBinOp(ast.OpUSH, ctx.USH()):
	case l.makeBinOp(ast.OpBitsAnd, ctx.BWAND()):
	case l.makeBinOp(ast.OpBitsXOR, ctx.XOR()):
	case l.makeBinOp(ast.OpBitsOR, ctx.BWOR()):
	}
}

func (l *parseListener) ExitComp(ctx *pantlr.CompContext) {
	switch {
	case l.makeBinOp(ast.OpCmpEQ, ctx.EQ()):
	case l.makeBinOp(ast.OpCmpEQR, ctx.EQR()):
	case l.makeBinOp(ast.OpCmpNE, ctx.NE()):
	case l.makeBinOp(ast.OpCmpNER, ctx.NER()):
	case l.makeBinOp(ast.OpCmpLT, ctx.LT()):
	case l.makeBinOp(ast.OpCmpLTE, ctx.LTE()):
	case l.makeBinOp(ast.OpCmpGT, ctx.GT()):
	case l.makeBinOp(ast.OpCmpGTE, ctx.GTE()):
	}
}

func (l *parseListener) ExitBool(ctx *pantlr.BoolContext) {
	switch {
	case l.makeBinOp(ast.OpAND, ctx.BOOLAND()):
	case l.makeBinOp(ast.OpOR, ctx.BOOLOR()):
	}
}

func (l *parseListener) ExitInstanceof(ctx *pantlr.InstanceofContext) {
	l.makeBinOp(ast.OpInstanceOf, ctx.INSTANCEOF())
}

func (l *parseListener) ExitConditional(ctx *pantlr.ConditionalContext) {
	exprs := l.st.PopN(3)
	l.st.Push(&ast.CondExpr{
		Check: exprs[0].(ast.Expr),
		Then:  exprs[1].(ast.Expr),
		Else:  exprs[2].(ast.Expr),
	})
}

func (l *parseListener) ExitElvis(ctx *pantlr.ElvisContext) {
	exprs := l.st.PopN(2)
	l.st.Push(&ast.CondExpr{
		Check: exprs[0].(ast.Expr),
		Then:  exprs[1].(ast.Expr),
	})
}

func (l *parseListener) ExitAssignment(ctx *pantlr.AssignmentContext) {
	switch {
	case l.makeAssign(ast.OpAssign, ctx.ASSIGN()):
	case l.makeAssign(ast.OpAdd, ctx.AADD()):
	case l.makeAssign(ast.OpSub, ctx.ASUB()):
	case l.makeAssign(ast.OpMul, ctx.AMUL()):
	case l.makeAssign(ast.OpDiv, ctx.ADIV()):
	case l.makeAssign(ast.OpRem, ctx.AREM()):
	case l.makeAssign(ast.OpBitsAnd, ctx.AAND()):
	case l.makeAssign(ast.OpBitsXOR, ctx.AXOR()):
	case l.makeAssign(ast.OpBitsOR, ctx.AOR()):
	case l.makeAssign(ast.OpLSH, ctx.ALSH()):
	case l.makeAssign(ast.OpRSH, ctx.ARSH()):
	case l.makeAssign(ast.OpUSH, ctx.AUSH()):
	}
}

func (l *parseListener) makeBinOp(kind ast.OpKind, opNode antlr.TerminalNode) bool {
	if opNode == nil {
		return false
	}

	x, y := l.st.Pop2()
	op := &ast.BinOp{
		OpPos: tokPos(l.source, opNode.GetSymbol()),
		Op:    kind,
		X:     x.(ast.Expr),
		Y:     y.(ast.Expr),
	}

	l.st.Push(op)
	return true
}

func (l *parseListener) makeAssign(op ast.OpKind, opNode antlr.TerminalNode) bool {
	if opNode == nil {
		return false
	}

	x, y := l.st.Pop2()
	l.st.Push(&ast.Assign{
		OpPos: tokPos(l.source, opNode.GetSymbol()),
		Op:    op,
		LHS:   x.(ast.Expr),
		RHS:   y.(ast.Expr),
	})
	return true
}

//
// handle rule: unary
// ------------------

func (l *parseListener) ExitPre(ctx *pantlr.PreContext) {
	switch {
	case l.makeUnary(ast.OpInc, true, ctx.INCR()):
	case l.makeUnary(ast.OpDec, true, ctx.DECR()):
	}
}

func (l *parseListener) ExitPost(ctx *pantlr.PostContext) {
	switch {
	case l.makeUnary(ast.OpInc, false, ctx.INCR()):
	case l.makeUnary(ast.OpDec, false, ctx.DECR()):
	}
}

func (l *parseListener) ExitOperator(ctx *pantlr.OperatorContext) {
	switch {
	case l.makeUnary(ast.OpBoolNot, true, ctx.BOOLNOT()):
	case l.makeUnary(ast.OpBitsNot, true, ctx.BWNOT()):
	case l.makeUnary(ast.OpPos, true, ctx.ADD()):
	case l.makeUnary(ast.OpNeg, true, ctx.SUB()):
	}
}

func (l *parseListener) ExitCast(ctx *pantlr.CastContext) {
	to, expr := l.st.Pop2()
	l.st.Push(&ast.CastExpr{
		CastPos: l.startPos(ctx),
		To:      to.(ast.TypeExpr),
		Expr:    expr.(ast.Expr),
	})
}

func (l *parseListener) makeUnary(op ast.OpKind, pre bool, node antlr.TerminalNode) bool {
	if node == nil {
		return false
	}

	expr := l.st.Pop()
	l.st.Push(&ast.UnaryOp{
		OpPos: tokPos(l.source, node.GetSymbol()),
		X:     expr.(ast.Expr),
		Pre:   pre,
		Op:    op,
	})
	return true
}

//
// handle rule: primary
// --------------------

func (l *parseListener) ExitNumeric(ctx *pantlr.NumericContext) {
	tokType := ctx.GetStart().GetTokenType()
	switch tokType {
	case pantlr.PainlessLexerDECIMAL:
		l.pushLit(ctx, ast.LitDecimal, ctx.GetText())

	default:
		l.pushLit(ctx, ast.LitInt, ctx.GetText())
	}
}

func (l *parseListener) ExitTrue(ctx *pantlr.TrueContext)   { l.pushLit(ctx, ast.LitBool, "true") }
func (l *parseListener) ExitFalse(ctx *pantlr.FalseContext) { l.pushLit(ctx, ast.LitBool, "false") }
func (l *parseListener) ExitNull(ctx *pantlr.NullContext)   { l.pushLit(ctx, ast.LitNull, "null") }

func (l *parseListener) ExitString(ctx *pantlr.StringContext) {
	l.pushLit(ctx, ast.LitString, ctx.GetText())
}

func (l *parseListener) ExitRegex(ctx *pantlr.RegexContext) {
	l.pushLit(ctx, ast.LitRegex, ctx.GetText())
}

func (l *parseListener) EnterListinit(ctx *pantlr.ListinitContext) { l.st.PushStack() }
func (l *parseListener) ExitListinit(ctx *pantlr.ListinitContext) {
	var elts []ast.Expr
	nodes := l.st.PopStack()
	if len(nodes) > 0 {
		panic("TODO")
	}

	l.st.Push(&ast.ListInit{Start: l.startPos(ctx), Elts: elts})
}

func (l *parseListener) EnterMapinit(ctx *pantlr.MapinitContext) { l.st.PushStack() }
func (l *parseListener) ExitMapinit(ctx *pantlr.MapinitContext) {
	var fields []ast.Expr
	var values []ast.Expr

	nodes := l.st.PopStack()
	if len(nodes) > 0 {
		panic("TODO")
	}

	l.st.Push(&ast.MapInit{Start: l.startPos(ctx), Fields: fields, Values: values})
}

func (l *parseListener) ExitVariable(ctx *pantlr.VariableContext) {
	l.st.Push(&ast.Ident{
		NamePos: l.startPos(ctx),
		Name:    ctx.GetText(),
		Obj: &ast.Object{
			Kind: ast.VarObj,
			Name: ctx.GetText(),
		},
	})
}

func (l *parseListener) EnterCalllocal(ctx *pantlr.CalllocalContext) { l.st.PushStack() }
func (l *parseListener) ExitCalllocal(ctx *pantlr.CalllocalContext) {
	l.st.Push(&ast.Call{
		CallPos: l.startPos(ctx),
		Kind:    ast.CallFunction,
		ID:      l.newIdent(ctx.ID(), ast.FunObj),
		Args:    l.popArgs(),
	})
}

func (l *parseListener) popArgs() []ast.Expr {
	raw := l.st.PopStack()
	if len(raw) == 0 {
		return nil
	}

	args := make([]ast.Expr, len(raw))
	for i, node := range raw {
		args[i] = node.(ast.Expr)
	}
	return args
}

func (l *parseListener) EnterNewobject(ctx *pantlr.NewobjectContext) { l.st.PushStack() }
func (l *parseListener) ExitNewobject(ctx *pantlr.NewobjectContext) {
	l.st.Push(&ast.Call{
		CallPos: l.startPos(ctx),
		Kind:    ast.CallNew,
		ID:      l.newIdent(ctx.TYPE(), ast.TypObj),
		Args:    l.popArgs(),
	})
}

func (l *parseListener) pushLit(ctx antlr.ParserRuleContext, kind ast.LiteralKind, value string) {
	l.st.Push(&ast.Literal{
		ValuePos: l.startPos(ctx),
		Kind:     kind,
		Value:    value,
	})
}

//
// chain expression handlers
// -------------------------

func (l *parseListener) EnterNewstandardarray(ctx *pantlr.NewstandardarrayContext) { l.st.PushStack() }
func (l *parseListener) ExitNewstandardarray(ctx *pantlr.NewstandardarrayContext) {
	id := l.newIdent(ctx.TYPE(), ast.TypObj)
	start := id.Pos()
	exprs := l.st.PopStack()

	current := ast.TypeExpr(id)
	for _, e := range exprs {
		current = &ast.ArrType{
			Start: start,
			Len:   e.(ast.Expr),
			Elt:   current,
		}
	}

	l.st.Push(&ast.NewArray{
		Typ: current.(*ast.ArrType),
	})
}

func (l *parseListener) EnterNewinitializedarray(ctx *pantlr.NewinitializedarrayContext) {
	l.st.PushStack()
}
func (l *parseListener) ExitNewinitializedarray(ctx *pantlr.NewinitializedarrayContext) {
	id := l.newIdent(ctx.TYPE(), ast.TypObj)
	start := id.Pos()

	init := l.st.PopStack()
	exprs := make([]ast.Expr, len(init))
	for i := range init {
		exprs[i] = init[i].(ast.Expr)
	}

	l.st.Push(&ast.NewArray{
		Typ: &ast.ArrType{
			Start: start,
			Elt:   id,
		},
		Init: exprs,
	})
}

//
// Post-Expressions (postfix, postdot)
// -----------------------------------

func (l *parseListener) EnterCallinvoke(ctx *pantlr.CallinvokeContext) { l.st.PushStack() }
func (l *parseListener) ExitCallinvoke(ctx *pantlr.CallinvokeContext) {
	dot, safeDot := ctx.DOT(), false
	if dot == nil {
		dot, safeDot = ctx.NSDOT(), true
	}

	args := l.popArgs()
	value := l.st.Pop().(ast.Expr)
	l.st.Push(&ast.Call{
		CallPos: tokPos(l.source, dot.GetSymbol()),
		Kind:    ast.CallMethod,
		SafeDot: safeDot,
		Value:   value,
		ID:      l.newIdent(ctx.DOTID(), ast.FieldObj),
		Args:    args,
	})
}

func (l *parseListener) ExitFieldaccess(ctx *pantlr.FieldaccessContext) {
	dot, safeDot := ctx.DOT(), false
	if dot == nil {
		dot, safeDot = ctx.NSDOT(), true
	}

	pos := tokPos(l.source, dot.GetSymbol())

	value := l.st.Pop().(ast.Expr)
	if id := ctx.DOTID(); id != nil {
		l.st.Push(&ast.FieldAccess{
			At:      pos,
			Value:   value,
			SafeDot: safeDot,
			ID:      l.newIdent(id, ast.FieldObj),
		})
	} else {
		i, _ := strconv.Atoi(ctx.DOTINTEGER().GetText())
		l.st.Push(&ast.IdxAccess{
			At:      pos,
			Value:   value,
			SafeDot: safeDot,
			Idx:     i,
		})
	}
}

func (l *parseListener) ExitBraceaccess(ctx *pantlr.BraceaccessContext) {
	accessor := l.st.Pop().(ast.Expr)

	value := l.st.Pop().(ast.Expr)
	l.st.Push(&ast.Access{
		At:       l.startPos(ctx),
		Value:    value,
		Accessor: accessor,
	})
}

func (l *parseListener) ExitDecltype(ctx *pantlr.DecltypeContext) {
	base := l.newIdent(ctx.TYPE(), ast.TypObj)
	start := base.NamePos
	braces := len(ctx.AllLBRACE()) - 1

	var current ast.TypeExpr = base
	for ; braces >= 0; braces-- {
		current = &ast.ArrType{
			Start: start,
			Elt:   current,
		}
	}

	l.st.Push(current)
}

//
// handle lambda
// -------------

func (l *parseListener) EnterLambda(ctx *pantlr.LambdaContext) { l.st.PushStack() }
func (l *parseListener) ExitLambda(ctx *pantlr.LambdaContext) {
	stack := l.st.PopStack()
	rhs := stack.Pop()
	lhs := stack

	var fields []*ast.Field
	if len(lhs) > 0 {
		fields := make([]*ast.Field, len(lhs))
		for i, node := range lhs {
			fields[i] = node.(*ast.Field)
		}
	}

	l.st.Push(&ast.Lambda{
		At:   l.startPos(ctx),
		Body: rhs,
		Args: fields,
	})
}

func (l *parseListener) EnterLamtype(ctx *pantlr.LamtypeContext) { l.st.PushStack() }
func (l *parseListener) ExitLamtype(ctx *pantlr.LamtypeContext) {
	stack := l.st.PopStack()
	if len(stack) > 1 {
		panic("Invalid lambda type stack")
	}

	var t ast.TypeExpr
	if len(stack) > 0 {
		t = stack[0].(ast.TypeExpr)
	}

	l.st.Push(&ast.Field{
		FieldPos: l.startPos(ctx),
		Type:     t,
		ID:       l.newIdent(ctx.ID(), ast.VarObj),
	})
}

//
// handle funcref
// --------------

func (l *parseListener) ExitClassfuncref(ctx *pantlr.ClassfuncrefContext) {
	l.st.Push(&ast.Ref{
		At:    l.startPos(ctx),
		Value: l.newIdent(ctx.TYPE(), ast.TypObj),
		ID:    l.newIdent(ctx.ID(), ast.FieldObj),
	})
}

func (l *parseListener) ExitConstructorfuncref(ctx *pantlr.ConstructorfuncrefContext) {
	l.st.Push(&ast.Ref{
		At:    l.startPos(ctx),
		Value: l.st.Pop().(ast.Expr),
		ID:    l.newIdent(ctx.NEW(), ast.FieldObj),
	})
}

func (l *parseListener) ExitCapturingfuncref(ctx *pantlr.CapturingfuncrefContext) {
	l.st.Push(&ast.Ref{
		At:    l.startPos(ctx),
		Value: l.newIdent(ctx.ID(0), ast.VarObj),
		ID:    l.newIdent(ctx.ID(1), ast.FieldObj),
	})
}

func (l *parseListener) ExitLocalfuncref(ctx *pantlr.LocalfuncrefContext) {
	l.st.Push(&ast.Ref{
		At:    l.startPos(ctx),
		Value: l.newIdent(ctx.THIS(), ast.VarObj),
		ID:    l.newIdent(ctx.ID(), ast.FieldObj),
	})
}

//
// parser utils
// ------------

func (l *parseListener) startPos(ctx antlr.ParserRuleContext) ast.Pos {
	return startPos(l.source, ctx)
}

func (l *parseListener) endPos(ctx antlr.ParserRuleContext) ast.Pos {
	return endPos(l.source, ctx)
}

func (l *parseListener) newIdent(node antlr.TerminalNode, kind ast.ObjKind) *ast.Ident {
	name := node.GetText()
	return &ast.Ident{
		NamePos: tokPos(l.source, node.GetSymbol()),
		Name:    name,
		Obj: &ast.Object{
			Kind: kind,
			Name: name,
		},
	}
}
