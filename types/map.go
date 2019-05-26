package types

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
)

// Mapper ...
type Mapper map[interface{}]interface{}

// Map ...
func (m Mapper) Map() map[interface{}]interface{} {
	return map[interface{}]interface{}(m)
}

// Map ...
func Map(i interface{}) (map[interface{}]interface{}, error) {
	switch t := i.(type) {
	case map[interface{}]interface{}:
		return t, nil
	case faces.Map:
		return t.Map(), nil
	}

	return nil, fmt.Errorf("%T is not a map", i)
}
