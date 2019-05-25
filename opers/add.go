package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/types"
)

type Adder interface {
	Add(b interface{}) (interface{}, error)
}

func Add(a, b interface{}) (interface{}, error) {
	switch at := a.(type) {
	case Adder:
		return at.Add(b)
	case int:
		switch bt := b.(type) {
		case int:
			return at + bt, nil
		case float64:
			return float64(at) + bt, nil
		case types.Integer:
			return at + bt.Int(), nil
		case types.Floater:
			return float64(at) + bt.Float(), nil
		}
	case float64:
		switch bt := b.(type) {
		case float64:
			return at + bt, nil
		case types.Integer:
			return at + float64(bt.Int()), nil
		case types.Floater:
			return at + bt.Float(), nil
		case int:
			return at + float64(bt), nil
		}
	case []interface{}:
		switch bt := b.(type) {
		case []interface{}:
			return append(at, bt...), nil
		default:
			return append(at, bt), nil
		}
	case string:
		switch bt := b.(type) {
		case string:
			return at + bt, nil
		}
	}

	return nil, fmt.Errorf("can't add %T and %T", a, b)
}
