package add

import "fmt"

func String(at string, b interface{}) (interface{}, error) {
	switch bt := b.(type) {
	case string:
		return at + bt, nil
	case fmt.Stringer:
		return at + bt.String(), nil
	}
	return 0, Cant(at, b)
}
