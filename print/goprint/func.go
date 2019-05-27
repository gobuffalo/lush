package goprint

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/lush/ast"
)

func (c Printer) astFunc(f ast.Func) error {
	fmt.Fprintf(c, "func(")
	var args []string
	for _, a := range f.Arguments {
		args = append(args, a.String()+" interface{}")
	}
	fmt.Fprintf(c, strings.Join(args, ", "))
	fmt.Fprintf(c, ")")
	if f.Block != nil {
		r, ok := returnFinder(f.Block)
		if ok {
			var lines []string
			for i := 0; i < len(r.Nodes); i++ {
				lines = append(lines, "interface{}")
			}
			fmt.Fprintf(c, "(%s)", strings.Join(lines, ", "))
		}
		fmt.Fprintf(c, "{")
		if err := c.astBlock(f.Block); err != nil {
			return err
		}
		fmt.Fprintf(c, "}")

	}
	fmt.Fprintln(c)
	return nil
}

func returnFinder(s ast.Node) (ast.Return, bool) {
	switch t := s.(type) {
	case ast.Nodes:
		for _, s := range t {
			r, ok := returnFinder(s)
			if ok {
				return r, ok
			}
		}
	case ast.Return:
		return t, true
	case *ast.Block:
		return returnFinder(t.Nodes)
	default:
	}

	return ast.Return{}, false
}
