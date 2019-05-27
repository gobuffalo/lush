package ast

import (
	"fmt"
)

type Visitor func(...Statement) error

type Visitable interface {
	Visit(Visitor) error
}

func Walk(st Statement, v Visitor) error {
	if st == nil {
		return fmt.Errorf("can not walk nil Statement")
	}
	vs, ok := st.(Visitable)
	if !ok {
		return fmt.Errorf("expected Visitable got %T", vs)
	}
	return vs.Visit(v)
}

func flatten(ii []interface{}) []interface{} {
	var res []interface{}
	for _, i := range ii {
		if i == nil {
			continue
		}
		switch t := i.(type) {
		case []interface{}:
			res = append(res, flatten(t)...)
		case Noop:
		default:
			res = append(res, i)
		}
	}

	return res
}
