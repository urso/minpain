package print

import (
	"fmt"

	"github.com/urso/minpain/ast"
)

func Print(node ast.Node) {
	contents := pp(nil, "", toSexp(node))
	fmt.Println(string(contents))
}

func pp(buf []byte, indent string, in interface{}) []byte {
	if l := len(buf); l > 0 && buf[l-1] == '\n' {
		buf = append(buf, indent...)
	}

	switch v := in.(type) {
	default:
		return append(buf, fmt.Sprintf("%v", v)...)

	case *sexp:
		buf = append(buf, fmt.Sprintf("(%v", v.name)...)
		indent += "   "
		kwIndent := indent + "  "
		pad := v.name != ""
		for _, kw := range v.keywords {
			if pad {
				buf = append(buf, fmt.Sprintf("\n%v", indent)...)
			}
			buf = append(buf, kw.keyword...)
			buf = pp(append(buf, ": "...), kwIndent, kw.value)
			pad = true
		}
		for _, node := range v.nodes {
			if pad {
				buf = append(buf, '\n')
			}
			buf = pp(buf, indent, node)
			pad = true
		}
		return append(buf, ')')
	}
}
