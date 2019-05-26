package modulus

import (
	"github.com/gobuffalo/lush/faces"
)

// Int ...
func Int(at int, b interface{}) (int, error) {
	var bi int
	switch bt := b.(type) {
	case int:
		bi = bt
	case faces.Int:
		bi = bt.Int()
	default:
		return 0, Cant(at, b)
	}
	if b == 0 {
		return 0, nil
	}
	return at % bi, nil
}
