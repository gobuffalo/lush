package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/gt"
)

// GreaterThan `a < b`
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
func GreaterThan(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case faces.GreaterThan:
		return at.GreaterThan(b)
	case int:
		return gt.Int(at, b)
	case float64:
		return gt.Float(at, b)
	case string:
		return gt.String(at, b)
	case fmt.Stringer:
		return gt.String(at.String(), b)
	case faces.Int:
		return gt.Int(at.Int(), b)
	case faces.Float:
		return gt.Float(at.Float(), b)
	}
	return false, gt.Cant(a, b)
}
