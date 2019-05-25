package types

import (
	"fmt"
	"reflect"
)

func Slice(i interface{}) []interface{} {
	if ii, ok := i.([]interface{}); ok {
		return ii
	}
	var res []interface{}

	rv := reflect.Indirect(reflect.ValueOf(i))
	kind := rv.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < rv.Len(); i++ {
			el := rv.Index(i)
			res = append(res, el.Interface())
		}
		return res
	}

	res = append(res, i)

	return res
}

func StringSlice(i interface{}) ([]string, error) {
	if s, ok := i.([]string); ok {
		return s, nil
	}
	var res []string
	for _, el := range Slice(i) {
		switch t := el.(type) {
		case string:
			res = append(res, t)
		case fmt.Stringer:
			res = append(res, t.String())
		default:
			return res, fmt.Errorf("%T is not a string", t)
		}
	}

	return res, nil
}
