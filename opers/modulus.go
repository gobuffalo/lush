package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/types"
)

type Modululerize interface {
	Modulus(b interface{}) (int, error)
}

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
	}

	return 0, fmt.Errorf("can't modululerize %T and %T", a, b)
}
