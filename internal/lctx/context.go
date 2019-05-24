package lctx

import (
	"context"
	"io"
)

type Context interface {
	context.Context
	io.Writer
	Clone() Context
	Has(key string) bool
	Set(key string, value interface{})
}

//
// type HelperContext interface {
// 	Context
// 	Block() (interface{}, error)
// 	BlockWith(Context) (interface{}, error)
// 	HasBlock() bool
// 	Exec((c *Context)) (interface{}, error)
// }

type upable interface {
	setup(key string, value interface{})
}
