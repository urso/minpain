package parser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type errStrategy struct {
	*antlr.DefaultErrorStrategy
}

func newErrStrategy() *errStrategy {
	return &errStrategy{
		DefaultErrorStrategy: antlr.NewDefaultErrorStrategy(),
	}
}
