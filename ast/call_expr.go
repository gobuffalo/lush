package ast

import (
	"errors"
	"reflect"
)

type AccessExpr struct {
	Callee   Execable
	Property string
}

func (c AccessExpr) Exec(ctx *Context) (interface{}, error) {
	res, err := c.Callee.Exec(ctx)
	if err != nil {
		return res, err
	}

	rv := reflect.ValueOf(res)
	if rv.Kind() != reflect.Struct {
		return nil, errors.New("Attempt to access properties on a non-struct type")
	}
	val := rv.FieldByName(c.Property).Interface()
	return val, nil
}
