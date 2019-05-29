package ast

import (
	"context"
	"fmt"
	"io"
	"sync"

	"github.com/gobuffalo/lush/builtins"
)

// type *Runtime interface {
// 	context.Runtime
// 	Clone() *Runtime
// 	Has(key string) bool
// 	Set(key string, value interface{})
// }

//
// type HelperContext interface {
// 	Runtime
// 	Block() (interface{}, error)
// 	BlockWith(Runtime) (interface{}, error)
// 	HasBlock() bool
// 	Exec((c *Runtime)) (interface{}, error)
// }

func NewContext(ctx context.Context, w io.Writer) *Runtime {
	c := &Runtime{
		Context: ctx,
		Writer:  w,
		wg:      &sync.WaitGroup{},
	}

	c.Imports.Store("fmt", builtins.NewFmt(w))
	c.Imports.Store("strings", builtins.Strings{})
	c.Imports.Store("time", builtins.Time{})
	c.Imports.Store("sort", builtins.Time{})

	c.Set("error", fmt.Errorf)
	return c
}

type Runtime struct {
	context.Context
	io.Writer
	Block   *Block
	data    sync.Map
	Imports sync.Map
	wg      *sync.WaitGroup
}

func (c *Runtime) Clone() *Runtime {
	fhc := NewContext(c, c.Writer)
	fhc.Context = c
	fhc.Block = c.Block
	fhc.wg = c.wg
	c.Imports.Range(func(k, v interface{}) bool {
		fhc.Imports.Store(k, v)
		return true
	})
	c.data.Range(func(k, v interface{}) bool {
		fhc.data.Store(k, v)
		return true
	})

	return fhc
}

func (c *Runtime) Has(key string) bool {
	return c.Value(key) != nil
}

func (c *Runtime) Set(key string, value interface{}) {
	c.data.Store(key, value)
}

type upable interface {
	setup(key string, value interface{})
}

func (c *Runtime) setup(key string, value interface{}) {
	c.data.Store(key, value)
	if c.Context.Value(key) != nil {
		if up, ok := c.Context.(upable); ok {
			up.setup(key, value)
		}
	}
}

func (f *Runtime) Value(key interface{}) interface{} {
	v, ok := f.data.Load(key)
	if ok {
		return v
	}
	return f.Context.Value(key)
}
