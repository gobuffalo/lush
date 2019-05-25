package lessthan

import (
	"fmt"
	"strconv"

	"github.com/gobuffalo/lush/faces"
)

func Float(at float64, b interface{}) (bool, error) {
	switch bt := b.(type) {
	case float64:
		return float64(at) < bt, nil
	case int:
		return at < float64(bt), nil
	case faces.Float:
		return float64(at) < bt.Float(), nil
	case faces.Int:
		return at < float64(bt.Int()), nil
	case string:
		toi, err := strconv.Atoi(bt)
		if err != nil {
			return false, err
		}
		return at < float64(toi), nil
	case fmt.Stringer:
		toi, err := strconv.Atoi(bt.String())
		if err != nil {
			return false, err
		}
		return at < float64(toi), nil
	}
	return false, fmt.Errorf("can not compare %T with %T", at, b)
}
