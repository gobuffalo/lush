package add

import (
	"github.com/gobuffalo/lush/types"
)

func Float(at float64, b interface{}) (interface{}, error) {
	switch bt := b.(type) {
	case float64:
		return at + bt, nil
	case types.Integer:
		return at + float64(bt.Int()), nil
	case types.Floater:
		return at + bt.Float(), nil
	case int:
		return at + float64(bt), nil
	}
	return 0.0, Cant(at, b)
}
