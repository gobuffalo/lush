package types

import (
	"fmt"
	"strconv"
)

type Integer interface {
	Int() int
}

func Int(i interface{}) (int, error) {
	switch t := i.(type) {
	case int:
		return t, nil
	case Integer:
		return t.Int(), nil
	default:
		return strconv.Atoi(fmt.Sprintf("%s", i))
	}
}
