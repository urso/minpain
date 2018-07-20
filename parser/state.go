package parser

import (
	"github.com/urso/minpain/ast"
)

type state struct {
	active  stateStack
	stacks  []stateStack
	stacks0 [8]stateStack
}

type stateStack []ast.Node

func newState() *state {
	s := &state{}
	s.stacks = s.stacks0[:0]
	return s
}

func (st *state) Push(node ast.Node)         { st.active.Push(node) }
func (st *state) Pop() ast.Node              { return st.active.Pop() }
func (st *state) PopN(n int) []ast.Node      { return st.active.PopN(n) }
func (st *state) Pop2() (ast.Node, ast.Node) { return st.active.Pop2() }

func (st *state) ActiveStack() stateStack {
	return st.active
}

func (st *state) PushStack() {
	st.stacks = append(st.stacks, st.active)
	st.active = nil
}

func (st *state) PopStack() stateStack {
	i := len(st.stacks) - 1
	ret := st.active
	st.active, st.stacks = st.stacks[i], st.stacks[:i]
	return ret
}

func (st *stateStack) Push(node ast.Node) {
	*st = append(*st, node)
}

func (st *stateStack) PopN(n int) []ast.Node {
	i := len(*st) - n
	if i < 0 {
		panic("pop from empty stack")
	}

	nodes, rest := (*st)[i:], (*st)[:i]
	(*st) = rest
	return nodes
}

func (st *stateStack) Pop() ast.Node {
	return st.PopN(1)[0]
}

func (st *stateStack) Pop2() (ast.Node, ast.Node) {
	nodes := st.PopN(2)
	return nodes[0], nodes[1]
}
