package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/types"
)

// Modululerize defines an interface to support the
// take the "modulus" of one type from an another.
type Modululerize interface {
	Modulus(b interface{}) (int, error)
}

// Modulus attempts to take the "modulus" of type `a` with type `b`.
// Supports:
//	* int
//	* Modululerize
//	* types.Integer
func Modulus(a, b interface{}) (int, error) {
	switch at := a.(type) {
	case Modululerize:
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
