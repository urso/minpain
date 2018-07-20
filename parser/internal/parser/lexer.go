package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Lexer struct {
	*PainlessLexer
}

type lexerBase struct {
	*antlr.BaseLexer

	// Check against the current whitelist to determine whether a token is a type
	// or not. Called by the {@code TYPE} token defined in {@code PainlessLexer.g4}.
	// See also
	// <a href="https://en.wikipedia.org/wiki/The_lexer_hack">The lexer hack</a>.
	IsType func(string) bool

	// SlashIsRegex true if the slash symbol `/` is the beginning of a regex.
	// In case of a division operator false is returned.
	SlashIsRegex func() bool

	current antlr.Token
}

func NewLexer(input antlr.CharStream, isType func(string) bool) *Lexer {
	l := &Lexer{}
	l.PainlessLexer = NewPainlessLexer(input)
	l.PainlessLexer.IsType = isType
	l.PainlessLexer.SlashIsRegex = l.slashIsRegex
	return l
}

func (l *Lexer) NextToken() antlr.Token {
	l.current = l.PainlessLexer.NextToken()
	return l.current
}

func (l *Lexer) slashIsRegex() bool {
	lastToken := l.current
	if lastToken == nil {
		return true
	}

	switch lastToken.GetTokenType() {
	case PainlessLexerRBRACE,
		PainlessLexerRP,
		PainlessLexerOCTAL,
		PainlessLexerHEX,
		PainlessLexerINTEGER,
		PainlessLexerDECIMAL,
		PainlessLexerID,
		PainlessLexerDOTINTEGER,
		PainlessLexerDOTID:
		return false
	default:
		return true
	}
}
