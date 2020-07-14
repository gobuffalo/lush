package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/gte"
)

// GreaterThanEqualTo `a <= b`
// Supports:
//	* int
//	* float64
//	* string
//	* []interface{}
//	* fmt.Stringer
//	* faces.Add
//	* faces.Int
//	* faces.Float
//	* faces.Slice
func GreaterThanEqualTo(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case faces.GreaterThanEqualTo:
		return at.GreaterThanEqualTo(b)
	case int:
		return gte.Int(at, b)
	case float64:
		return gte.Float(at, b)
	case string:
		return gte.String(at, b)
	case faces.Int:
		return gte.Int(at.Int(), b)
	case faces.Float:
		return gte.Float(at.Float(), b)
	case fmt.Stringer:
		return gte.String(at.String(), b)
	}
	return false, gte.Cant(a, b)
}
