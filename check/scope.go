package check

import (
	"fmt"

	"github.com/urso/minpain/types"
)

type Scope struct {
	name     string
	idx      int
	kind     ScopeKind
	parent   *Scope
	children []*Scope
	syms     map[string]Object

	classes    map[string]types.Class
	interfaces map[string]types.Interface
}

type ScopeKind uint8

//go:generate stringer -type=ScopeKind
const (
	ScopeUniversal ScopeKind = iota
	ScopePrelude             // user provided prelude
	ScopeScript
	ScopeFunction // function scope
	ScopeBlock
)

func NewScope(name string, kind ScopeKind, parent *Scope) *Scope {
	s := &Scope{name: name, idx: -1, kind: kind, parent: parent}
	if parent != nil {
		k := parent.kind
		if k != ScopeUniversal && k != ScopePrelude {
			s.idx = len(parent.children)
			parent.children = append(parent.children, s)
		}
	}

	return s
}

func (s *Scope) Name() string {
	return s.name
}

func (s *Scope) FullName() string {
	parent := s.Parent()
	if parent == nil {
		return s.Name()
	}
	return fmt.Sprintf("%v.%v.%v", parent.FullName(), s.idx, s.Name())
}

func (s *Scope) Kind() ScopeKind {
	return s.kind
}

func (s *Scope) Parent() *Scope {
	return s.parent
}

func (s *Scope) LookupDirect(sym string) Object {
	if s.syms == nil {
		return nil
	}
	return s.syms[sym]
}

func (s *Scope) LookupParents(sym string) Object {
	if s.parent == nil {
		return nil
	}
	return s.parent.Lookup(sym)
}

func (s *Scope) Lookup(sym string) Object {
	if s.syms != nil {
		if obj := s.syms[sym]; obj != nil {
			return obj
		}
	}
	return s.LookupParents(sym)
}

func (s *Scope) add(obj Object) {
	if s.syms == nil {
		s.syms = map[string]Object{}
	}
	s.syms[obj.Name()] = obj
}

func (s *Scope) Each(fn func(Object)) {
	for _, obj := range s.syms {
		fn(obj)
	}

	for _, child := range s.children {
		child.Each(fn)
	}
}
