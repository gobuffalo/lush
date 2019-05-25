package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/types"
)

// Modulus attempts to take the "modulus" of type `a` with type `b`.
// Supports:
//	* int
//	* faces.Modulus
//	* types.Integer
func Modulus(a, b interface{}) (int, error) {
	switch at := a.(type) {
	case faces.Modulus:
		return at.Modulus(b)
	case int:
		switch bt := b.(type) {
		case int:
			if bt == 0 {
				return 0, nil
			}
			return at % bt, nil
		case types.Integer:
			bi := bt.Int()
			if bi == 0 {
				return 0, nil
			}
			return at % bi, nil
		}
	case types.Integer:
		switch bt := b.(type) {
		case int:
			if bt == 0 {
				return 0, nil
			}
			return at.Int() % bt, nil
		case types.Integer:
			bi := bt.Int()
			if bi == 0 {
				return 0, nil
			}
			return at.Int() % bi, nil
		}

	}

	return 0, fmt.Errorf("can't modululerize %T and %T", a, b)
}
