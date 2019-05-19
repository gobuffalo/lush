package ast

import "fmt"

type Holder struct {
	Value interface{}
}

func (v Holder) Bool(*Context) (bool, error) {
	return v.Value != nil, nil
}

func (v Holder) String() string {
	return fmt.Sprint(v.Value)
}

func (v Holder) Exec(c *Context) (interface{}, error) {
	return v.Value, nil
}
