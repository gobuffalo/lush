package ast

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/gobuffalo/lush/builtins"
)

// type *Context interface {
// 	context.Context
// 	Clone() *Context
// 	Has(key string) bool
// 	Set(key string, value interface{})
// }

//
// type HelperContext interface {
// 	Context
// 	Block() (interface{}, error)
// 	BlockWith(Context) (interface{}, error)
// 	HasBlock() bool
// 	Exec((c *Context)) (interface{}, error)
// }

func NewContext(ctx context.Context, w io.Writer) *Context {
	c := &Context{
		Context: ctx,
		Writer:  w,
	}

	c.Set("fmt", builtins.NewFmt(w))
	c.Set("strings", builtins.Strings{})
	c.Set("time", builtins.Time{})
	c.Set("error", fmt.Errorf)
	return c
}

type Context struct {
	context.Context
	io.Writer
	data  sync.Map
	Block *Block
}

func (c *Context) Clone() *Context {
	fhc := NewContext(c, c.Writer)
	fhc.Context = c
	fhc.Block = c.Block
	c.data.Range(func(k, v interface{}) bool {
		fhc.data.Store(k, v)
		return true
	})

	return fhc
}

func (c *Context) Has(key string) bool {
	return c.Value(key) != nil
}

func (c *Context) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

type upable interface {
	setup(key string, value interface{})
}

func (c *Context) setup(key string, value interface{}) {
	c.data.Store(key, value)
	if c.Context.Value(key) != nil {
		if up, ok := c.Context.(upable); ok {
			up.setup(key, value)
		}
	}
}

func (f *Context) Value(key interface{}) interface{} {
	v, ok := f.data.Load(key)
	if ok {
		return v
	}
	return f.Context.Value(key)
}
