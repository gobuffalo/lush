package opers

import "fmt"

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
		case Integer:
			return at + bt.Int(), nil
		case Float:
			return float64(at) + bt.Float(), nil
		case float64:
			return float64(at) + bt, nil
		}
	case float64:
		switch bt := b.(type) {
		case float64:
			return at + bt, nil
		case Integer:
			return at + float64(bt.Int()), nil
		case Float:
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
			// default:
			// 	fmt.Sprintf("bt %T", bt)
		}
		// default:
		// 	panic(fmt.Sprintf("at %T", at))
	}

	return nil, fmt.Errorf("can't add %T and %T", a, b)
}
