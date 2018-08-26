package print

type node interface{}

type kw struct {
	keyword string
	value   node
}

type sexp struct {
	name     string
	keywords []kw
	nodes    []node
}

func newSexp(name string, nodes ...interface{}) *sexp {
	s := &sexp{name: name}
	return s.Add(nodes...)
}

func (s *sexp) Add(nodes ...interface{}) *sexp {
	for _, node := range nodes {
		switch v := node.(type) {
		case kw:
			s.keywords = append(s.keywords, v)
		default:
			s.nodes = append(s.nodes, v)
		}
	}
	return s
}
