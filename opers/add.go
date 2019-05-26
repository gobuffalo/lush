package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/add"
	"github.com/gobuffalo/lush/types"
)

// Add attempts to "add" type `a` with type `b`.
// Supports:
//	* int
//	* float64
//	* faces.Add
//	* types.Floater
//	* types.Integer
func Add(a, b interface{}) (interface{}, error) {
	switch at := a.(type) {
	case faces.Add:
		return at.Add(b)
	case int:
		return add.Int(at, b)
	case float64:
		return add.Float(at, b)
	case []interface{}:
		return add.Slice(at, b)
	case string:
		switch bt := b.(type) {
		case string:
			return at + bt, nil
		}
	case types.Integer:
		return add.Int(at.Int(), b)
	case types.Floater:
		return add.Float(at.Float(), b)
	}

	return nil, add.Cant(a, b)
}
