package ast

import (
	"fmt"
)

type Goroutine struct {
	Call Call
}

func NewGoroutine(c Call) (Goroutine, error) {
	return Goroutine{Call: c}, nil
}

func (g Goroutine) String() string {
	return fmt.Sprintf("go %s", g.Call)
}

func (g Goroutine) Format(st fmt.State, verb rune) {
	format(g, st, verb)
}

func (g Goroutine) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"Call": g.Call,
	}

	return toJSON(g, m)
}

func (g Goroutine) Visit(c *Context) (interface{}, error) {
	return g.Call.Visit(c)
}
