package opers

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
	"github.com/gobuffalo/lush/opers/internal/add"
)

// Add `a + b`
// Supports:
//	* int
//	* float64
//	* string
//	* []interface{}
//	* map[string]interface{}
//	* fmt.Stringer
//	* faces.Add
//	* faces.Int
//	* faces.Float
//	* faces.Slice
//	* faces.Map
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
		return add.String(at, b)
	case map[string]interface{}:
		return add.Map(at, b)
	case faces.Map:
		return add.Map(at.Map(), b)
	case fmt.Stringer:
		return add.String(at.String(), b)
	case faces.Int:
		return add.Int(at.Int(), b)
	case faces.Float:
		return add.Float(at.Float(), b)
	case faces.Slice:
		return add.Slice(at.Slice(), b)
	}

	return nil, add.Cant(a, b)
}
