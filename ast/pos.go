package ast

import "fmt"

type Pos struct {
	Source       string
	Line, Column int
}

func (p Pos) String() string {
	if p.Source == "" {
		return fmt.Sprintf("%v:%v", p.Line, p.Column)
	}
	return fmt.Sprintf("%v:%v:%v", p.Source, p.Line, p.Column)
}
