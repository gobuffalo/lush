package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/types"
)

type Multiplyer interface {
	Multiply(b interface{}) (interface{}, error)
}

func Multiply(a, b interface{}) (interface{}, error) {
	switch at := a.(type) {
	case Multiplyer:
		return at.Multiply(b)
	case int:
		switch bt := b.(type) {
		case int:
			return at * bt, nil
		case float64:
			return float64(at) * bt, nil
		case types.Integer:
			return at * bt.Int(), nil
		case types.Floater:
			return float64(at) * bt.Float(), nil
		}
	case float64:
		switch bt := b.(type) {
		case float64:
			return at * bt, nil
		case types.Integer:
			return at * float64(bt.Int()), nil
		case types.Floater:
			return at * bt.Float(), nil
		case int:
			return at * float64(bt), nil
		}
	}

	return nil, fmt.Errorf("can't add %T and %T", a, b)
}
