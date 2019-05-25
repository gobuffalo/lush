package lessthan

import "fmt"

func String(at string, b interface{}) (bool, error) {
	switch bt := b.(type) {
	case string:
		return at < bt, nil
	case fmt.Stringer:
		return at < bt.String(), nil
	}
	return false, fmt.Errorf("can not compare %T with %T", at, b)
}
