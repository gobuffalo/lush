package types

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
)

// Mapper ...
type Mapper map[string]interface{}

// Map ...
func (m Mapper) Map() map[string]interface{} {
	return map[string]interface{}(m)
}

// Map ...
func Map(i interface{}) (map[string]interface{}, error) {
	switch t := i.(type) {
	case map[string]interface{}:
		return t, nil
	case faces.Map:
		return t.Map(), nil
	}

	return nil, fmt.Errorf("%T is not a map", i)
}
