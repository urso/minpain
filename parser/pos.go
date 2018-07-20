package parser

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/urso/minpain/ast"
)

func tokPos(source string, tok antlr.Token) ast.Pos {
	return ast.Pos{Source: source, Line: tok.GetLine(), Column: tok.GetColumn()}
}

func startPos(source string, ctx antlr.ParserRuleContext) ast.Pos {
	return tokPos(source, ctx.GetStart())
}

func endPos(source string, ctx antlr.ParserRuleContext) ast.Pos {
	return tokPos(source, ctx.GetStop())
}
