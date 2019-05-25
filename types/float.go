package types

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/lush/faces"
)

type Floater float64

func (f Floater) Float() float64 {
	return float64(f)
}

// Float ...
func Float(i interface{}) (float64, error) {
	switch t := i.(type) {
	case float64:
		return t, nil
	case faces.Float:
		return t.Float(), nil
	default:
		a, err := strconv.Atoi(fmt.Sprintf("%s", i))
		if err != nil {
			return 0, err
		}
		return float64(a), nil
	}
}
