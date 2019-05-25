package types

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/lush/faces"
)

type Integer int

func (i Integer) Int() int {
	return int(i)
}

func Int(i interface{}) (int, error) {
	switch t := i.(type) {
	case int:
		return t, nil
	case faces.Int:
		return t.Int(), nil
	default:
		return strconv.Atoi(fmt.Sprintf("%s", i))
	}
}
