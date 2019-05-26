package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/lt"
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
		return lt.Int(at, b)
	case float64:
		return lt.Float(at, b)
	case string:
		return lt.String(at, b)
	case fmt.Stringer:
		return lt.String(at.String(), b)
	case faces.Int:
		return lt.Int(at.Int(), b)
	case faces.Float:
		return lt.Float(at.Float(), b)
	}
	return false, lt.Cant(a, b)
}
