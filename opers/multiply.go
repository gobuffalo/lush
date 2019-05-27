package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/multiply"
)

// Multiply attempts to "multiply" type `a` with type `b`.
// Supports:
//	* int
//	* float64
//	* faces.Multiply
//	* faces.Float
//	* faces.Int
func Multiply(a, b interface{}) (interface{}, error) {
	switch at := a.(type) {
	case faces.Multiply:
		return at.Multiply(b)
	case int:
		return multiply.Int(at, b)
	case float64:
		return multiply.Float(at, b)
	case faces.Int:
		return multiply.Int(at.Int(), b)
	case faces.Float:
		return multiply.Float(at.Float(), b)
	}

	return nil, multiply.Cant(a, b)
}
