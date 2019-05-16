package ast

import "fmt"

type Returned struct {
	Value interface{}
}

func (r Returned) String() string {
	if st, ok := r.Value.(fmt.Stringer); ok {
		return st.String()
	}
	return fmt.Sprint(r.Value)
}

func NewReturned(i interface{}) Returned {
	if r, ok := i.(Returned); ok {
		return r
	}
	if i == nil {
		return Returned{}
	}
	if ii, err := toII(i); err == nil {
		if len(ii) == 0 {
			return Returned{}
		}
		if len(ii) == 1 {
			i = ii[0]
			if r, ok := i.(Returned); ok {
				return r
			}
			if ri, ok := i.(interfacer); ok {
				i = ri.Interface()
			}
			if ri, ok := i.([]interface{}); ok {
				return NewReturned(ri)
			}
			return Returned{Value: i}
		}
		return Returned{Value: ii}
	}
	if ri, ok := i.(interfacer); ok {
		i = ri.Interface()
	}
	return Returned{Value: i}
}
