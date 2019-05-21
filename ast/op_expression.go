package ast

import (
	"fmt"
	"regexp"

	"github.com/google/go-cmp/cmp"
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
	return toJSON("ast.OpExpression", m)
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
		res := cmp.Equal(a, b)
		return res, nil
	case "!=":
		res := cmp.Equal(a, b)
		return !res, nil
	case "~=":
		sb, ok := b.(string)
		if !ok {
			return false, e.Meta.Errorf("expected string got %T", b)
		}
		rx, err := regexp.Compile(sb)
		if err != nil {
			return false, err
		}
		s, ok := a.(string)
		if !ok {
			return false, e.Meta.Errorf("expected string got %T", a)
		}
		return rx.MatchString(s), nil
	case "<":
		return fmt.Sprint(a) < fmt.Sprint(b), nil
	case ">":
		return fmt.Sprint(a) > fmt.Sprint(b), nil
	case "<=":
		return fmt.Sprint(a) <= fmt.Sprint(b), nil
	case ">=":
		return fmt.Sprint(a) >= fmt.Sprint(b), nil
	}
	return false, nil
}

func (e OpExpression) Add(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, err
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, err
	}

	if fl, err := ints(a, b); err == nil {
		var f int
		for _, x := range fl {
			f += x
		}
		return f, nil
	}

	if fl, err := floats(a, b); err == nil {
		var f float64
		for _, x := range fl {
			f += x
		}
		return f, nil
	}

	if fl, err := stringSlice(c, a, b); err == nil {
		var f string
		for _, x := range fl {
			f += x
		}
		return f, nil
	}

	if at, ok := a.([]interface{}); ok {
		if bt, ok := b.([]interface{}); ok {
			return append(at, bt...), nil
		}
	}

	return nil, e.Meta.Errorf("can not add %T and %T", a, b)
}

func (e OpExpression) Sub(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, err
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, err
	}

	switch at := a.(type) {
	case int:
		switch bt := b.(type) {
		case int:
			return at - bt, nil
		case float64:
			return float64(at) - bt, nil
		}
	case float64:
		switch bt := b.(type) {
		case int:
			return at - float64(bt), nil
		case float64:
			return at - bt, nil
		}
	}

	return nil, e.Meta.Errorf("can not subtract %T and %T", a, b)
}

func (e OpExpression) Multiply(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, err
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, err
	}

	switch at := a.(type) {
	case int:
		switch bt := b.(type) {
		case int:
			return at * bt, nil
		case float64:
			return float64(at) * bt, nil
		}
	case float64:
		switch bt := b.(type) {
		case int:
			return at * float64(bt), nil
		case float64:
			return at * bt, nil
		}
	}

	return nil, e.Meta.Errorf("can not multiply %T and %T", a, b)
}

func (e OpExpression) Divide(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, err
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, err
	}

	switch at := a.(type) {
	case int:
		switch bt := b.(type) {
		case int:
			return float64(at) / float64(bt), nil
		case float64:
			return float64(at) / bt, nil
		}
	case float64:
		switch bt := b.(type) {
		case int:
			return at / float64(bt), nil
		case float64:
			return at / bt, nil
		}
	}

	return nil, e.Meta.Errorf("can not divide %T and %T", a, b)
}

func (e OpExpression) Modulus(c *Context) (interface{}, error) {
	a, err := exec(c, e.A)
	if err != nil {
		return nil, err
	}

	b, err := exec(c, e.B)
	if err != nil {
		return nil, err
	}

	switch at := a.(type) {
	case int:
		switch bt := b.(type) {
		case int:
			if bt == 0 {
				return 0, nil
			}
			return at % bt, nil
		}
	}

	return nil, e.Meta.Errorf("can not modulus %T and %T", a, b)
}
