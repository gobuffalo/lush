package add

import "github.com/gobuffalo/lush/faces"

// Slice ...
func Slice(at []interface{}, b interface{}) (interface{}, error) {
	switch bt := b.(type) {
	case []interface{}:
		return append(at, bt...), nil
	case faces.Slice:
		return append(at, bt.Slice()...), nil
	default:
		return append(at, bt), nil
	}
	return nil, Cant(at, b)
}
