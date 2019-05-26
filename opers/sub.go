package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/sub"
)

// Sub attempts to "subtract" type `b` from type `a`.
// Supports:
//	* int
//	* float64
//	* faces.Sub
//	* types.Floater
//	* types.Integer
func Sub(a, b interface{}) (interface{}, error) {
	switch at := a.(type) {
	case faces.Sub:
		return at.Sub(b)
	case int:
		return sub.Int(at, b)
	case float64:
		return sub.Float(at, b)
	case faces.Int:
		return sub.Int(at.Int(), b)
	case faces.Float:
		return sub.Float(at.Float(), b)
	}

	return nil, sub.Cant(a, b)
}
