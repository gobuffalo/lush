package add

import (
	"github.com/gobuffalo/lush/faces"
)

// Float ...
func Float(at float64, b interface{}) (interface{}, error) {
	switch bt := b.(type) {
	case float64:
		return at + bt, nil
	case int:
		return at + float64(bt), nil
	case faces.Float:
		return at + bt.Float(), nil
	case faces.Int:
		return at + float64(bt.Int()), nil
	}
	return 0.0, Cant(at, b)
}
