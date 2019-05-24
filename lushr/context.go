package lushr

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/gobuffalo/lush/ast"
	"github.com/gobuffalo/lush/builtins"
	"github.com/gobuffalo/lush/internal/lctx"
)

type Context = lctx.Context

func NewContext(ctx context.Context, w io.Writer) *Ctx {
	c := &Ctx{
		Context: ctx,
		Writer:  w,
	}

	c.Imports.Store("fmt", builtins.NewFmt(w))
	c.Imports.Store("strings", builtins.Strings{})
	c.Imports.Store("time", builtins.Time{})

	c.Set("error", fmt.Errorf)
	return c
}

type Ctx struct {
	context.Context
	io.Writer
	data    sync.Map
	Block   *ast.Block
	wg      sync.WaitGroup
	Imports sync.Map
}

func (c *Ctx) Clone() Context {
	fhc := NewContext(c, c.Writer)
	fhc.wg = c.wg
	fhc.Context = c
	fhc.Block = c.Block
	fhc.Imports = c.Imports
	c.data.Range(func(k, v interface{}) bool {
		fhc.data.Store(k, v)
		return true
	})

	return fhc
}

func (c *Ctx) Has(key string) bool {
	return c.Value(key) != nil
}

func (c *Ctx) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

type upable interface {
	setup(key string, value interface{})
}

func (c *Ctx) setup(key string, value interface{}) {
	c.data.Store(key, value)
	if c.Context.Value(key) != nil {
		if up, ok := c.Context.(upable); ok {
			up.setup(key, value)
		}
	}
}

func (f *Ctx) Value(key interface{}) interface{} {
	v, ok := f.data.Load(key)
	if ok {
		return v
	}
	return f.Context.Value(key)
}
