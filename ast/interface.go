package ast

import "fmt"

func toII(i interface{}) ([]interface{}, error) {
	ii, ok := i.([]interface{})
	if !ok {
		return ii, fmt.Errorf("expected []interface{} got %T", i)
	}
	return ii, nil
}

func flatten(ii []interface{}) []interface{} {
	var res []interface{}
	for _, i := range ii {
		if i == nil {
			continue
		}
		switch t := i.(type) {
		case []interface{}:
			res = append(res, flatten(t)...)
		case Noop:
		default:
			res = append(res, i)
		}
	}

	return res
}
