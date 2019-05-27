package opers

import (
	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/divide"
)

// Divide attempts to "divide" type `b` from type `a`.
// Supports:
//	* int
//	* float64
//	* faces.Divide
//	* faces.Float
//	* faces.Int
func Divide(a, b interface{}) (interface{}, error) {
	switch at := a.(type) {
	case faces.Divide:
		return at.Divide(b)
	case int:
		return divide.Int(at, b)
	case float64:
		return divide.Float(at, b)
	case faces.Int:
		return divide.Int(at.Int(), b)
	case faces.Float:
		return divide.Float(at.Float(), b)
	}

	return nil, divide.Cant(a, b)
}
