package ast

import (
	"fmt"

	"github.com/gobuffalo/lush/opers"
	"github.com/gobuffalo/lush/types"
)

// BinaryExpr is an expression involving two operands. The operands can either
// be Unary expressions (e.g. !true) or another BinaryExpr for more complicated
// expressions (e.g. 4 + (6 / 2))
type BinaryExpr struct {
	Op   string
	A    Execable
	B    Execable
	Meta Meta
}

func (e BinaryExpr) String() string {
	return fmt.Sprintf("(%s %s %s)", e.A, e.Op, e.B)
}

func (e BinaryExpr) Format(st fmt.State, verb rune) {
	format(e, st, verb)
}

func (e BinaryExpr) MarshalJSON() ([]byte, error) {
	m := map[string]interface{}{
		"A":    e.A,
		"B":    e.B,
		"op":   e.Op,
		"Meta": e.Meta,
	}
	return toJSON(e, m)
}

// Exec applies the operation of the BinaryExpr to the left and right subtrees
// (LHS and RHS) after recursively Exec-ing each subtree. This produces a
// depth-first evaluation order.
func (e BinaryExpr) Exec(c *Context) (interface{}, error) {
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

func (e BinaryExpr) And(c *Context) (bool, error) {
	a, _ := boolExec(e.A, c)
	b, _ := boolExec(e.B, c)
	return a && b, nil
}

func (e BinaryExpr) Or(c *Context) (bool, error) {
	a, _ := boolExec(e.A, c)
	b, _ := boolExec(e.B, c)
	return a || b, nil
}

func (e BinaryExpr) Bool(c *Context) (bool, error) {
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
		return opers.NotEqual(a, b)
	case "~=":
		return opers.Match(a, types.String(b))
	case "<":
		return opers.LessThan(a, b)
	case ">":
		return opers.GreaterThan(a, b)
	case "<=":
		return opers.LessThanEqualTo(a, b)
	case ">=":
		return opers.GreaterThanEqualTo(a, b)
	}
	return false, nil
}

func (e BinaryExpr) Add(c *Context) (interface{}, error) {
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

func (e BinaryExpr) Sub(c *Context) (interface{}, error) {
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

func (e BinaryExpr) Multiply(c *Context) (interface{}, error) {
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

func (e BinaryExpr) Divide(c *Context) (interface{}, error) {
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

func (e BinaryExpr) Modulus(c *Context) (interface{}, error) {
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
