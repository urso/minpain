package parser

// Generate lexer
// --------------

// Generate code
//go:generate go run gentool.go antlr -o internal/parser -DsuperClass=lexerBase PainlessLexer.g4

// Fixup generated code
//go:generate go run gentool.go replace *lexerBase lexerBase internal/parser/painless_lexer.go

// Generate Parser
// ---------------

//go:generate go run gentool.go antlr -o internal/parser PainlessParser.g4
