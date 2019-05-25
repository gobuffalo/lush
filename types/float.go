package types

import (
	"fmt"
	"strconv"
)

// Floater ...
type Floater interface {
	Float() float64
}

func Float(i interface{}) (float64, error) {
	switch t := i.(type) {
	case float64:
		return t, nil
	case Floater:
		return t.Float(), nil
	default:
		a, err := strconv.Atoi(fmt.Sprintf("%s", i))
		if err != nil {
			return 0, err
		}
		return float64(a), nil
	}
}
