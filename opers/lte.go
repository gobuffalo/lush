package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/lte"
)

// LessThanEqualTo `a <= b`
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
func LessThanEqualTo(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case faces.LessThanEqualTo:
		return at.LessThanEqualTo(b)
	case int:
		return lte.Int(at, b)
	case float64:
		return lte.Float(at, b)
	case string:
		return lte.String(at, b)
	case fmt.Stringer:
		return lte.String(at.String(), b)
	case faces.Int:
		return lte.Int(at.Int(), b)
	case faces.Float:
		return lte.Float(at.Float(), b)
	}
	return false, lte.Cant(a, b)
}
