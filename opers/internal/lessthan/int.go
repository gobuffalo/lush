package lessthan

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/lush/faces"
)

// Int ...
func Int(at int, b interface{}) (bool, error) {
	switch bt := b.(type) {
	case int:
		return at < bt, nil
	case float64:
		return float64(at) < bt, nil
	case faces.Int:
		return at < bt.Int(), nil
	case faces.Float:
		return float64(at) < bt.Float(), nil
	case string:
		toi, err := strconv.Atoi(bt)
		if err != nil {
			return false, err
		}
		return at < toi, nil
	case fmt.Stringer:
		toi, err := strconv.Atoi(bt.String())
		if err != nil {
			return false, err
		}
		return at < toi, nil
	}
	return false, Cant(at, b)

}
