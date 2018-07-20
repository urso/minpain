package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type traceListener struct {
}

var _ antlr.ParseTreeListener = &traceListener{}

func (l *traceListener) VisitTerminal(node antlr.TerminalNode) {
}

func (l *traceListener) VisitErrorNode(node antlr.ErrorNode) {
}

func (l *traceListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
}

func (l *traceListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}
