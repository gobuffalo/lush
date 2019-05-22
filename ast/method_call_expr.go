package ast

import (
	"errors"
	"reflect"
)

// MethodCallExpr represents the invocation of a method on something that
// returns a struct.
type MethodCallExpr struct {
	Callee Execable
	Method string
	Args   []Execable
}

// Exec invokes the method referred to by Method using the arguments derived
// from Args, after evaluation them sequentially (left-to-right order)
func (m *MethodCallExpr) Exec(ctx *Context) (interface{}, error) {
	res, err := m.Callee.Exec(ctx)
	if err != nil {
		return res, err
	}

	rv := reflect.Indirect(reflect.ValueOf(res))
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("Attempt to call a method on a non-struct type")
	}

	meth := rv.MethodByName(m.Method)
	callRes := meth.Call([]reflect.Value{})
	return callRes[0].Interface(), nil
}
