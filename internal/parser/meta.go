package parser

import (
	"github.com/gobuffalo/lush/ast"
)

type metable interface {
	SetMeta(ast.Meta)
}

func setMeta(s interface{}, c *current) {
	m, ok := s.(metable)
	if !ok {
		return
	}
	m.SetMeta(meta(c))
}

func meta(c *current) ast.Meta {
	var fn string
	f, ok := c.globalStore["filename"]
	if ok {
		fn, _ = f.(string)
	}
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
	p.cur.globalStore["filename"] = p.filename
	return nil
}
