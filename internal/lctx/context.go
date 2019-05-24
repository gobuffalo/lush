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

type upable interface {
	setup(key string, value interface{})
}
