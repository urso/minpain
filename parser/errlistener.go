package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/urso/minpain/ast"
)

type errorListener struct {
	op     string
	source string
	errs   multiError
}

var _ antlr.ErrorListener = &errorListener{}

func (e *errorListener) Error() error {
	return e.errs.Err()
}

func (l *errorListener) SyntaxError(
	recognizer antlr.Recognizer,
	offendingSymbol interface{},
	line, column int, msg string,
	_ antlr.RecognitionException,
) {
	l.errs.Add(&ParseError{
		op:  l.op,
		pos: ast.Pos{Source: l.source, Line: line, Column: column},
		msg: msg,
	})
}

func (l *errorListener) ReportAmbiguity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex, stopIndex int, exact bool,
	ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet,
) {
}

func (l *errorListener) ReportAttemptingFullContext(
	recognizer antlr.Parser,
	dfa *antlr.DFA, startIndex,
	stopIndex int,
	conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet,
) {
	// nothing
}

func (l *errorListener) ReportContextSensitivity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex, stopIndex, prediction int,
	configs antlr.ATNConfigSet,
) {
	// nothing
}
