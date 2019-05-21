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

func floats(i ...interface{}) ([]float64, error) {
	var fl []float64
	for _, x := range i {
		switch f := x.(type) {
		case Float:
			fl = append(fl, f.Value)
		case float64:
			fl = append(fl, f)
		case int:
			fl = append(fl, float64(f))
		default:
			return fl, fmt.Errorf("expected float64 got %T", x)
		}
	}
	return fl, nil
}

func ints(i ...interface{}) ([]int, error) {
	var fl []int
	for _, x := range i {
		switch f := x.(type) {
		case Integer:
			fl = append(fl, f.Value)
		case int:
			fl = append(fl, f)
		default:
			return fl, fmt.Errorf("expected int got %T", x)
		}
	}
	return fl, nil
}

func stringSlice(c *Context, i ...interface{}) ([]string, error) {
	var fl []string
	for _, x := range i {
		switch f := x.(type) {
		case string:
			fl = append(fl, f)
		case String:
			fl = append(fl, f.Original)
		case Ident:
			s, err := f.Exec(c)
			if err != nil {
				return fl, err
			}
			sx, err := stringSlice(c, s)
			if err != nil {
				return fl, err
			}
			fl = append(fl, sx...)
		default:
			return fl, fmt.Errorf("expected string got %T", x)
		}
	}
	return fl, nil
}
