package ast

import (
	"fmt"

	"github.com/gobuffalo/lush/opers"
	"github.com/gobuffalo/lush/types"
)

func NewOpExpression(a Statement, op string, b Statement) (*OpExpression, error) {
	e := OpExpression{
		format: "%s %s %s",
		A:      a,
		B:      b,
		Op:     op,
	}
	return &e, nil
}

func NewPopExpression(a Statement, op string, b Statement) (*OpExpression, error) {
	e, err := NewOpExpression(a, op, b)
	e.format = "(%s %s %s)"
	return e, err
}

type OpExpression struct {
	A    Statement
	B    Statement
	Op   string
	Meta Meta

	format string
}

func (e OpExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", e.A, e.Op, e.B)
}

func (e OpExpression) Format(st fmt.State, verb rune) {
	format(e, st, verb)
}

func (e OpExpression) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"A":      e.A,
		"B":      e.B,
		"op":     e.Op,
		"Meta":   e.Meta,
		"Format": e.format,
	}
	return toJSON(e, m)
}

func (e OpExpression) Exec(c *Context) (interface{}, error) {
	switch e.Op {
	case "==", "!=", "~=", "<", ">", "<=", ">=", "&&":
		return e.Bool(c)
	case "+":
		return e.Add(c)
	case "-":
		return e.Sub(c)
	case "*":
		return e.Multiply(c)
	case "/":
		return e.Divide(c)
	case "%":
		return e.Modulus(c)
	}

	return nil, nil
}

func (e OpExpression) And(c *Context) (bool, error) {
	a, _ := boolExec(e.A, c)
	b, _ := boolExec(e.B, c)
	return a && b, nil
}

func (e OpExpression) Or(c *Context) (bool, error) {
	a, _ := boolExec(e.A, c)
	b, _ := boolExec(e.B, c)
	return a || b, nil
}

func (e OpExpression) Bool(c *Context) (bool, error) {
	if e.Op == "&&" {
		return e.And(c)
	}

	if e.Op == "||" {
		return e.Or(c)
	}

	a, _ := exec(c, e.A)
	if ia, ok := a.(interfacer); ok {
		a = ia.Interface()
	}

	b, _ := exec(c, e.B)
	if ib, ok := b.(interfacer); ok {
		b = ib.Interface()
	}

	switch e.Op {
	case "==":
		return opers.Equal(a, b)
	case "!=":
		res, err := opers.Equal(a, b)
		return !res, err
	case "~=":
		return opers.Match(a, types.Value(b))
	case "<":
		return opers.LessThan(a, b)
	case ">":
		return types.Value(a) > types.Value(b), nil
	case "<=":
		return types.Value(a) <= types.Value(b), nil
	case ">=":
		return types.Value(a) >= types.Value(b), nil
	}
	return false, nil
}

func (e OpExpression) Add(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	i, err := opers.Add(a, b)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}
	return i, nil
}

func (e OpExpression) Sub(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	i, err := opers.Sub(a, b)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}
	return i, nil
}

func (e OpExpression) Multiply(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	i, err := opers.Multiply(a, b)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}
	return i, nil
}

func (e OpExpression) Divide(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	i, err := opers.Divide(a, b)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}
	return i, nil
}

func (e OpExpression) Modulus(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}

	i, err := opers.Modulus(a, b)
	if err != nil {
		return nil, e.Meta.Wrap(err)
	}
	return i, nil
}
