package types

import (
	"fmt"

	"github.com/gobuffalo/lush/faces"
)

type Valuer string

func (s Valuer) Value() string {
	return fmt.Sprintf("%v", s)
}

// Value ...
func Value(i interface{}) string {
	switch t := i.(type) {
	case faces.Value:
		return t.Value()
	default:
		return fmt.Sprintf("%v", i)
	}
}
