package sub

import (
	"github.com/gobuffalo/lush/types"
)

func Int(at int, b interface{}) (interface{}, error) {
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
	return 0, Cant(at, b)
}
