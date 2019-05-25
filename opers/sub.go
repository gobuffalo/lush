package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/types"
)

// Suber defines an interface to support the
// "subtracting" of one type from an another.
type Suber interface {
	Sub(b interface{}) (interface{}, error)
}

// Sub attempts to "subtract" type `b` from type `a`.
// Supports:
//	* int
//	* float64
//	* Suber
//	* types.Floater
//	* types.Integer
func Sub(a, b interface{}) (interface{}, error) {
	switch at := a.(type) {
	case Suber:
		return at.Sub(b)
	case int:
		switch bt := b.(type) {
		case int:
			return at - bt, nil
		case float64:
			return float64(at) - bt, nil
		case types.Integer:
			return at - bt.Int(), nil
		case types.Floater:
			return float64(at) - bt.Float(), nil
		}
	case float64:
		switch bt := b.(type) {
		case float64:
			return at - bt, nil
		case types.Integer:
			return at - float64(bt.Int()), nil
		case types.Floater:
			return at - bt.Float(), nil
		case int:
			return at - float64(bt), nil
		}
	case types.Integer:
		a := at.Int()
		switch bt := b.(type) {
		case int:
			return a - bt, nil
		case float64:
			return float64(a) - bt, nil
		case types.Integer:
			return a - bt.Int(), nil
		case types.Floater:
			return float64(a) - bt.Float(), nil
		}
	case types.Floater:
		a := at.Float()
		switch bt := b.(type) {
		case float64:
			return a - bt, nil
		case types.Integer:
			return a - float64(bt.Int()), nil
		case types.Floater:
			return a - bt.Float(), nil
		case int:
			return a - float64(bt), nil
		}
	}

	return nil, fmt.Errorf("can't subtract %T and %T", a, b)
}
