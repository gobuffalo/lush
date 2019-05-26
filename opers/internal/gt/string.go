package gt

import "fmt"

func String(at string, b interface{}) (bool, error) {
	switch bt := b.(type) {
	case string:
		return at > bt, nil
	case fmt.Stringer:
		return at > bt.String(), nil
	}
	return false, Cant(at, b)
}
