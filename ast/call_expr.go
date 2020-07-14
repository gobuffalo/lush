package ast

import (
	"errors"
	"fmt"
	"reflect"
)

type AccessExpr struct {
	Callee   Execable
	Property string
	Meta     Meta
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

func (c AccessExpr) String() string {
	fmt.Println("AccessExpr", c.Callee, c.Property, c.Meta)
	return ""
}
