package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/lessthan"
)

// LessThan `a < b`
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
func LessThan(a, b interface{}) (bool, error) {
	switch at := a.(type) {
	case faces.LessThan:
		return at.LessThan(b)
	case int:
		return lessthan.Int(at, b)
	case float64:
		return lessthan.Float(at, b)
	case string:
		return lessthan.String(at, b)
	case fmt.Stringer:
		return lessthan.String(at.String(), b)
	case faces.Int:
		return lessthan.Int(at.Int(), b)
	case faces.Float:
		return lessthan.Float(at.Float(), b)
	}
	return false, lessthan.Cant(a, b)
}
