package parser

import (
	"fmt"

	"github.com/gobuffalo/lush/ast"
)

func meta(c *current) ast.Meta {
	fmt.Printf(">>>TODO internal/parser/meta.go:10: c (%T): %v \n", c, c)
	var fn string
	// f, ok := c.globalStore["filename"]
	// if ok {
	// 	fn, _ = f.(string)
	// }
	m := ast.Meta{
		Filename: fn,
		Line:     c.pos.line,
		Col:      c.pos.col,
		Offset:   c.pos.offset,
		Original: string(c.text),
	}
	return m
}

func MetaOption(p *parser) Option {
	fmt.Printf(">>>TODO internal/parser/meta.go:24: p (%T): %v \n", p, p)
	// p.cur.globalStore["filename"] = p.filename
	return nil
}
