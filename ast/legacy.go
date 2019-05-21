package ast

import (
	"context"
	"fmt"
	"time"
)

type legacyContext interface {
	context.Context
	New() legacyContext
	Has(key string) bool
	Set(key string, value interface{})
}

type legacyHelperContext interface {
	legacyContext
	Block() (string, error)
	BlockWith(legacyContext) (string, error)
	HasBlock() bool
	Render(s string) (string, error)
}

var _ legacyHelperContext = legacyCtx{}

type legacyCtx struct {
	Ctx *Context
}

func (l legacyCtx) Value(i interface{}) interface{} {
	return l.Ctx.Value(i)
}

func (l legacyCtx) Set(k string, v interface{}) {
	l.Ctx.Set(k, v)
}

func (l legacyCtx) Deadline() (time.Time, bool) {
	return l.Ctx.Deadline()
}

func (l legacyCtx) Done() <-chan struct{} {
	return l.Ctx.Done()
}

func (l legacyCtx) Err() error {
	return l.Ctx.Err()
}

func (l legacyCtx) Has(k string) bool {
	return l.Ctx.Has(k)
}

func (l legacyCtx) HasBlock() bool {
	return l.Ctx.Block != nil
}

func (l legacyCtx) New() legacyContext {
	return legacyCtx{Ctx: l.Ctx.Clone()}
}

func (l legacyCtx) Block() (string, error) {
	return l.BlockWith(l)
}

func (l legacyCtx) BlockWith(lc legacyContext) (string, error) {
	if l.Ctx.Block == nil {
		return "", fmt.Errorf("no block given")
	}
	c, ok := lc.(legacyCtx)
	if !ok {
		return "", fmt.Errorf("expected legacy context")
	}
	i, err := l.Ctx.Block.Exec(c.Ctx)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", i), nil
}

func (l legacyCtx) Render(s string) (string, error) {
	return s, nil
}
